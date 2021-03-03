package rewarder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/etherscan"
)

// LoadTxsTask load txs task
type LoadTxsTask struct {
	rootpath   string
	startBlock uint64
	endBlock   uint64

	filepath string
	txs      []common.Hash
}

// LoadFromFile load from file
func (t *LoadTxsTask) LoadFromFile() error {
	t.filepath = path.Join(t.rootpath, "txs.json")

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

// GetFromEtherscan get txs from etherscan
func (t *LoadTxsTask) GetFromEtherscan() error {
	log.Printf("getting txs from etherscan")

	apiKey := os.Getenv("ETHERSCAN_API_KEY")

	if len(apiKey) == 0 {
		log.Fatalln("env ETHERSCAN_API_KEY can't be blank")
	}

	client := etherscan.NewClient(&http.Client{Timeout: 10 * time.Second}, apiKey)

	t.txs = make([]common.Hash, 0)
	for _, address := range furucombo.ProxyAddresses() {
		params := etherscan.Params{
			"address":    address.String(),
			"startBlock": t.startBlock,
			"endBlock":   t.endBlock,
			"sort":       "asc",
		}

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
	}

	return nil
}

// SaveToFile save txs to file
func (t *LoadTxsTask) SaveToFile() error {
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

// Execute execute
func (t *LoadTxsTask) Execute() error {
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
