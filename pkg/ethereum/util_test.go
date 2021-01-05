package ethereum

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestToBigUnit(t *testing.T) {
	testCases := []struct {
		amount   string
		decimals int32
		expected string
	}{
		{
			amount:   "1234567890012340567",
			decimals: 18,
			expected: "1.234567890012340567",
		},
		{
			amount:   "123456789001234",
			decimals: 18,
			expected: "0.000123456789001234",
		},
		{
			amount:   "12345678901",
			decimals: 8,
			expected: "123.45678901",
		},
		{
			amount:   "123456789",
			decimals: 8,
			expected: "1.23456789",
		},
		{
			amount:   "123456789",
			decimals: 6,
			expected: "123.456789",
		},
		{
			amount:   "12345",
			decimals: 2,
			expected: "123.45",
		},
		{
			amount:   "123",
			decimals: 0,
			expected: "123",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			sAmount, _ := new(big.Int).SetString(testCase.amount, 10)
			amount := ToBigUnit(sAmount, testCase.decimals)
			assert.Equal(t, testCase.expected, amount.String())
		})
	}
}

func TestToSmallUnit(t *testing.T) {
	testCases := []struct {
		amount   string
		decimals int32
		expected string
	}{
		{
			amount:   "1.234567890012340567",
			decimals: 18,
			expected: "1234567890012340567",
		},
		{
			amount:   "0.000123456789001234",
			decimals: 18,
			expected: "123456789001234",
		},
		{
			amount:   "123.45678901",
			decimals: 8,
			expected: "12345678901",
		},
		{
			amount:   "1.23456789",
			decimals: 8,
			expected: "123456789",
		},
		{
			amount:   "123.456789",
			decimals: 6,
			expected: "123456789",
		},
		{
			amount:   "123.45",
			decimals: 2,
			expected: "12345",
		},
		{
			amount:   "123",
			decimals: 0,
			expected: "123",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			amount, err := decimal.NewFromString(testCase.amount)
			assert.Nil(t, err)
			sAmount := ToSmallUnit(amount, testCase.decimals)
			assert.Equal(t, testCase.expected, sAmount.String())
		})
	}
}
