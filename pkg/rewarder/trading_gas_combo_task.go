package rewarder

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/coingecko"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/etherscan"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

var divisionPrecision int32 = 18

// LoadTradingGasComboTask load trading gas combo task
type LoadTradingGasComboTask struct {
	rootpath       string
	startBlock     uint64
	endBlock       uint64
	startTimestamp uint64
	endTimestamp   uint64
	maxGasUsed     decimal.Decimal

	filepath         string
	gasPrices        []decimal.Decimal
	comboPricesInEth []decimal.Decimal

	// To files
	gasUsedMap         GasUsedMap
	txs                []common.Hash
	params             TradingGasComboParams
	tradingGasComboMap TradingGasComboMap
}

// LoadFromFile load from file
func (t *LoadTradingGasComboTask) LoadFromFile() error {
	t.filepath = path.Join(t.rootpath, "trading_gas_combo.json")

	log.Printf("loading trading gas combo from ./%s", t.filepath)

	if _, err := os.Stat(t.filepath); err != nil {
		log.Println("trading gas combo file not found")
		return err
	}

	data, err := ioutil.ReadFile(t.filepath)
	if err != nil {
		return err
	}

	if err := jsonex.Unmarshal(data, &t.tradingGasComboMap); err != nil {
		return err
	}

	return nil
}

// GetFromEtherscan get gas used from etherscan
func (t *LoadTradingGasComboTask) GetFromEtherscan() error {
	log.Printf("getting gas used from etherscan")

	apiKey := os.Getenv("ETHERSCAN_API_KEY")

	if len(apiKey) == 0 {
		log.Fatalln("env ETHERSCAN_API_KEY can't be blank")
	}

	client := etherscan.NewClient(&http.Client{Timeout: 10 * time.Second}, apiKey)

	t.gasUsedMap = make(GasUsedMap)
	t.txs = make([]common.Hash, 0)
	for _, address := range furucombo.ProxyAddresses() {
		params := etherscan.Params{
			"address":    address.String(),
			"startBlock": t.startBlock,
			"endBlock":   t.endBlock,
			"sort":       "asc",
		}

		txs1, err := client.AccountTxs(params)
		if err != nil {
			return err
		}

		if len(txs1) == 0 {
			return errors.New("0 tx from etherscan")
		}

		for _, tx := range txs1 {
			// Accumulate gas used
			t.gasUsedMap[tx.From] = t.gasUsedMap[tx.From].Add(decimal.New(int64(tx.GasUsed), 0))

			// Store tx hash
			t.txs = append(t.txs, tx.Hash)

			// Store gas prices
			gasPrice := decimal.NewFromBigInt((*big.Int)(tx.GasPrice), -18) // in ether unit
			t.gasPrices = append(t.gasPrices, gasPrice)

			log.Printf("found account %s uses gas %d price %s in tx: %s", tx.From.String(), tx.GasUsed, gasPrice.String(), tx.Hash.String())
		}
	}

	return nil
}

func (t *LoadTradingGasComboTask) GetComboPricesInEth() error {
	// Get coingecko price
	client := coingecko.NewClient(&http.Client{Timeout: 10 * time.Second})
	comboAddress := "0xfFffFffF2ba8F66D4e51811C5190992176930278"
	tokenMarketChart, err := client.CoinsContractMarketChartRange(comboAddress, "eth", t.startTimestamp, t.endTimestamp)
	if err != nil {
		return err
	}
	prices := tokenMarketChart.Prices
	for _, v := range prices {
		t.comboPricesInEth = append(t.comboPricesInEth, decimal.NewFromFloat(v[1]))
	}

	return nil
}

func (t *LoadTradingGasComboTask) CalcParams() error {
	// Calc gas price
	sum := decimal.Zero
	for _, v := range t.gasPrices {
		sum = sum.Add(v)
	}
	t.params.GasPrice = sum.DivRound(decimal.New(int64(len(t.gasPrices)), 0), divisionPrecision)

	// Calc combo price
	sum = decimal.Zero
	for _, v := range t.comboPricesInEth {
		sum = sum.Add(v)
	}
	t.params.ComboPrice = sum.DivRound(decimal.New(int64(len(t.comboPricesInEth)), 0), divisionPrecision)

	return nil
}

func (t *LoadTradingGasComboTask) CalcCombo() error {
	// Calc reimbursement combo amount = gasUsed * gasPrice / comboPrice
	sum := decimal.Zero
	t.tradingGasComboMap = make(TradingGasComboMap)
	for user, gasUsed := range t.gasUsedMap {
		// Limit to max gas used
		if gasUsed.GreaterThan(t.maxGasUsed) {
			gasUsed = t.maxGasUsed
		}
		amount := gasUsed.Mul(t.params.GasPrice).DivRound(t.params.ComboPrice, divisionPrecision)
		t.tradingGasComboMap[user] = amount
		sum = sum.Add(amount)
	}
	log.Printf("print combo sum: %s", sum.String())

	return nil
}

// SaveGasUsed save gas used to file
func (t *LoadTradingGasComboTask) SaveGasUsed() error {
	filepath := path.Join(t.rootpath, "trading_gas_used.json")
	log.Printf("saving gas used to ./%s", filepath)

	data, err := json.MarshalIndent(t.gasUsedMap, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}

// SaveTxs save txs to file
func (t *LoadTradingGasComboTask) SaveTxs() error {
	filepath := path.Join(t.rootpath, "txs.json")
	log.Printf("saving txs to ./%s", filepath)

	data, err := json.MarshalIndent(t.txs, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}

// SaveParams save params to file
func (t *LoadTradingGasComboTask) SaveParams() error {
	filepath := path.Join(t.rootpath, "params.json")
	log.Printf("saving params to ./%s", filepath)

	data, err := json.MarshalIndent(t.params, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}

// SaveToFile save trading gas combo to file
func (t *LoadTradingGasComboTask) SaveToFile() error {
	log.Printf("saving trading gas combo to ./%s", t.filepath)

	data, err := json.MarshalIndent(t.tradingGasComboMap, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(t.filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}

// Execute execute
func (t *LoadTradingGasComboTask) Execute() error {
	if err := t.LoadFromFile(); err != nil {
		if err := t.GetFromEtherscan(); err != nil {
			return err
		}

		if err := t.GetComboPricesInEth(); err != nil {
			return err
		}

		if err := t.CalcParams(); err != nil {
			return err
		}

		if err := t.CalcCombo(); err != nil {
			return err
		}

		if err := t.SaveGasUsed(); err != nil {
			return err
		}

		if err := t.SaveTxs(); err != nil {
			return err
		}

		if err := t.SaveParams(); err != nil {
			return err
		}

		if err := t.SaveToFile(); err != nil {
			return err
		}
	}

	return nil
}
