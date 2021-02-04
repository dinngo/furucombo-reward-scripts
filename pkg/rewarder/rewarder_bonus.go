package rewarder

import (
	"path"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// BonusRewarderRequiredFieldNames bonus rewarder required field names
var BonusRewarderRequiredFieldNames = []string{"name", "startBlock", "endBlock", "blocks", "cubes", "pool"}

// BonusRewarder struct
type BonusRewarder struct {
	config *Config

	txs              []common.Hash
	tradingMap       TradingMap
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
		filepath:   path.Join(r.config.RoundDir(), "txs.json"),
		startBlock: r.config.StartBlock,
		endBlock:   r.config.EndBlock,
	}

	if err := task.LoadTxsFromFile(); err != nil {
		if err := task.GetTxsFromEtherscan(); err != nil {
			return err
		}

		if err := task.SaveTxsToFile(); err != nil {
			return err
		}
	}

	r.txs = task.txs

	return nil
}

// LoadTradings load tradings
func (r *BonusRewarder) LoadTradings() error {
	task := LoadTradingsTask{
		filepath:       path.Join(r.config.RoundDir(), "tradings.json"),
		startTimestamp: r.config.startTimestamp,
		endTimestamp:   r.config.endTimestamp,
		txs:            r.txs,
		cubeFinders:    r.config.cubeFinders,
	}

	if err := task.LoadTradingsFromFile(); err != nil {
		if err := task.GetTradingsFromTxs(); err != nil {
			return err
		}

		task.RankTradings()

		if err := task.SaveTradingsToFile(); err != nil {
			return err
		}
	}

	r.tradingMap = task.tradingMap

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

// LoadStakingStaked load staked
func (r *BonusRewarder) LoadStakingStaked() error {
	task := LoadStakingStakedTask{
		filepath:        path.Join(r.config.RoundDir(), "staked.json"),
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
		filepath:         path.Join(r.config.RoundDir(), "stakings.json"),
		round:            r.config.Round,
		duration:         decimal.NewFromInt(int64(r.config.Blocks())),
		baseAmount:       r.config.Pool.BaseAmount,
		stakingStakedMap: r.stakingStakedMap,
		tradingMap:       r.tradingMap,
		poolAddress:      r.config.Pool.Address,
		startBlock:       r.config.StartBlock,
		endBlock:         r.config.EndBlock,
	}

	if err := task.LoadStakingsFromFile(); err != nil {
		if err := task.MakeStakingPoolDir(); err != nil {
			return err
		}

		if err := task.InitStakings(); err != nil {
			return err
		}

		if err := task.UpdateStakingsByStakingEvents(); err != nil {
			return err
		}

		if err := task.CalcStakingsWeightWithTradingRank(); err != nil {
			return err
		}

		if err := task.SaveStakingsToFile(); err != nil {
			return err
		}

		r.stakingMap = task.stakingMap
	}

	return nil
}

// LoadRewards load rewards
func (r *BonusRewarder) LoadRewards() error {
	task := LoadRewardsTask{
		filepath:        path.Join(r.config.RoundDir(), "rewards.json"),
		rewardWeightMap: r.stakingMap.ToRewardWeightMap(),
		rewardAmount:    r.config.Pool.RewardAmount,
	}

	if err := task.LoadRewardsFromFile(); err != nil {
		if err := task.CalcRewardsByWeight(); err != nil {
			return err
		}

		if err := task.SaveRewardsToFile(); err != nil {
			return err
		}
	}

	r.rewardMap = task.rewardMap

	return nil
}

// GenerateRewardsMerkleTree generate rewards merkle tree
func (r *BonusRewarder) GenerateRewardsMerkleTree() error {
	task := GenerateRewardMerkleTreeTask{
		rewardMap:                      r.rewardMap,
		rewardMerkleTreeLeavesFilepath: path.Join(r.config.RoundDir(), "merkle_tree_leaves.json"),
		rewardMerkleProofsFilepath:     path.Join(r.config.RoundDir(), "merkle_proofs.json"),
	}

	if err := task.CheckMerkleTreeFiles(); err != nil {
		if err := task.GenerateRewardMerkleTree(); err != nil {
			return err
		}

		if err := task.SaveRewardMerkleTreeLeavesToFile(); err != nil {
			return err
		}

		if err := task.SaveRewardMerkleProofsToFile(); err != nil {
			return err
		}
	}

	return nil
}
