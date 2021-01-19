package rewarder

import (
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/oneinch"
)

func init() {
	cubeFinderMap["OneInchSwap"] = findOneInchSwapCube
}

func isOneInchSwapCube(txLog *types.Log) bool {
	// 1. check is 1inch History event or not
	// 2. check is to furucombo proxy or not
	if oneinch.IsHistoryEvent(txLog.Topics[0]) && furucombo.IsProxy(txLog.Topics[1]) {
		return true
	}

	return false
}

func findOneInchSwapCube(txLog *types.Log) (*Cube, error) {
	if !isOneInchSwapCube(txLog) {
		return nil, nil
	}

	contractABI, err := abi.JSON(strings.NewReader(oneinch.ExchangeContractABI))
	if err != nil {
		return nil, err
	}

	event := new(oneinch.ExchangeContractHistory)
	if err := contractABI.UnpackIntoInterface(event, "History", txLog.Data); err != nil {
		return nil, err
	}

	// replace eETH with WETH
	tokenAddress := event.ToToken
	if ethereum.IsToken("eETH", tokenAddress) {
		tokenAddress = ethereum.GetToken("WETH")
	}

	// check token is listed or not
	if !IsTokenListed(tokenAddress) {
		log.Printf("OneInch Swap: %s is not listed", tokenAddress.String())
		return nil, nil
	}

	cube := Cube{
		Name:         "OneInch Swap",
		TokenAddress: tokenAddress,
		TokenAmount:  event.ToAmount,
	}

	return &cube, nil
}
