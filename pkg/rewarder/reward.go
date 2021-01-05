package rewarder

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// RewardMap represents user's reward amount
type RewardMap map[common.Address]decimal.Decimal

// RewardWeight interface
type RewardWeight interface {
	GetWeight() decimal.Decimal
}

// RewardWeightMap represents user's weight of reward amount
type RewardWeightMap map[common.Address]RewardWeight
