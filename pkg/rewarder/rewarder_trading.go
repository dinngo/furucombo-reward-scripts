package rewarder

import (
	"path"

	"github.com/ethereum/go-ethereum/common"
)

// TradingRewarderRequiredFieldNames trading rewarder required field names
var TradingRewarderRequiredFieldNames = []string{"name", "blocks", "cubes", "reward_amount"}

// TradingRewarder struct
type TradingRewarder struct {
	config *Config

	tokenMap   TokenMap
	txs        []common.Hash
	tradingMap TradingMap
	rewardMap  RewardMap
}

// NewTradingRewarder new staking rewarder
func NewTradingRewarder(config *Config) *TradingRewarder {
	return &TradingRewarder{config: config}
}

// LoadTokens load tokens
func (r *TradingRewarder) LoadTokens() error {
	task := LoadTokensTask{
		filepath:       path.Join(r.config.RoundDir(), "tokens.json"),
		startTimestamp: r.config.startTimestamp,
		endTimestamp:   r.config.endTimestamp,
	}

	if err := task.LoadTokensFromFile(); err != nil {
		if err := task.LoadTokensWithPrices(); err != nil {
			return err
		}

		if err := task.SaveTokensToFile(); err != nil {
			return err
		}
	}

	r.tokenMap = task.tokenMap

	return nil
}

// LoadTxs load txs
func (r *TradingRewarder) LoadTxs() error {
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
func (r *TradingRewarder) LoadTradings() error {
	task := LoadTradingsTask{
		filepath:    path.Join(r.config.RoundDir(), "tradings.json"),
		txs:         r.txs,
		cubeFinders: r.config.cubeFinders,
		tokenMap:    r.tokenMap,
	}

	if err := task.LoadTradingsFromFile(); err != nil {
		if err := task.GetTradingsFromTxs(); err != nil {
			return err
		}

		task.RankTradings()

		if err := task.WeightTradings(); err != nil {
			return err
		}

		if err := task.SaveTradingsToFile(); err != nil {
			return err
		}
	}

	r.tradingMap = task.tradingMap

	return nil
}

// LoadRewards load rewards
func (r *TradingRewarder) LoadRewards() error {
	task := LoadRewardsTask{
		filepath:        path.Join(r.config.RoundDir(), "rewards.json"),
		rewardWeightMap: r.tradingMap.ToRewardWeightMap(),
		rewardAmount:    r.config.RewardAmount,
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
func (r *TradingRewarder) GenerateRewardsMerkleTree() error {
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
