package rewarder

import (
	"path"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// StakingRewarderRequiredFieldNames staking rewarder required field names
var StakingRewarderRequiredFieldNames = []string{"name", "blocks", "cubes", "pools"}

// StakingRewarder struct
type StakingRewarder struct {
	config *Config

	tokenMap    TokenMap
	txs         []common.Hash
	tradingMap  TradingMap
	stakingsMap map[common.Address]StakingMap
	rewardsMap  map[common.Address]RewardMap
}

// NewStakingRewarder new staking rewarder
func NewStakingRewarder(config *Config) *StakingRewarder {
	return &StakingRewarder{config: config}
}

// LoadTokens load tokens
func (r *StakingRewarder) LoadTokens() error {
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
func (r *StakingRewarder) LoadTxs() error {
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
func (r *StakingRewarder) LoadTradings() error {
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

		if err := task.SaveTradingsToFile(); err != nil {
			return err
		}
	}

	r.tradingMap = task.tradingMap

	return nil
}

// LoadStakings load stakings
func (r *StakingRewarder) LoadStakings() error {
	r.stakingsMap = make(map[common.Address]StakingMap)

	for _, pool := range r.config.Pools {
		task := LoadStakingsTask{
			filepath:    path.Join(r.config.RoundDir(), pool.Address.String(), "stakings.json"),
			round:       r.config.Round,
			duration:    decimal.NewFromInt(int64(r.config.Blocks())),
			tradingMap:  r.tradingMap,
			poolAddress: pool.Address,
			baseAmount:  pool.BaseAmount,
			startBlock:  r.config.StartBlock,
			endBlock:    r.config.EndBlock,
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
		}

		r.stakingsMap[pool.Address] = task.stakingMap
	}

	return nil
}

// LoadRewards load rewards
func (r *StakingRewarder) LoadRewards() error {
	r.rewardsMap = make(map[common.Address]RewardMap)

	for _, pool := range r.config.Pools {
		task := LoadRewardsTask{
			filepath:        path.Join(r.config.RoundDir(), pool.Address.String(), "rewards.json"),
			rewardWeightMap: r.stakingsMap[pool.Address].ToRewardWeightMap(),
			rewardAmount:    pool.RewardAmount,
		}

		if err := task.LoadRewardsFromFile(); err != nil {
			if err := task.CalcRewardsByWeight(); err != nil {
				return err
			}

			if err := task.SaveRewardsToFile(); err != nil {
				return err
			}
		}

		r.rewardsMap[pool.Address] = task.rewardMap
	}

	return nil
}

// GenerateRewardsMerkleTree generate rewards merkle tree
func (r *StakingRewarder) GenerateRewardsMerkleTree() error {
	for _, pool := range r.config.Pools {
		task := GenerateRewardMerkleTreeTask{
			rewardMap:                      r.rewardsMap[pool.Address],
			rewardMerkleTreeLeavesFilepath: path.Join(r.config.RoundDir(), pool.Address.String(), "merkle_tree_leaves.json"),
			rewardMerkleProofsFilepath:     path.Join(r.config.RoundDir(), pool.Address.String(), "merkle_proofs.json"),
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
	}

	return nil
}
