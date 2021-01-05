package rewarder

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"garage.dinngo.co/furucombo/mining-scripts/pkg/ethereum"
)

func TestFindYearnDepositCube(t *testing.T) {
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
		// 3poolCRV
		{
			txHash: common.HexToHash("0xa1f3bc3c7cdec553ec4d13ddccc2d87da317abca9beefdde59e605d068b440b5"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        8,
					tokenAddress: "0x6c3F90f043a72FA612cbac8115EE7e52BDe6E490",
					tokenAmount:  "985228734648762369335",
				},
			},
		},
		// DAI
		{
			txHash: common.HexToHash("0x6fdf9893f1f02f560f69187ab44e71a54ce97594b6a7060382a7b18b6fcec007"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        19,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "1000000000000000000000",
				},
			},
		},
		// yCRV and DAI
		{
			txHash: common.HexToHash("0x33ec3b2dda4bc6d14742f7e8400204b0f5186f060b58d345708fa13b24181456"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        22,
					tokenAddress: "0xdF5e0e81Dff6FAF3A7e52BA697820c5e32D806A8",
					tokenAmount:  "190000000000000000000",
				},
				{
					index:        44,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "195000000000000000000",
				},
			},
		},
		// ETH
		{
			txHash: common.HexToHash("0x5cc32b6c4bb423d5faee11a2079690b8fe6bf79fce63d9594d0b8311d8f01e8f"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        0,
					tokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					tokenAmount:  "1000000000000000",
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
				cube, err := findYearnDepositCube(txLog)
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
