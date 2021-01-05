package rewarder

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// Staking represents user's staking base/start/end amount, area and weight
type Staking struct {
	BaseAmount  decimal.Decimal `json:"base_amount"` // pool_0 > 0, otherwise = 0
	StartAmount decimal.Decimal `json:"start_amount"`
	EndAmount   decimal.Decimal `json:"end_amount"`
	Area        decimal.Decimal `json:"area"`      // staking duration * amount
	RankArea    decimal.Decimal `json:"rank_area"` // area * rank
	Weight      decimal.Decimal `json:"weight"`
}

// GetWeight return staking weight
func (s *Staking) GetWeight() decimal.Decimal {
	return s.Weight
}

// StakingMap represents user's staking data
type StakingMap map[common.Address]*Staking

// Add add staking
func (m StakingMap) Add(account common.Address, duration decimal.Decimal, baseAmount decimal.Decimal, preStaking *Staking) {
	log.Printf("adding %s staking", account.String())

	startAmount := decimal.Zero
	endAmount := decimal.Zero
	if preStaking != nil {
		startAmount = preStaking.EndAmount
		endAmount = preStaking.EndAmount
	}
	area := baseAmount.Add(startAmount).Mul(duration)

	m[account] = &Staking{
		BaseAmount:  baseAmount,
		StartAmount: startAmount,
		EndAmount:   endAmount,
		Area:        area,
	}

	log.Printf("baseAmount: %s startAmount: %s endAmount: %s area: %s", baseAmount, startAmount, endAmount, area)
}

// ToRewardWeightMap to reward weight map
func (m StakingMap) ToRewardWeightMap() RewardWeightMap {
	rewardWeightMap := make(RewardWeightMap)
	for account, staking := range m {
		rewardWeightMap[account] = staking
	}

	return rewardWeightMap
}
