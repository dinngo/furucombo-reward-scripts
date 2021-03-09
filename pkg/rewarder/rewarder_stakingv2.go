package rewarder

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// StakingV2RewarderRequiredFieldNames staking v2 rewarder required field names
var StakingV2RewarderRequiredFieldNames = []string{"name", "startBlock", "endBlock", "blocks", "pools"}

// StakingV2Rewarder struct
type StakingV2Rewarder struct {
	config *Config

	tradingCountMap   TradingCountMap
	stakingsEventMap  map[common.Address]StakingEventMap
	stakingsStakedMap map[common.Address]StakingStakedMap
	stakingsMap       map[common.Address]StakingMap
	rewardsMap        map[common.Address]RewardMap
}

// NewStakingV2Rewarder new staking v2 rewarder
func NewStakingV2Rewarder(config *Config) *StakingV2Rewarder {
	return &StakingV2Rewarder{config: config}
}

// MakeStakingPoolDir make staking pool dir
func (r *StakingV2Rewarder) MakeStakingPoolDir() error {
	for _, pool := range r.config.Pools {
		poolDir := path.Join(r.config.RoundDir(), pool.Address.String())

		log.Printf("making staking pool dir: ./%s/", poolDir)

		if err := os.MkdirAll(poolDir, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

// LoadTradingCount load trading count
func (r *StakingV2Rewarder) LoadTradingCount() error {
	task := LoadTradingCountsTask{
		rootpath:   r.config.RoundDir(),
		startBlock: r.config.StartBlock,
		endBlock:   r.config.EndBlock,
	}

	if err := task.Execute(); err != nil {
		return err
	}

	r.tradingCountMap = task.tradingCountMap

	return nil
}

// LoadStakingsDataset load stakings dataset
func (r *StakingV2Rewarder) LoadStakingsDataset() error {
	r.stakingsEventMap = make(map[common.Address]StakingEventMap)

	for _, pool := range r.config.Pools {
		task := LoadStakingDatasetTask{
			filepath:    pool.DatasetFilepath(),
			poolAddress: pool.Address,
			startBlock:  pool.Dataset.StartBlock,
		}

		if err := task.Execute(); err != nil {
			return err
		}

		r.stakingsEventMap[pool.Address] = task.stakingDataset.EventMap
	}

	return nil
}

// LoadStakingsStaked load stakings staked
func (r *StakingV2Rewarder) LoadStakingsStaked() error {
	r.stakingsStakedMap = make(map[common.Address]StakingStakedMap)

	for _, pool := range r.config.Pools {
		fmt.Println("ZD", pool.Address.String())
		task := LoadStakingStakedTask{
			rootpath:        path.Join(r.config.RoundDir(), pool.Address.String()),
			stakingEventMap: r.stakingsEventMap[pool.Address],
			startBlock:      r.config.StartBlock,
		}

		if err := task.Execute(); err != nil {
			return err
		}

		r.stakingsStakedMap[pool.Address] = task.stakingStakedMap
	}

	return nil
}

// LoadStakings load stakings
func (r *StakingV2Rewarder) LoadStakings() error {
	r.stakingsMap = make(map[common.Address]StakingMap)

	for _, pool := range r.config.Pools {
		task := LoadStakingsTask{
			rootpath:         path.Join(r.config.RoundDir(), pool.Address.String()),
			round:            r.config.Round,
			duration:         decimal.NewFromInt(int64(r.config.Blocks())),
			baseAmount:       pool.BaseAmount,
			stakingStakedMap: r.stakingsStakedMap[pool.Address],
			tradingCountMap:  r.tradingCountMap,
			poolAddress:      pool.Address,
			startBlock:       r.config.StartBlock,
			endBlock:         r.config.EndBlock,
		}

		if err := task.Execute(); err != nil {
			return err
		}

		r.stakingsMap[pool.Address] = task.stakingMap
	}

	return nil
}

// LoadRewards load rewards
func (r *StakingV2Rewarder) LoadRewards() error {
	r.rewardsMap = make(map[common.Address]RewardMap)

	for _, pool := range r.config.Pools {
		task := LoadRewardsTask{
			rootpath:        path.Join(r.config.RoundDir(), pool.Address.String()),
			rewardWeightMap: r.stakingsMap[pool.Address].ToRewardWeightMap(),
			rewardAmount:    pool.RewardAmount,
		}

		if err := task.Execute(); err != nil {
			return err
		}

		r.rewardsMap[pool.Address] = task.rewardMap
	}

	return nil
}

// GenerateRewardsMerkleTree generate rewards merkle tree
func (r *StakingV2Rewarder) GenerateRewardsMerkleTree() error {
	for _, pool := range r.config.Pools {
		task := GenerateRewardMerkleTreeTask{
			rootpath:  path.Join(r.config.RoundDir(), pool.Address.String()),
			rewardMap: r.rewardsMap[pool.Address],
		}

		if err := task.Execute(); err != nil {
			return err
		}
	}

	return nil
}
