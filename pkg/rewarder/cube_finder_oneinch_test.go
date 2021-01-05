package rewarder

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"garage.dinngo.co/furucombo/mining-scripts/pkg/ethereum"
)

func TestFindOneInchSwapCube(t *testing.T) {
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
			txHash: common.HexToHash("0x80680260b4deb972e6288327ca4a6cf1b6fead25c8d5bf7eef557953f869e580"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        23,
					tokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					tokenAmount:  "1518107870817153810",
				},
			},
		},
		// yCRV
		{
			txHash: common.HexToHash("0x5ab53bdd9b708b778e634c5575ca68ce0f718dceaea752ac6e2a18e974795e8c"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        35,
					tokenAddress: "0xdF5e0e81Dff6FAF3A7e52BA697820c5e32D806A8",
					tokenAmount:  "14968887700686513609848",
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
				cube, err := findOneInchSwapCube(txLog)
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
