package rewarder

import (
	"log"
	"math"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/coingecko"
)

// Token represents token data
type Token struct {
	Decimals    int32        `json:"decimals"`
	Symbol      string       `json:"symbol"`
	Prices      []TokenPrice `json:"prices"`
	CoingeckoID string       `json:"coingecko_id,omitempty"`
}

// TokenPrice represents usd price at a certain point in time
type TokenPrice struct {
	Timestamp uint64          `json:"timestamp"`
	USD       decimal.Decimal `json:"usd"`
}

// GetClosestUSDPrice returns usd price given a timestamp
func (t *Token) GetClosestUSDPrice(timestamp uint64) decimal.Decimal {
	// timestamp to millisecond
	ts := float64(timestamp * 1000)

	// search closest usd price
	var closest float64
	var usd decimal.Decimal
	for i, v := range t.Prices {
		current := float64(v.Timestamp)
		if i == 0 || math.Abs(current-ts) < math.Abs(closest-ts) {
			closest = current
			usd = v.USD
			continue
		}
	}

	return usd
}

// TokenMap represents token data map with address key
type TokenMap map[common.Address]*Token

// GetTokenPrices get token prices by timestamp range
func (m TokenMap) GetTokenPrices(from, to uint64) (err error) {
	log.Printf("getting %d tokens coingecko prices data from %d to %d", len(m), from, to)

	client := coingecko.NewClient(&http.Client{Timeout: 10 * time.Second})

	for tokenAddress, token := range m {
		sTokenAddress := tokenAddress.String()

		var marketChart *coingecko.CoinsIDMarketChart
		if len(token.CoingeckoID) > 0 {
			marketChart, err = client.CoinsIDMarketChartRange(token.CoingeckoID, "usd", from, to)
			if err != nil {
				return err
			}
		} else {
			marketChart, err = client.CoinsContractMarketChartRange(sTokenAddress, "usd", from, to)
			if err != nil {
				return err
			}
		}

		pricesLen := len(marketChart.Prices)
		prices := make([]TokenPrice, pricesLen)

		if pricesLen > 0 {
			for i, chartItem := range marketChart.Prices {
				prices[i] = TokenPrice{
					Timestamp: uint64(chartItem[0]),
					USD:       decimal.NewFromFloat(chartItem[1]),
				}
			}
			log.Printf("found %s %d prices data", token.Symbol, pricesLen)
		} else {
			log.Printf("found %s %d prices data (address: %s)", token.Symbol, pricesLen, sTokenAddress)
		}

		m[tokenAddress].Prices = prices
	}

	return nil
}
