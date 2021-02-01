package rewarder

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
)

func TestFindUniswapV2SwapCube(t *testing.T) {
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
		// ETH -> Token + Token -> ETH
		{
			txHash: common.HexToHash("0x9ce9f0e2edb9713e532adcec3cc30d8d5af2f05e33a2d233301a926c630eede6"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        7,
					tokenAddress: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
					tokenAmount:  "2333924258",
				},
				{
					index:        12,
					tokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					tokenAmount:  "1985163242248875307",
				},
			},
		},
		// Token -> ETH
		{
			txHash: common.HexToHash("0x1262fc4bb26c0c0b3e99903a6722d1da945bc380ffeedf302c099b816bff8d95"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        9,
					tokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					tokenAmount:  "108498737630701318",
				},
			},
		},
		// Token -> ETH
		{
			txHash: common.HexToHash("0x92aad6a7651fe0edfd7b6777e6671c2e7521d86aeee127268c7fee48e5e6acd6"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        19,
					tokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					tokenAmount:  "993845503594863868",
				},
			},
		},
		// Token -> Token
		{
			txHash: common.HexToHash("0x482f5fce99fc7889a095c7d05912b7a2f3790c7d5e4fe55e144674bf66652162"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        15,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "402705740484498616789",
				},
				{
					index:        24,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "246413094375332222540",
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
				cube, err := findUniswapV2SwapCube(txLog)
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
