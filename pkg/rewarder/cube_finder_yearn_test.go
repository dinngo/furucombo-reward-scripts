package rewarder

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
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
			txHash: common.HexToHash("0xf4d11d692077670d641fa2d69c68047f1bfc011942c3917d6c64308194e27583"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        7,
					tokenAddress: "0x6c3F90f043a72FA612cbac8115EE7e52BDe6E490",
					tokenAmount:  "260151400000000000000",
				},
			},
		},
		// DAI
		{
			txHash: common.HexToHash("0x9724e78e958b5bda16e46e37085c43f73b00dcb1b078d5751985b08eff18d2d4"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        11,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "2500000000000000000000",
				},
			},
		},
		// yCRV
		{
			txHash: common.HexToHash("0xafd4a7d7556b2598c4cfc4858a83a6cf70ec99931f066c48ac216f6196d76b43"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        24,
					tokenAddress: "0xdF5e0e81Dff6FAF3A7e52BA697820c5e32D806A8",
					tokenAmount:  "372956772091292222946",
				},
			},
		},
		// ETH
		{
			txHash: common.HexToHash("0x8df0d0015c895a83d53789188e6ffe96ce0669e4a2c039715f0286d7f9692993"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        0,
					tokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					tokenAmount:  "200000000000000000000",
				},
				{
					index:        4,
					tokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					tokenAmount:  "200000000000000000000",
				},
				{
					index:        8,
					tokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					tokenAmount:  "200000000000000000000",
				},
				{
					index:        12,
					tokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					tokenAmount:  "200000000000000000000",
				},
				{
					index:        16,
					tokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					tokenAmount:  "200000000000000000000",
				},
				{
					index:        20,
					tokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					tokenAmount:  "200000000000000000000",
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
