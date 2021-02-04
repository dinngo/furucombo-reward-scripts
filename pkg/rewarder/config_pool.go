package rewarder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// PoolConfig pool config
type PoolConfig struct {
	Address      common.Address    `json:"address"`
	RewardAmount decimal.Decimal   `json:"reward_amount"`
	BaseAmount   decimal.Decimal   `json:"base_amount"`
	Dataset      PoolDatasetConfig `json:"dataset"`
}

// PoolDatasetConfig dataset config
type PoolDatasetConfig struct {
	Name       string `json:"name"`
	StartBlock uint64 `json:"startBlock"`
}

// DatasetFilepath get dataset filepath
func (c *PoolConfig) DatasetFilepath() string {
	return path.Join("dataset", fmt.Sprintf("%s.json", c.Dataset.Name))
}

// SavePool save pool to file
func (c *Config) SavePool() error {
	filepath := path.Join(c.RoundDir(), "pool.json")

	if _, err := os.Stat(filepath); err == nil {
		return nil
	}

	log.Printf("saving pool: ./%s", filepath)

	data, err := json.MarshalIndent(c.Pool, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}

// SavePools save pools to file
func (c *Config) SavePools() error {
	filepath := path.Join(c.RoundDir(), "pools.json")

	if _, err := os.Stat(filepath); err == nil {
		return nil
	}

	log.Printf("saving pools: ./%s", filepath)

	data, err := json.MarshalIndent(c.Pools, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}
