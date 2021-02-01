package rewarder

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
)

func TestFindCurveSwapCube(t *testing.T) {
	TestLoadDotEnv(t)
	TestResetTokenWhitelistFilepath(t)

	testCases := []struct {
		txHash   common.Hash
		expected []struct {
			index        int
			tokenAddress string
			tokenAmount  string
		}
	}{
		// single pool - renBTC
		{
			txHash: common.HexToHash("0x8b08e181b39544cf6a3abc76c6a36a0fec1079eb0ddd40bee813e25e961aec05"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        10,
					tokenAddress: "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
					tokenAmount:  "46562906",
				},
			},
		},
		// single pool - stable coin
		{
			txHash: common.HexToHash("0xae21c1ef3c37e872454f0da5df798288c3c30cf0c48ed94f4841621849d43c5c"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        3,
					tokenAddress: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
					tokenAmount:  "168054568316",
				},
				{
					index:        8,
					tokenAddress: "0xdAC17F958D2ee523a2206206994597C13D831ec7",
					tokenAmount:  "167943212299",
				},
				{
					index:        13,
					tokenAddress: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
					tokenAmount:  "167890246543",
				},
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			receipt, err := ethereum.Client().TransactionReceipt(context.Background(), testCase.txHash)
			assert.Nil(t, err)

			count := 0
			for i, txLog := range receipt.Logs {
				cube, err := findCurveSwapCube(txLog)
				assert.Nil(t, err)

				if cube != nil {
					assert.Equal(t, testCase.expected[count].index, i)
					assert.Equal(t, testCase.expected[count].tokenAddress, cube.TokenAddress.String())
					assert.Equal(t, testCase.expected[count].tokenAmount, cube.TokenAmount.String())
					count++
				}
			}
			assert.Equal(t, len(testCase.expected), count)
		})
	}
}
