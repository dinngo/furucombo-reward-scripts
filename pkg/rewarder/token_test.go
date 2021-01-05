package rewarder

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestTokenGetClosestUSDPrice(t *testing.T) {
	// Set test cases
	testCases := []struct {
		token    Token
		expected []TokenPrice
	}{
		{
			token: Token{
				Prices: []TokenPrice{
					{
						Timestamp: 1606801683400,
						USD:       decimal.NewFromFloat(0.2312742896966513),
					},
					{
						Timestamp: 1606802100610,
						USD:       decimal.NewFromFloat(0.24310468641154304),
					},
					{
						Timestamp: 1606802415260,
						USD:       decimal.NewFromFloat(0.24351064282208018),
					},
					{
						Timestamp: 1606802812020,
						USD:       decimal.NewFromFloat(0.23682055793028398),
					},
					{
						Timestamp: 1606803161914,
						USD:       decimal.NewFromFloat(0.24366628540503518),
					},
				},
			},
			expected: []TokenPrice{
				{
					Timestamp: 1500000000,
					USD:       decimal.NewFromFloat(0.2312742896966513),
				},
				{
					Timestamp: 1606802200,
					USD:       decimal.NewFromFloat(0.24310468641154304),
				},
				{
					Timestamp: 1606802315,
					USD:       decimal.NewFromFloat(0.24351064282208018),
				},
				{
					Timestamp: 1606804161,
					USD:       decimal.NewFromFloat(0.24366628540503518),
				},
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			for _, expectedPrice := range testCase.expected {
				usdPrice := testCase.token.GetClosestUSDPrice(expectedPrice.Timestamp)
				assert.True(t, expectedPrice.USD.Equal(usdPrice))
			}
		})
	}
}

func TestTokenMapGetTokenPrices(t *testing.T) {
	// Set test cases
	testCases := []struct {
		tokenMap TokenMap
		from     uint64
		to       uint64
		expected map[common.Address][]TokenPrice
	}{
		{
			tokenMap: TokenMap{
				common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"): {
					Symbol:   "WETH",
					Decimals: 18,
				},
				common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"): {
					Symbol:   "USDC",
					Decimals: 6,
				},
				common.HexToAddress("0x674C6Ad92Fd080e4004b2312b45f796a192D27a0"): {
					Symbol:      "USDN",
					Decimals:    18,
					CoingeckoID: "neutrino",
				},
			},
			from: 1606752000,
			to:   1606753800,
			expected: map[common.Address][]TokenPrice{
				common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"): {
					{
						Timestamp: 1606753631724,
						USD:       decimal.NewFromFloat(602.3835770509737),
					},
				},
				common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"): {
					{
						Timestamp: 1606752511255,
						USD:       decimal.NewFromFloat(0.9966286659088316),
					},
				},
				common.HexToAddress("0x674C6Ad92Fd080e4004b2312b45f796a192D27a0"): {
					{
						Timestamp: 1606753221480,
						USD:       decimal.NewFromFloat(0.9974055194160021),
					},
				},
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			err := testCase.tokenMap.GetTokenPrices(testCase.from, testCase.to)
			assert.Nil(t, err)
			for tokenAddress, token := range testCase.tokenMap {
				assert.True(t, reflect.DeepEqual(testCase.expected[tokenAddress], token.Prices))
			}
		})
	}
}
