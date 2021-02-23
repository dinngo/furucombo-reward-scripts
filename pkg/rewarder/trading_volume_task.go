package rewarder

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
)

// LoadTradingVolumesTask load trading volumes task
type LoadTradingVolumesTask struct {
	rootpath       string
	startTimestamp uint64
	endTimestamp   uint64
	txs            []common.Hash
	cubeFinders    CubeFinders

	filepath         string
	priceOracle      *PriceOracle
	tradingVolumeMap TradingVolumeMap
}

// LoadFromFile load from file
func (t *LoadTradingVolumesTask) LoadFromFile() error {
	t.filepath = path.Join(t.rootpath, "trading_volumes.json")

	log.Printf("loading trading volumes from ./%s", t.filepath)

	if _, err := os.Stat(t.filepath); err != nil {
		log.Println("trading volumes file not found")
		return err
	}

	data, err := ioutil.ReadFile(t.filepath)
	if err != nil {
		return err
	}

	if err := jsonex.Unmarshal(data, &t.tradingVolumeMap); err != nil {
		return err
	}

	return nil
}

// GetFromTxs get trading volumes from txs
func (t *LoadTradingVolumesTask) GetFromTxs() error {
	log.Printf("getting trading volumes from %d txs", len(t.txs))

	priceOracle, err := NewPriceOracle(t.rootpath, t.startTimestamp, t.endTimestamp)
	if err != nil {
		return err
	}

	t.priceOracle = priceOracle
	t.tradingVolumeMap = make(TradingVolumeMap)

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

			// get token
			token, err := ethereum.GetToken(cube.TokenAddress)
			if err != nil {
				return err
			}

			// amount to big unit
			amount := ethereum.ToBigUnit(cube.TokenAmount, token.Decimals)

			// get price
			price := priceOracle.GetClosestPrice(token, timestamp)

			// calc trading volume
			tradingVolume := amount.Mul(price)
			if tradingVolume.IsZero() {
				log.Printf("0 trading volume: [%s] %s", cube.Name, token.Symbol)
			}

			log.Printf("found %s cube: %s %s ($%s)", cube.Name, amount, token.Symbol, tradingVolume.Truncate(3))

			// Add user trading volume
			t.tradingVolumeMap[from] = t.tradingVolumeMap[from].Add(tradingVolume)

			count++
		}

		log.Printf("total found %d cubes", count)
	}

	return nil
}

// SaveToFile save trading volumes to file
func (t *LoadTradingVolumesTask) SaveToFile() error {
	log.Printf("saving trading volumes to ./%s", t.filepath)

	data, err := json.MarshalIndent(t.tradingVolumeMap, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(t.filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}

// Execute execute
func (t *LoadTradingVolumesTask) Execute() error {
	if err := t.LoadFromFile(); err != nil {
		if err := t.GetFromTxs(); err != nil {
			return err
		}

		if err := t.SaveToFile(); err != nil {
			return err
		}

		if err := t.priceOracle.SavePricesToFile(); err != nil {
			return err
		}
	}

	return nil
}
