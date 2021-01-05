package rewarder

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// Trading represents trading volume, rank for each trading
type Trading struct {
	Volume decimal.Decimal  `json:"volume"`
	Rank   int              `json:"rank"`
	Weight *decimal.Decimal `json:"weight,omitempty"`
}

// GetWeight return trading weight
func (s *Trading) GetWeight() decimal.Decimal {
	return *s.Weight
}

// TradingMap represents address to trading map
type TradingMap map[common.Address]*Trading

// ToRewardWeightMap to reward weight map
func (m TradingMap) ToRewardWeightMap() RewardWeightMap {
	rewardWeightMap := make(RewardWeightMap)
	for account, trading := range m {
		rewardWeightMap[account] = trading
	}

	return rewardWeightMap
}
