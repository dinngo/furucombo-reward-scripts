package rewarder

import (
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/coingecko"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/uniswapv2"
)

// PrintStakingTokenPrice print staking token price
func (c *Config) PrintStakingTokenPrice() error {
	to := time.Now()
	from := to.Add(-24 * time.Hour) // 1 day ago

	for _, pool := range c.Pools {
		var price decimal.Decimal
		staking, err := furucombo.NewStakingContract(pool.Address, ethereum.Client())
		stakingToken, err := staking.StakingToken(new(bind.CallOpts))
		if err != nil {
			return err
		}

		price, err = getTokenLatestPrice(stakingToken, from, to)
		if err != nil {
			price, err = getPairLatestPrice(stakingToken, from, to)
			if err != nil {
				return err
			}
		}
		log.Printf("print the latest price %s: %s", stakingToken.Hex(), price.String())
	}

	return nil
}

func getTokenLatestPrice(address common.Address, from, to time.Time) (decimal.Decimal, error) {
	// get coingecko price
	client := coingecko.NewClient(&http.Client{Timeout: 10 * time.Second})
	tokenMarketChart, err := client.CoinsContractMarketChartRange(address.Hex(), "usd", uint64(from.Unix()), uint64(to.Unix()))
	if err != nil {
		return decimal.Zero, err
	}
	prices := tokenMarketChart.Prices
	price := decimal.NewFromFloat(prices[len(prices)-1][1]) // the latest price

	return price, nil
}

func getPairLatestPrice(address common.Address, from, to time.Time) (decimal.Decimal, error) {
	var err error

	// read uniswap v2 pair
	pair, err := uniswapv2.NewPairContract(address, ethereum.Client())
	if err != nil {
		return decimal.Zero, err
	}
	token0, err := pair.Token0(new(bind.CallOpts))
	if err != nil {
		return decimal.Zero, err
	}
	token1, err := pair.Token1(new(bind.CallOpts))
	if err != nil {
		return decimal.Zero, err
	}
	totalSupply, err := pair.TotalSupply(new(bind.CallOpts))
	if err != nil {
		return decimal.Zero, err
	}
	supply := decimal.NewFromBigInt(totalSupply, -18)
	reserves := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	*reserves, err = pair.GetReserves(new(bind.CallOpts))
	if err != nil {
		return decimal.Zero, err
	}
	reserve0 := decimal.NewFromBigInt(reserves.Reserve0, -18)
	reserve1 := decimal.NewFromBigInt(reserves.Reserve1, -18)

	token0Price, err := getTokenLatestPrice(token0, from, to)
	if err != nil {
		return decimal.Zero, err
	}
	token1Price, err := getTokenLatestPrice(token1, from, to)
	if err != nil {
		return decimal.Zero, err
	}
	pairPrice := reserve0.Mul(token0Price).Add(reserve1.Mul(token1Price)).Div(supply)

	return pairPrice, nil
}
