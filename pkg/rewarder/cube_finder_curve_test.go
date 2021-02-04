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
		// seth + susd pool - eth to seth and usdc to susd
		{
			txHash: common.HexToHash("0x54c19a399bc13e5c6f75b6414ebd441c68d00d1159b809bd029e7cd5968b0a63"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        2,
					tokenAddress: "0x5e74C9036fb86BD7eCdcb084a0673EFc32eA31cb",
					tokenAmount:  "1006011043273916126",
				},
				{
					index:        5,
					tokenAddress: "0x57Ab1ec28D129707052df4dF418D58a2D46d5f51",
					tokenAmount:  "1469986114991477575443",
				},
			},
		},
		// steth pool - steth to eth
		{
			txHash: common.HexToHash("0xa23ce0068cc7535dd3be084ed16878024b4cf1103f2f953034b261ab64c28b90"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        5,
					tokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					tokenAmount:  "99920019748817702467",
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
