package rewarder

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"path"
	"time"

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

	log.Printf("found block timestamps: %d (%s), %d (%s)",
		c.startTimestamp,
		time.Unix(int64(c.startTimestamp), 0).UTC().Format(time.RFC3339),
		c.endTimestamp,
		time.Unix(int64(c.endTimestamp), 0).UTC().Format(time.RFC3339),
	)

	return nil
}

// SaveBlocks save blocks to file
func (c *Config) SaveBlocks() error {
	filepath := path.Join(c.RoundDir(), "blocks.json")
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

// UpdateEndBlockToCurrentBlock update end block to current block
func (c *Config) UpdateEndBlockToCurrentBlock() error {
	currentBlock, err := ethereum.Client().BlockNumber(context.Background())
	if err != nil {
		return err
	}

	c.EndBlock = currentBlock

	log.Printf("update end block to current block: %d", c.EndBlock)

	return nil
}

// UpdatePolygonEndBlockToCurrentBlock update polygon end block to current block
func (c *Config) UpdatePolygonEndBlockToCurrentBlock() error {
	currentBlock, err := ethereum.ClientPolygon().BlockNumber(context.Background())
	if err != nil {
		return err
	}

	c.PolygonEndBlock = currentBlock

	log.Printf("update polygon end block to current block: %d", c.PolygonEndBlock)

	return nil
}
