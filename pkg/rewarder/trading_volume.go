package rewarder

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// TradingVolumeMap represents address to trading volume map
type TradingVolumeMap map[common.Address]decimal.Decimal
