package rewarder

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// TradingGasComboMap represents user's trading gas reimbursement in combo map
type TradingGasComboMap map[common.Address]decimal.Decimal

type TradingGasComboParams struct {
	GasPrice   decimal.Decimal `json:"gas_price"`
	ComboPrice decimal.Decimal `json:"combo_price"`
}

// GasUsedMap represents user's gas used map
type GasUsedMap map[common.Address]decimal.Decimal
