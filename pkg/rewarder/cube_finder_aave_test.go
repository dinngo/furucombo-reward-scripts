package rewarder

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
)

func TestFindAaveDepositCube(t *testing.T) {
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
		// aETH
		{
			txHash: common.HexToHash("0xcc22166408ea4950e27e1c3f08e3c8f20638fa9d3eb8dad6f63013e167b82878"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        3,
					tokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					tokenAmount:  "40000000000000000000",
				},
			},
		},
		// aDAI
		{
			txHash: common.HexToHash("0xffdfe36dfc790a96bf7f42eae421ea6db5dca74203a4354b623d2996325e504c"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        16,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "2457521533496077815",
				},
			},
		},
		// aUNI
		{
			txHash: common.HexToHash("0xcfab62f0ff6535a6cb4f36ece4ae08de723a799b59019723a903d28bcd2896f3"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        14,
					tokenAddress: "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984",
					tokenAmount:  "3392720433781194898367",
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
				cube, err := findAaveDepositCube(txLog)
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
