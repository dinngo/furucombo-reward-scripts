package rewarder

import (
	"log"
	"os"
	"path"

	"github.com/ethereum/go-ethereum/common"
)

// PolygonUsageRewarderRequiredFieldNames polygon usage rewarder required field names
var PolygonUsageRewarderRequiredFieldNames = []string{"name", "endBlock", "polygonStartBlock", "polygonEndBlock", "polygonBlocks", "maxLoyalTxCount", "nft"}

// PolygonUsageRewarder struct
type PolygonUsageRewarder struct {
	config *Config

	bridgeTxMap       BridgeTxMap
	tradingLoyaltyMap TradingLoyaltyMap
	rewardsMap        map[common.Address]RewardMap
}

// NewPolygonUsageRewarder new polygon usage rewarder
func NewPolygonUsageRewarder(config *Config) *PolygonUsageRewarder {
	return &PolygonUsageRewarder{config: config}
}

// LoadBridgeTxs load user's valid bridge txs
func (r *PolygonUsageRewarder) LoadBridgeTxs() error {
	task := LoadBridgeTxsTask{
		rootpath: r.config.RoundDir(),
		endBlock: r.config.EndBlock,
	}

	if err := task.Execute(); err != nil {
		return err
	}

	r.bridgeTxMap = task.bridgeTxMap

	return nil
}

// LoadTradingLoyalty load trading loyalty
func (r *PolygonUsageRewarder) LoadTradingLoyalty() error {
	task := LoadTradingLoyaltyTask{
		rootpath:        r.config.RoundDir(),
		startBlock:      r.config.PolygonStartBlock,
		endBlock:        r.config.PolygonEndBlock,
		maxLoyalTxCount: r.config.MaxLoyalTxCount,
		nft:             r.config.Nft,
		bridgeTxMap:     r.bridgeTxMap,
	}

	if err := task.Execute(); err != nil {
		return err
	}

	r.tradingLoyaltyMap = task.tradingLoyaltyMap

	return nil
}

// MakeUsageDir make usage dir
func (r *PolygonUsageRewarder) MakeUsageDir() error {
	for _, usage := range r.config.Usages {
		usageDir := path.Join(r.config.RoundDir(), usage.Address.String())

		log.Printf("making usage dir: ./%s/", usageDir)

		if err := os.MkdirAll(usageDir, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

// LoadRewards load rewards
func (r *PolygonUsageRewarder) LoadRewards() error {
	r.rewardsMap = make(map[common.Address]RewardMap)

	for _, usage := range r.config.Usages {
		task := LoadLoyaltyRewardsTask{
			rootpath:          path.Join(r.config.RoundDir(), usage.Address.String()),
			baseReward:        usage.BaseReward,
			maxReward:         usage.MaxReward,
			tradingLoyaltyMap: r.tradingLoyaltyMap,
		}

		if err := task.Execute(); err != nil {
			return err
		}

		r.rewardsMap[usage.Address] = task.rewardMap
	}

	return nil
}

// GenerateRewardsMerkleTree generate rewards merkle tree
func (r *PolygonUsageRewarder) GenerateRewardsMerkleTree() error {
	for _, usage := range r.config.Usages {
		task := GenerateRewardMerkleTreeTask{
			rootpath:  path.Join(r.config.RoundDir(), usage.Address.String()),
			rewardMap: r.rewardsMap[usage.Address],
		}

		if err := task.Execute(); err != nil {
			return err
		}
	}

	return nil
}
