package ethereum

import (
	"math/big"

	"github.com/shopspring/decimal"
)

// ToBigUnit small unit to big unit
func ToBigUnit(v *big.Int, d uint8) decimal.Decimal {
	return decimal.NewFromBigInt(v, -int32(d))
}

// ToSmallUnit big unit to small unit
func ToSmallUnit(v decimal.Decimal, d uint8) decimal.Decimal {
	return v.Mul(decimal.New(1, int32(d)))
}
