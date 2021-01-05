package rewarder

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"garage.dinngo.co/furucombo/mining-scripts/pkg/ethereum"
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
			txHash: common.HexToHash("0x61aa3180aca3df1cd85bfb2044ef5bd9462afd8cb3f223f2a6e0f94e868d2a80"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        11,
					tokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					tokenAmount:  "1442957397805663216",
				},
			},
		},
		// aDAI
		{
			txHash: common.HexToHash("0x95488f3c6cfed23f1e45eea0fca7ccfd3af2e5febf91b8964e814ea65856afb0"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        9,
					tokenAddress: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					tokenAmount:  "20000000000000000000",
				},
			},
		},
		// aYFI
		{
			txHash: common.HexToHash("0x234d68dcd6e90b75e1f280d0d42945ea95747a4009180033b068234b4967cdea"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        11,
					tokenAddress: "0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e",
					tokenAmount:  "586715522625526572",
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
