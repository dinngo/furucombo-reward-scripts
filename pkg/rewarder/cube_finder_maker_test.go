package rewarder

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"garage.dinngo.co/furucombo/mining-scripts/pkg/ethereum"
)

func TestFindMakerDepositCube(t *testing.T) {
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
		// ETH
		{
			txHash: common.HexToHash("0xd2f4505025bdc4a36e9ddd7f064c90d76b48139b2212f665520c1436a9ad37ce"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        12,
					tokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					tokenAmount:  "20000000000000000000",
				},
			},
		},
		{
			txHash: common.HexToHash("0x274df7c06b0c964b13ee2e10e34cfb29be6d901c67b6cd33b05419b97e8030a8"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        6,
					tokenAddress: "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
					tokenAmount:  "5900000000",
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
				cube, err := findMakerDepositCube(txLog)
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

func TestFindMakerGenerateCube(t *testing.T) {
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
		{
			txHash: common.HexToHash("0xd2f4505025bdc4a36e9ddd7f064c90d76b48139b2212f665520c1436a9ad37ce"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        23,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "7887000000000000000000",
				},
			},
		},
		{
			txHash: common.HexToHash("0x274df7c06b0c964b13ee2e10e34cfb29be6d901c67b6cd33b05419b97e8030a8"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        18,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "340000000000000000000000",
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
				cube, err := findMakerGenerateCube(txLog)
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

func TestFindMakerPayBackCube(t *testing.T) {
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
		{
			txHash: common.HexToHash("0x4bfa3b488aedf8ae78fbc3b1f2641ef964640b642860626815ec7fa7e76bae6e"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        6,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "1100000000000000000000",
				},
			},
		},
		{
			txHash: common.HexToHash("0x042709fe9631f43fef00ef8ac197a2ef81a7b04d7fc2021f14ad6c508460460c"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        7,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "872059994167197600000",
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
				cube, err := findMakerPayBackCube(txLog)
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
