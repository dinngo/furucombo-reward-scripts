package rewarder

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/curve"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/erc20"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
)

func init() {
	cubeFinderMap["CurveSwap"] = findCurveSwapCube
}

func isCurveSwapCube(txLog *types.Log) bool {
	// 1. check is erc20 Transfer event or not
	// 2. check is swap / one split contract or not
	// 3. check is to furucombo proxy or not
	if erc20.IsTransferEvent(txLog.Topics[0]) &&
		(curve.IsSwapContract(txLog.Topics[1]) || curve.IsOneSplitContract(txLog.Topics[1])) &&
		furucombo.IsProxy(txLog.Topics[2]) {
		return true
	}

	return false
}

func isCurveEthSwapCube(txLog *types.Log) bool {
	// for eth pool there is no erc20 transfer event if output token is eth
	// so decode token exchange event
	// 1. check is eth swap contract
	// 2. check is token exchange event or not
	// 3. check is to furucombo proxy or not
	if curve.IsEthSwapContract(txLog.Address) &&
		curve.IsTokenExchangeEvent(txLog.Topics[0]) &&
		furucombo.IsProxy(txLog.Topics[1]) {
		return true
	}

	return false
}

func findCurveSwapCube(txLog *types.Log) (*Cube, error) {
	if isCurveSwapCube(txLog) {
		cube := Cube{
			Name:         "Curve Swap",
			TokenAddress: txLog.Address,
			TokenAmount:  new(big.Int).SetBytes(txLog.Data),
		}

		return &cube, nil
	}

	// for output eth case that we need to decode event
	if isCurveEthSwapCube(txLog) {
		contractABI, err := abi.JSON(strings.NewReader(curve.SwapContractABI))
		if err != nil {
			return nil, err
		}

		event := new(curve.SwapContractTokenExchange)
		if err := contractABI.UnpackIntoInterface(event, "TokenExchange", txLog.Data); err != nil {
			return nil, err
		}

		// get output token address
		swap, err := curve.NewSwapContract(txLog.Address, ethereum.Client())
		if err != nil {
			return nil, err
		}
		tokenAddress, err := swap.Coins(new(bind.CallOpts), event.BoughtId)
		if err != nil {
			return nil, err
		}
		if ethereum.IsToken("eETH", tokenAddress) {
			tokenAddress = ethereum.GetTokenAddress("WETH")
		}

		cube := Cube{
			Name:         "Curve Swap",
			TokenAddress: tokenAddress,
			TokenAmount:  event.TokensBought,
		}

		return &cube, nil
	}

	return nil, nil
}
