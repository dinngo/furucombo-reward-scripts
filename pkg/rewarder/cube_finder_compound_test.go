package rewarder

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"garage.dinngo.co/furucombo/mining-scripts/pkg/ethereum"
)

func TestFindCompoundSupplyCube(t *testing.T) {
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
		// cETH
		{
			txHash: common.HexToHash("0x1c417610bb191b5e037321a6cc3e5ccbc1587e3847fed65139ea2beca5a6e383"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        2,
					tokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					tokenAmount:  "222000000000000000000",
				},
			},
		},
		// cUSDC
		{
			txHash: common.HexToHash("0x69ce50e7dacd870e0e91dd448b8cb2272c7ead79f116e1c3f598d2888ea6d23f"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        21,
					tokenAddress: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
					tokenAmount:  "8256079890",
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
				cube, err := findCompoundSupplyCube(txLog)
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

func TestFindCompoundRepayCube(t *testing.T) {
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
		// cUSDC
		{
			txHash: common.HexToHash("0x6faa9eb08371ecfe7eb735e55d4ee2bd774172be72976513fdd0eb9c6a4658c5"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        7,
					tokenAddress: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
					tokenAmount:  "3227045000",
				},
			},
		},
		// cDAI
		{
			txHash: common.HexToHash("0xf859c2839bba7d61c4eb18878ddcfbdea087465d8aa78539bfd285e86327f23d"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        32,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "14069261646692017265170",
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
				cube, err := findCompoundRepayCube(txLog)
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
