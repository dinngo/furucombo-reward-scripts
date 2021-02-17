package rewarder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/coingecko"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
)

var coingeckoTokenIDMapFilepath = path.Join("config", "coingecko_token_ids.json")

// PriceOracle struct
type PriceOracle struct {
	rootpath string
	filepath string
	from     uint64
	to       uint64

	client      *coingecko.Client
	tokenIDMap  map[common.Address]string
	tokenPrices map[common.Address][]CoingeckoTokenPrice
}

// CoingeckoTokenPrice represents usd price at a certain point in time
type CoingeckoTokenPrice struct {
	Timestamp uint64          `json:"timestamp"`
	USD       decimal.Decimal `json:"usd"`
}

// NewPriceOracle new price oracle
func NewPriceOracle(rootpath string, from, to uint64) (*PriceOracle, error) {
	priceOracle := &PriceOracle{
		rootpath: rootpath,
		filepath: path.Join(rootpath, "prices.json"),
		from:     from,
		to:       to,
		client:   coingecko.NewClient(&http.Client{Timeout: 10 * time.Second}),
	}

	if err := priceOracle.LoadTokenIDs(); err != nil {
		return nil, err
	}

	if err := priceOracle.LoadPricesFromFile(); err != nil {
		priceOracle.tokenPrices = map[common.Address][]CoingeckoTokenPrice{}
	}

	return priceOracle, nil
}

// LoadPricesFromFile load prices from file
func (t *PriceOracle) LoadPricesFromFile() error {
	log.Printf("loading prices from ./%s", t.filepath)

	if _, err := os.Stat(t.filepath); err != nil {
		log.Println("prices file not found")
		return err
	}

	data, err := ioutil.ReadFile(t.filepath)
	if err != nil {
		return err
	}

	if err := jsonex.Unmarshal(data, &t.tokenPrices); err != nil {
		return err
	}

	return nil
}

// LoadTokenIDs load token ids
func (t *PriceOracle) LoadTokenIDs() error {
	if _, err := os.Stat(coingeckoTokenIDMapFilepath); err != nil {
		return err
	}

	data, err := ioutil.ReadFile(coingeckoTokenIDMapFilepath)
	if err != nil {
		return err
	}

	if err := jsonex.Unmarshal(data, &t.tokenIDMap); err != nil {
		return err
	}

	return nil
}

// GetTokenPrices get token prices
func (t *PriceOracle) GetTokenPrices(token *ethereum.Token) []CoingeckoTokenPrice {
	if _, ok := t.tokenPrices[token.Address]; !ok {
		var marketChart *coingecko.CoinsIDMarketChart
		var err error

		// Get coingecko price and retry
		for i := 0; i < 3; i++ {
			if tokenID, ok := t.tokenIDMap[token.Address]; ok {
				marketChart, err = t.client.CoinsIDMarketChartRange(tokenID, "usd", t.from, t.to)
				if err != nil {
					log.Printf("failed to get coingecko id: %s", tokenID)
					continue
				}
			} else {
				marketChart, err = t.client.CoinsContractMarketChartRange(token.Address.String(), "usd", t.from, t.to)
				if err != nil {
					log.Printf("failed to get coingecko address: %s", token.Address.String())
					continue
				}
			}
			break
		}

		if err != nil {
			t.tokenPrices[token.Address] = make([]CoingeckoTokenPrice, 0)
		} else {
			pricesLen := len(marketChart.Prices)
			t.tokenPrices[token.Address] = make([]CoingeckoTokenPrice, pricesLen)

			if pricesLen > 0 {
				for i, chartItem := range marketChart.Prices {
					t.tokenPrices[token.Address][i] = CoingeckoTokenPrice{
						Timestamp: uint64(chartItem[0]),
						USD:       decimal.NewFromFloat(chartItem[1]),
					}
				}

				log.Printf("found %s %d prices data", token.Symbol, pricesLen)
			} else {
				log.Printf("found %s %d prices data (address: %s)", token.Symbol, pricesLen, token.Address.String())
			}
		}
	}

	return t.tokenPrices[token.Address]
}

// GetClosestPrice returns usd price given a timestamp
func (t *PriceOracle) GetClosestPrice(token *ethereum.Token, timestamp uint64) decimal.Decimal {
	usd := decimal.Zero

	// timestamp to millisecond
	ts := float64(timestamp * 1000)

	// search closest usd price
	var closest float64
	for i, v := range t.GetTokenPrices(token) {
		current := float64(v.Timestamp)
		if i == 0 || math.Abs(current-ts) < math.Abs(closest-ts) {
			closest = current
			usd = v.USD
			continue
		}
	}

	return usd
}

// SavePricesToFile save prices to file
func (t *PriceOracle) SavePricesToFile() error {
	log.Printf("saving prices to ./%s", t.filepath)

	data, err := json.MarshalIndent(t.tokenPrices, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(t.filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}
