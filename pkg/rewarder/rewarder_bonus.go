package rewarder

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// BonusRewarderRequiredFieldNames bonus rewarder required field names
var BonusRewarderRequiredFieldNames = []string{"name", "startBlock", "endBlock", "blocks", "cubes", "pool"}

// BonusRewarder struct
type BonusRewarder struct {
	config *Config

	txs              []common.Hash
	tradingVolumeMap TradingVolumeMap
	tradingRankMap   TradingRankMap
	stakingEventMap  StakingEventMap
	stakingStakedMap StakingStakedMap
	stakingMap       StakingMap
	rewardMap        RewardMap
}

// NewBonusRewarder new bonus rewarder
func NewBonusRewarder(config *Config) *BonusRewarder {
	return &BonusRewarder{config: config}
}

// LoadTxs load txs
func (r *BonusRewarder) LoadTxs() error {
	task := LoadTxsTask{
		rootpath:   r.config.RoundDir(),
		startBlock: r.config.StartBlock,
		endBlock:   r.config.EndBlock,
	}

	if err := task.Execute(); err != nil {
		return err
	}

	r.txs = task.txs

	return nil
}

// LoadTradingVolumes load trading volumes
func (r *BonusRewarder) LoadTradingVolumes() error {
	task := LoadTradingVolumesTask{
		rootpath:       r.config.RoundDir(),
		startTimestamp: r.config.startTimestamp,
		endTimestamp:   r.config.endTimestamp,
		txs:            r.txs,
		cubeFinders:    r.config.cubeFinders,
	}

	if err := task.Execute(); err != nil {
		return err
	}

	r.tradingVolumeMap = task.tradingVolumeMap

	return nil
}

// LoadTradingRanks load trading ranks
func (r *BonusRewarder) LoadTradingRanks() error {
	task := LoadTradingRanksTask{
		rootpath:         r.config.RoundDir(),
		volumeCap:        r.config.Pool.VolumeCap,
		tradingVolumeMap: r.tradingVolumeMap,
	}

	if err := task.Execute(); err != nil {
		return err
	}

	r.tradingRankMap = task.tradingRankMap

	return nil
}

// LoadStakingDataset load staking dataset
func (r *BonusRewarder) LoadStakingDataset() error {
	task := LoadStakingDatasetTask{
		filepath:    r.config.Pool.DatasetFilepath(),
		poolAddress: r.config.Pool.Address,
		startBlock:  r.config.Pool.Dataset.StartBlock,
	}

	if err := task.Execute(); err != nil {
		return err
	}

	r.stakingEventMap = task.stakingDataset.EventMap

	return nil
}

// LoadStakingStaked load staking staked
func (r *BonusRewarder) LoadStakingStaked() error {
	task := LoadStakingStakedTask{
		rootpath:        r.config.RoundDir(),
		stakingEventMap: r.stakingEventMap,
		startBlock:      r.config.StartBlock,
	}

	if err := task.Execute(); err != nil {
		return err
	}

	r.stakingStakedMap = task.stakingStakedMap

	return nil
}

// LoadStakings load stakings
func (r *BonusRewarder) LoadStakings() error {
	task := LoadStakingsTask{
		rootpath:         r.config.RoundDir(),
		round:            r.config.Round,
		duration:         decimal.NewFromInt(int64(r.config.Blocks())),
		baseAmount:       r.config.Pool.BaseAmount,
		stakingStakedMap: r.stakingStakedMap,
		tradingRankMap:   r.tradingRankMap,
		poolAddress:      r.config.Pool.Address,
		startBlock:       r.config.StartBlock,
		endBlock:         r.config.EndBlock,
	}

	if err := task.Execute(); err != nil {
		return err
	}

	r.stakingMap = task.stakingMap

	return nil
}

// LoadRewards load rewards
func (r *BonusRewarder) LoadRewards() error {
	task := LoadRewardsTask{
		rootpath:        r.config.RoundDir(),
		rewardWeightMap: r.stakingMap.ToRewardWeightMap(),
		rewardAmount:    r.config.Pool.RewardAmount,
	}

	if err := task.Execute(); err != nil {
		return err
	}

	r.rewardMap = task.rewardMap

	return nil
}

// GenerateRewardsMerkleTree generate rewards merkle tree
func (r *BonusRewarder) GenerateRewardsMerkleTree() error {
	task := GenerateRewardMerkleTreeTask{
		rootpath:  r.config.RoundDir(),
		rewardMap: r.rewardMap,
	}

	if err := task.Execute(); err != nil {
		return err
	}

	return nil
}
