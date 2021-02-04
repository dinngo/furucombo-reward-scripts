package ethereum

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestGetToken(t *testing.T) {
	TestLoadDotEnv(t)

	testCases := []struct {
		tokenAddress common.Address
		expected     Token
	}{
		{
			tokenAddress: common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"),
			expected: Token{
				Address:  common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"),
				Decimals: 18,
				Symbol:   "DAI",
			},
		},
		{
			tokenAddress: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
			expected: Token{
				Address:  common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
				Decimals: 6,
				Symbol:   "USDC",
			},
		},
		{
			tokenAddress: common.HexToAddress("0x0000000000004946c0e9F43F4Dee607b0eF1fA1c"),
			expected: Token{
				Address:  common.HexToAddress("0x0000000000004946c0e9F43F4Dee607b0eF1fA1c"),
				Decimals: 0,
				Symbol:   "CHI",
			},
		},
		{
			tokenAddress: common.HexToAddress("0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2"),
			expected: Token{
				Address:  common.HexToAddress("0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2"),
				Decimals: 18,
				Symbol:   "UNKNOWN",
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			tokenInfo, err := GetToken(testCase.tokenAddress)
			assert.Nil(t, err)

			assert.Equal(t, testCase.expected.Address, tokenInfo.Address)
			assert.Equal(t, testCase.expected.Decimals, tokenInfo.Decimals)
			assert.Equal(t, testCase.expected.Symbol, tokenInfo.Symbol)
		})
	}
}
