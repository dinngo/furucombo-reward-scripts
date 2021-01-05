package rewarder

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"garage.dinngo.co/furucombo/mining-scripts/pkg/ethereum"
)

func TestFindCurveSwapCube(t *testing.T) {
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
		// single pool - renBTC
		{
			txHash: common.HexToHash("0x0f6c546ff421d17074aaef7db15ad15fcee373221da4cd2716ba9faef10b34e3"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        3,
					tokenAddress: "0xEB4C2781e4ebA804CE9a9803C67d0893436bB27D",
					tokenAmount:  "100130088",
				},
			},
		},
		// single pool - stable coin
		{
			txHash: common.HexToHash("0xbd72335f7800897a3fb1c2aeb3f2364f0bbd8adf83e6e9bd5c62f663ac8f4114"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        33,
					tokenAddress: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
					tokenAmount:  "2036399759",
				},
				{
					index:        38,
					tokenAddress: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
					tokenAmount:  "5520403",
				},
			},
		},
		// one split
		{
			txHash: common.HexToHash("0x1fd8df4d244e87b6efdff2c941fe578b698f1a34e87265446b5d4dc6f0d872af"),
			expected: []struct {
				index        int
				tokenAddress string
				tokenAmount  string
			}{
				{
					index:        4,
					tokenAddress: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
					tokenAmount:  "645736104",
				},
				{
					index:        14,
					tokenAddress: "0x0000000000085d4780B73119b644AE5ecd22b376",
					tokenAmount:  "645663045659510433432",
				},
				{
					index:        24,
					tokenAddress: "0x4Fabb145d64652a948d72533023f6E7A623C7C53",
					tokenAmount:  "371128605362290943981",
				},
				{
					index:        41,
					tokenAddress: "0x8E870D67F660D95d5be530380D0eC0bd388289E1",
					tokenAmount:  "648453222569354140961",
				},
				{
					index:        45,
					tokenAddress: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
					tokenAmount:  "646732943",
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
