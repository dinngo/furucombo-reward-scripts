package rewarder

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"garage.dinngo.co/furucombo/mining-scripts/pkg/ethereum"
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
			txHash: common.HexToHash("0xf5b0bb3fed051c8e3ba6841d2fe55da6b4f11a6d9dc80d660d692f9bbd55124f"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        5,
					tokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					tokenAmount:  "2500000000000000000",
				},
				{
					index:        12,
					tokenAddress: "0x6B3595068778DD592e39A122f4f5a5cF09C90fE2",
					tokenAmount:  "112775964742293072453",
				},
			},
		},
		// Token -> ETH
		{
			txHash: common.HexToHash("0xd2f4505025bdc4a36e9ddd7f064c90d76b48139b2212f665520c1436a9ad37ce"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        5,
					tokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					tokenAmount:  "20153251452157105001",
				},
			},
		},
		// Token -> ETH
		{
			txHash: common.HexToHash("0xac22e49a072dac4f6a92b1e7ebbbbe2440385c623c20fab2e2cf4f469f3e48b0"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        7,
					tokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					tokenAmount:  "1174013744947632149",
				},
			},
		},
		// Token -> Token
		{
			txHash: common.HexToHash("0xefe29c9885e54f73316596c9a8f99f0b188816d618f1afeb18bb57db4b23c21b"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        18,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "1754363292667535728136",
				},
				{
					index:        27,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "3750228056898634734263",
				},
				{
					index:        36,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "1900812689938819392076",
				},
				{
					index:        42,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "2703917976415005639723",
				},
				{
					index:        51,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "2323737161187031097926",
				},
				{
					index:        61,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "5707097150159011893807",
				},
				{
					index:        71,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "1770127275414896103851",
				},
				{
					index:        81,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "1688187348731527753278",
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
