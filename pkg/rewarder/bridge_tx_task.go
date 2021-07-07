package rewarder

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

var bridgeMinAmountsMapFilepath = path.Join("config", "bridge_min_amounts.json")

// LoadBridgeTxsTask load bridge txs task
type LoadBridgeTxsTask struct {
	// Exterior
	rootpath string
	endBlock uint64

	// Intermedium
	filepath            string
	bridgeMinAmountsMap map[common.Address]decimal.Decimal

	// File
	bridgeTxMap BridgeTxMap
}

// LoadFromFile load from file
func (t *LoadBridgeTxsTask) LoadFromFile() error {
	t.filepath = path.Join(t.rootpath, "bridge_txs.json")

	log.Printf("loading bridge txs from ./%s", t.filepath)

	if _, err := os.Stat(t.filepath); err != nil {
		log.Println("bridge txs file not found")
		return err
	}

	data, err := ioutil.ReadFile(t.filepath)
	if err != nil {
		return err
	}

	if err := jsonex.Unmarshal(data, &t.bridgeTxMap); err != nil {
		return err
	}

	return nil
}

// LoadBridgeMinAmounts load bridge min amounts
func (t *LoadBridgeTxsTask) LoadBridgeMinAmounts() error {
	if _, err := os.Stat(bridgeMinAmountsMapFilepath); err != nil {
		return err
	}

	data, err := ioutil.ReadFile(bridgeMinAmountsMapFilepath)
	if err != nil {
		return err
	}

	if err := jsonex.Unmarshal(data, &t.bridgeMinAmountsMap); err != nil {
		return err
	}

	return nil
}

func (t *LoadBridgeTxsTask) GetValidBridgeTxs() error {
	t.bridgeTxMap = make(BridgeTxMap)

	events, err := furucombo.GetTokenBridgeEvents(12676516, t.endBlock)
	if err != nil {
		return err
	}
	for _, event := range events {
		user := common.BytesToAddress(event.Topics[1].Bytes())
		token := common.BytesToAddress(event.Topics[2].Bytes())
		amount := decimal.NewFromBigInt(new(big.Int).SetBytes(event.Data), 0)

		// Is valid token and amount
		minAmount, ok := t.bridgeMinAmountsMap[token]
		if !ok || amount.LessThan(minAmount) {
			continue
		}

		t.bridgeTxMap[user] = append(t.bridgeTxMap[user], event.TxHash)
	}

	return nil
}

func (t *LoadBridgeTxsTask) FilterEtherBalance() error {
	for user := range t.bridgeTxMap {
		balance, err := ethereum.Client().BalanceAt(context.Background(), user, big.NewInt(12694370)) // Jun-24-2021 03:00:01 AM +UTC
		if err != nil {
			return (err)
		}

		if balance.Cmp(big.NewInt(0)) == 0 {
			delete(t.bridgeTxMap, user)
		}
	}

	return nil
}

// SaveToFile save bridge txs to file
func (t *LoadBridgeTxsTask) SaveToFile() error {
	log.Printf("saving bridge txs to ./%s", t.filepath)

	data, err := json.MarshalIndent(t.bridgeTxMap, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(t.filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}

// Execute execute
func (t *LoadBridgeTxsTask) Execute() error {
	if err := t.LoadFromFile(); err != nil {
		if err := t.LoadBridgeMinAmounts(); err != nil {
			return err
		}

		if err := t.GetValidBridgeTxs(); err != nil {
			return err
		}

		if err := t.FilterEtherBalance(); err != nil {
			return err
		}

		if err := t.SaveToFile(); err != nil {
			return err
		}
	}

	return nil
}
