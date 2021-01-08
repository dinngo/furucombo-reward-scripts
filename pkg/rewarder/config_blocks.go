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
)

// Blocks blocks
func (c *Config) Blocks() uint64 {
	if c.blocks == 0 {
		c.blocks = c.EndBlock - c.StartBlock
	}

	return c.blocks
}

// GetBlockTimestamps get block timestamps
func (c *Config) GetBlockTimestamps() error {
	log.Printf("getting block timestamps: %d, %d", c.StartBlock, c.EndBlock)

	startHeader, err := ethereum.Client().HeaderByNumber(context.Background(), big.NewInt(int64(c.StartBlock)))
	if err != nil {
		return err
	}
	c.startTimestamp = startHeader.Time

	endHeader, err := ethereum.Client().HeaderByNumber(context.Background(), big.NewInt(int64(c.EndBlock)))
	if err != nil {
		return err
	}
	c.endTimestamp = endHeader.Time

	return nil
}

// SaveBlocks save blocks to file
func (c *Config) SaveBlocks() error {
	filepath := path.Join(c.RoundDir(), "blocks.json")

	if _, err := os.Stat(filepath); err == nil {
		return nil
	}

	log.Printf("saving blocks: ./%s", filepath)

	config := map[string]interface{}{
		"startBlock": c.StartBlock,
		"endBlock":   c.EndBlock,
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}
