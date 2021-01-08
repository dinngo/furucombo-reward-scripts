package rewarder

import (
	"errors"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/uniswapv2"
)

func init() {
	cubeFinderMap["UniswapV2Swap"] = findUniswapV2SwapCube
}

func isUniswapV2SwapCube(txLog *types.Log) bool {
	// 1. check is uniswap v2 swap event or not
	// 2. check is uniswap v2 router 02 contract or not
	if uniswapv2.IsSwapEvent(txLog.Topics[0]) && uniswapv2.IsRouter02Contract(txLog.Topics[1]) {

		// 3-1. check is (ETH -> Token) and (Token -> Token) or not
		if furucombo.IsProxy(txLog.Topics[2]) {
			return true
		}

		// 3-2. check is (Token -> ETH) or not
		if uniswapv2.IsRouter02Contract(txLog.Topics[2]) {
			return true
		}
	}

	return false
}

func findUniswapV2SwapCube(txLog *types.Log) (*Cube, error) {
	if !isUniswapV2SwapCube(txLog) {
		return nil, nil
	}

	// get pair's token0, token1
	pair, err := uniswapv2.NewPairContract(txLog.Address, ethereum.Client())
	if err != nil {
		return nil, err
	}
	token0, err := pair.Token0(new(bind.CallOpts))
	if err != nil {
		return nil, err
	}
	token1, err := pair.Token1(new(bind.CallOpts))
	if err != nil {
		return nil, err
	}

	swap, err := pair.ParseSwap(*txLog)
	if err != nil {
		return nil, err
	}

	cube := Cube{
		Name: "UniswapV2 Swap",
	}

	// check token amount out
	if swap.Amount0Out.Cmp(new(big.Int)) == 1 {
		// check token is listed or not
		if !IsTokenListed(token0) {
			log.Printf("UniswapV2 Swap: %s is not listed", token0.String())
			return nil, nil
		}

		cube.TokenAddress = token0
		cube.TokenAmount = swap.Amount0Out
	} else if swap.Amount1Out.Cmp(new(big.Int)) == 1 {
		// check token is listed or not
		if !IsTokenListed(token1) {
			log.Printf("UniswapV2 Swap: %s is not listed", token0.String())
			return nil, nil
		}

		cube.TokenAddress = token1
		cube.TokenAmount = swap.Amount1Out
	} else {
		return nil, errors.New("UniswapV2 Swap: unexpected 0 amount out")
	}

	return &cube, nil
}
