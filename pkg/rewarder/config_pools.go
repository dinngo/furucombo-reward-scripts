package rewarder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// PoolConfig pool config
type PoolConfig struct {
	Address      common.Address  `json:"address"`
	RewardAmount decimal.Decimal `json:"reward_amount"`
	BaseAmount   decimal.Decimal `json:"base_amount"`
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
