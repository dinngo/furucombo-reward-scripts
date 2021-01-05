package rewarder

import (
	"fmt"
	"path"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestResetTokenWhitelistFilepath(t *testing.T) {
	tokenWhitelistFilepath = path.Join("..", "..", "config", "token_whitelist.json")
}

func TestGetTokenWhitelist(t *testing.T) {
	TestResetTokenWhitelistFilepath(t)
	tokenMap, err := GetTokenWhitelist()
	assert.Nil(t, err)
	assert.Greater(t, len(tokenMap), 0)
}

func TestIsTokenListed(t *testing.T) {
	TestResetTokenWhitelistFilepath(t)

	testCases := []struct {
		tokenAddress common.Address
		expected     bool
	}{
		{
			tokenAddress: common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"),
			expected:     true,
		},
		{
			tokenAddress: common.HexToAddress("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"),
			expected:     true,
		},
		{
			tokenAddress: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
			expected:     true,
		},
		{
			tokenAddress: common.HexToAddress("0x6d38574be6C230272DAAd16Fa5F291F825Bd0Da1"),
			expected:     false,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			ok := IsTokenListed(testCase.tokenAddress)
			assert.Equal(t, testCase.expected, ok)
		})
	}
}
