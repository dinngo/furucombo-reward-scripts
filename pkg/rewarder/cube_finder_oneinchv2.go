package rewarder

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/oneinchv2"
)

func init() {
	cubeFinderMap["OneInchV2Swap"] = findOneInchV2SwapCube
}

func isOneInchV2SwapCube(txLog *types.Log) bool {
	// 1. check is 1inch v2 Swapped event or not
	// 2. check is to furucombo proxy or not
	if oneinchv2.IsSwappedEvent(txLog.Topics[0]) && furucombo.IsProxy(txLog.Topics[1]) {
		return true
	}

	return false
}

func findOneInchV2SwapCube(txLog *types.Log) (*Cube, error) {
	if !isOneInchV2SwapCube(txLog) {
		return nil, nil
	}

	contractABI, err := abi.JSON(strings.NewReader(oneinchv2.ExchangeContractABI))
	if err != nil {
		return nil, err
	}

	event := new(oneinchv2.ExchangeContractSwapped)
	if err := contractABI.UnpackIntoInterface(event, "Swapped", txLog.Data); err != nil {
		return nil, err
	}

	// replace eETH with WETH
	tokenAddress := common.HexToAddress(txLog.Topics[3].Hex()) // DstToken
	if ethereum.IsToken("eETH", tokenAddress) {
		tokenAddress = ethereum.GetTokenAddress("WETH")
	}

	cube := Cube{
		Name:         "OneInchV2 Swap",
		TokenAddress: tokenAddress,
		TokenAmount:  event.ReturnAmount,
	}

	return &cube, nil
}
