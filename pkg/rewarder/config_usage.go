package rewarder

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// UsageConfig usage config
type UsageConfig struct {
	Address    common.Address  `json:"address"`
	BaseReward decimal.Decimal `json:"baseReward"`
	MaxReward  decimal.Decimal `json:"maxReward"`
}
