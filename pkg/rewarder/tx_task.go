package rewarder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/etherscan"
)

// LoadTxsTask load txs task
type LoadTxsTask struct {
	filepath   string
	startBlock uint64
	endBlock   uint64

	txs []common.Hash
}

// LoadTxsFromFile load txs from file
func (t *LoadTxsTask) LoadTxsFromFile() error {
	log.Printf("loading txs from ./%s", t.filepath)

	if _, err := os.Stat(t.filepath); err != nil {
		log.Println("txs file not found")
		return err
	}

	data, err := ioutil.ReadFile(t.filepath)
	if err != nil {
		return err
	}

	if err := jsonex.Unmarshal(data, &t.txs); err != nil {
		return err
	}

	return nil
}

// GetTxsFromEtherscan get txs from etherscan
func (t *LoadTxsTask) GetTxsFromEtherscan() error {
	log.Printf("getting txs from etherscan")

	apiKey := os.Getenv("ETHERSCAN_API_KEY")

	if len(apiKey) == 0 {
		log.Fatalln("env ETHERSCAN_API_KEY can't be blank")
	}

	client := etherscan.NewClient(&http.Client{Timeout: 10 * time.Second}, apiKey)

	params := etherscan.Params{
		"address":    furucombo.ProxyAddress().String(),
		"startBlock": t.startBlock,
		"endBlock":   t.endBlock,
		"sort":       "asc",
	}

	t.txs = make([]common.Hash, 0)
	txHashMap := map[common.Hash]struct{}{}

	txs1, err := client.AccountTxs(params)
	if err != nil {
		return err
	}

	for _, tx := range txs1 {
		if tx.IsError == 1 {
			continue
		}

		if _, ok := txHashMap[tx.Hash]; !ok {
			txHashMap[tx.Hash] = struct{}{}
			t.txs = append(t.txs, tx.Hash)

			log.Printf("found tx: %s", tx.Hash.String())
		}
	}

	return nil
}

// SaveTxsToFile save txs to file
func (t *LoadTxsTask) SaveTxsToFile() error {
	log.Printf("saving txs to ./%s", t.filepath)

	data, err := json.MarshalIndent(t.txs, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(t.filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}
