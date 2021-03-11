package rewarder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/etherscan"
)

// LoadTradingCountsTask load trading counts task
type LoadTradingCountsTask struct {
	rootpath   string
	startBlock uint64
	endBlock   uint64

	filepath        string
	tradingCountMap TradingCountMap
}

// LoadFromFile load from file
func (t *LoadTradingCountsTask) LoadFromFile() error {
	t.filepath = path.Join(t.rootpath, "trading_counts.json")

	log.Printf("loading trading counts from ./%s", t.filepath)

	if _, err := os.Stat(t.filepath); err != nil {
		log.Println("trading counts file not found")
		return err
	}

	data, err := ioutil.ReadFile(t.filepath)
	if err != nil {
		return err
	}

	if err := jsonex.Unmarshal(data, &t.tradingCountMap); err != nil {
		return err
	}

	return nil
}

// GetFromEtherscan get trading counts from etherscan
func (t *LoadTradingCountsTask) GetFromEtherscan() error {
	log.Printf("getting trading counts from etherscan")

	apiKey := os.Getenv("ETHERSCAN_API_KEY")

	if len(apiKey) == 0 {
		log.Fatalln("env ETHERSCAN_API_KEY can't be blank")
	}

	client := etherscan.NewClient(&http.Client{Timeout: 10 * time.Second}, apiKey)

	t.tradingCountMap = make(TradingCountMap)
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

		for _, tx := range txs1 {
			if tx.IsError == 1 {
				continue
			}

			t.tradingCountMap[tx.From]++
			log.Printf("found account %s in tx: %s", tx.From.String(), tx.Hash.String())
		}
	}

	return nil
}

// SaveToFile save trading counts to file
func (t *LoadTradingCountsTask) SaveToFile() error {
	log.Printf("saving trading counts to ./%s", t.filepath)

	data, err := json.MarshalIndent(t.tradingCountMap, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(t.filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}

// Execute execute
func (t *LoadTradingCountsTask) Execute() error {
	if err := t.LoadFromFile(); err != nil {
		if err := t.GetFromEtherscan(); err != nil {
			return err
		}

		if err := t.SaveToFile(); err != nil {
			return err
		}
	}

	return nil
}
