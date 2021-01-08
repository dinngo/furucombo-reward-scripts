package rewarder

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
)

// LoadTradingsTask load tradings task
type LoadTradingsTask struct {
	filepath    string
	txs         []common.Hash
	cubeFinders CubeFinders
	tokenMap    TokenMap

	tradingMap TradingMap
}

// LoadTradingsFromFile load tradings from file
func (t *LoadTradingsTask) LoadTradingsFromFile() error {
	log.Printf("loading tradings from ./%s", t.filepath)

	if _, err := os.Stat(t.filepath); err != nil {
		log.Println("tradings file not found")
		return err
	}

	data, err := ioutil.ReadFile(t.filepath)
	if err != nil {
		return err
	}

	if err := jsonex.Unmarshal(data, &t.tradingMap); err != nil {
		return err
	}

	return nil
}

// GetTradingsFromTxs get tradings from txs
func (t *LoadTradingsTask) GetTradingsFromTxs() error {
	log.Printf("getting tradings from %d txs", len(t.txs))

	t.tradingMap = make(TradingMap)

	for _, txHash := range t.txs {
		log.Printf("txhash: %s", txHash.String())

		// get tx
		tx, isPending, err := ethereum.Client().TransactionByHash(context.Background(), txHash)
		if err != nil {
			return err
		}
		if isPending {
			return fmt.Errorf("unexpected pending tx: %s", txHash.String())
		}

		// get from address
		from, err := types.NewEIP155Signer(tx.ChainId()).Sender(tx)
		if err != nil {
			return err
		}
		log.Printf("from: %s", from.String())

		// get receipt
		receipt, err := ethereum.Client().TransactionReceipt(context.Background(), txHash)
		if err != nil {
			return err
		}
		if receipt.Status == types.ReceiptStatusFailed {
			return fmt.Errorf("unexpected failed tx: %s", txHash.String())
		}

		// get timestamp
		header, err := ethereum.Client().HeaderByNumber(context.Background(), receipt.BlockNumber)
		if err != nil {
			return err
		}
		timestamp := header.Time
		log.Printf("timestamp: %d block_number: %s", timestamp, receipt.BlockNumber.String())

		// find cubes
		count := 0
		for _, txLog := range receipt.Logs {
			cube, err := t.cubeFinders.Find(txLog)
			if err != nil {
				return err
			}
			if cube == nil {
				continue
			}

			// create trading's data if not exists
			if _, ok := t.tradingMap[from]; !ok {
				t.tradingMap[from] = &Trading{}
			}

			// get token
			token := t.tokenMap[cube.TokenAddress]

			// amount to big unit
			amount := ethereum.ToBigUnit(cube.TokenAmount, token.Decimals)

			// get price
			price := token.GetClosestUSDPrice(timestamp)

			// calc trading volume
			tradingVolume := amount.Mul(price)
			if tradingVolume.IsZero() {
				return fmt.Errorf("unexpected 0 trading volume: [%s] %s", cube.Name, token.Symbol)
			}

			log.Printf("found %s cube: %s %s ($%s)", cube.Name, amount, token.Symbol, tradingVolume.Truncate(3))

			// Add user trading volume
			t.tradingMap[from].Volume = t.tradingMap[from].Volume.Add(tradingVolume)

			count++
		}

		log.Printf("total found %d cubes", count)
	}

	return nil
}

// RankTradings rank tradings
func (t *LoadTradingsTask) RankTradings() {
	for account, curTrading := range t.tradingMap {
		rank := 1
		for _, trading := range t.tradingMap {
			if curTrading.Volume.GreaterThan(trading.Volume) {
				rank++
			}
		}

		t.tradingMap[account].Rank = rank
	}
}

// WeightTradings weight tradings
func (t *LoadTradingsTask) WeightTradings() error {
	rankSum := decimal.Zero
	rankMap := map[common.Address]decimal.Decimal{}

	for account, trading := range t.tradingMap {
		rankMap[account] = decimal.NewFromInt(int64(trading.Rank))
		rankSum = rankSum.Add(rankMap[account])
	}

	for account := range t.tradingMap {
		weight, _ := rankMap[account].QuoRem(rankSum, 27)
		t.tradingMap[account].Weight = &weight
	}

	return nil
}

// SaveTradingsToFile save tradings to file
func (t *LoadTradingsTask) SaveTradingsToFile() error {
	log.Printf("saving tradings to ./%s", t.filepath)

	data, err := json.MarshalIndent(t.tradingMap, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(t.filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}
