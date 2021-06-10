package rewarder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
)

// TokenPrices token prices
type TokenPrices struct {
	Combo float64 `json:"combo"`
	Matic float64 `json:"matic"`
}

// SaveTokenPrices save token prices to file
func (c *Config) SaveTokenPrices() error {
	filepath := path.Join(c.RoundDir(), "token_prices.json")

	if _, err := os.Stat(filepath); err == nil {
		return nil
	}

	log.Printf("saving token prices: ./%s", filepath)

	if err := c.GetTokenPrices(); err != nil {
		return err
	}

	data, err := json.MarshalIndent(c.tokenPrices, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}

// GetTokenPrices get token prices
func (c *Config) GetTokenPrices() error {
	to := time.Now()
	from := to.Add(-24 * time.Hour) // 1 day ago

	price, err := getTokenLatestPrice(ethereum.GetTokenAddress("COMBO"), from, to)
	if err != nil {
		return err
	}
	c.tokenPrices.Combo, _ = price.Float64()
	log.Printf("print the latest price %s: %s", ethereum.GetTokenAddress("COMBO").Hex(), price.String())

	price, err = getTokenLatestPrice(ethereum.GetTokenAddress("MATIC"), from, to)
	if err != nil {
		return err
	}
	c.tokenPrices.Matic, _ = price.Float64()
	log.Printf("print the latest price %s: %s", ethereum.GetTokenAddress("MATIC").Hex(), price.String())

	return nil
}
