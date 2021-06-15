package rewarder

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// validate validate field
func (c *Config) validate(fieldName string) error {

	switch fieldName {
	case "rewards":
		if len(c.RewardMap) == 0 {
			return errors.New("rewards can't be blank")
		}

	case "startBlock":
		if c.StartBlock == 0 {
			return errors.New("startBlock can't be blank or 0")
		}

	case "endBlock":
		if c.EndBlock == 0 {
			return errors.New("endBlock can't be blank or 0")
		}

	case "blocks":
		if c.StartBlock > c.EndBlock {
			return errors.New("startBlock should be less than endBlock")
		}

	case "cubes":
		if len(c.CubeNames) == 0 {
			return errors.New("cubes can't be blank")
		}

	case "pools":
		if len(c.Pools) == 0 {
			return errors.New("pools can't be blank")
		}

		for i, pool := range c.Pools {
			if pool.Address == (common.Address{}) {
				return fmt.Errorf("pools.%d.address can't be blank", i)
			}

			if pool.RewardAmount.IsZero() {
				return fmt.Errorf("pools.%d.reward_amount can't be blank or 0", i)
			}

			if len(pool.Dataset.Name) == 0 {
				return fmt.Errorf("pool.%d.dataset.name can't be blank", i)
			}
		}

	case "pool":
		if c.Pool.Address == (common.Address{}) {
			return errors.New("pool.address can't be blank")
		}

		if c.Pool.RewardAmount.IsZero() {
			return errors.New("pool.reward_amount can't be blank or 0")
		}

		if len(c.Pool.Dataset.Name) == 0 {
			return errors.New("pool.dataset.name can't be blank")
		}

	case "rewardAmount":
		if c.RewardAmount.IsZero() {
			return errors.New("rewardAmount can't be blank or 0")
		}

	case "maxGasUsed":
		if c.MaxGasUsed.IsZero() {
			return errors.New("maxGasUsed can't be blank or 0")
		}

	case "nft":
		if len(c.Nft.Ethereum) == 0 {
			return errors.New("ethereum can't be blank")
		}

		if len(c.Nft.Polygon) == 0 {
			return errors.New("polygon can't be blank")
		}

		for i, address := range c.Nft.Ethereum {
			if address == (common.Address{}) {
				return fmt.Errorf("ethereum.%d.address can't be blank", i)
			}
		}

		for i, address := range c.Nft.Polygon {
			if address == (common.Address{}) {
				return fmt.Errorf("polygon.%d.address can't be blank", i)
			}
		}

		if c.Nft.Boost == 0 {
			return errors.New("boost can't be blank or 0")
		}

		if c.Nft.MaxBoost == 0 {
			return errors.New("maxBoost can't be blank or 0")
		}

	case "polygonEndBlock":
		if c.PolygonEndBlock == 0 {
			return errors.New("polygonEndBlock can't be blank or 0")
		}
	}

	return nil
}

// Validates validate fields
func (c *Config) Validates(fieldNames []string) error {
	if len(c.Name) == 0 {
		return errors.New("name can't be blank")
	}

	for _, fieldName := range fieldNames {
		if err := c.validate(fieldName); err != nil {
			return err
		}
	}

	return nil
}
