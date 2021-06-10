package rewarder

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// Loyalty represents user's loyalty data
type Loyalty struct {
	NftBoost         decimal.Decimal `json:"nftBoost"`
	FurucomboTxCount decimal.Decimal `json:"furucomboTxCount"`
	OfficialTxCount  decimal.Decimal `json:"officialTxCount"`
	TxCount          decimal.Decimal `json:"txCount"`
	LoyalTxCount     decimal.Decimal `json:"loyalTxCount"`
	Loyalty          decimal.Decimal `json:"loyalty"`
}

// TradingLoyaltyMap represents user's loyalty data
type TradingLoyaltyMap map[common.Address]*Loyalty
