package rewarder

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/erc20"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/yearn"
)

func init() {
	cubeFinderMap["YearnDeposit"] = findYearnDepositCube
}

func isYearnDepositCube(txLog *types.Log) bool {
	// 1. check is supported token or not
	if yearn.IsSupportedToken(txLog.Address) {

		// 2-1. check is erc20 Transfer event or not
		// 2-2. check is to furucombo proxy or not
		// 2-3. check is correct vault address or not
		if erc20.IsTransferEvent(txLog.Topics[0]) &&
			furucombo.IsProxy(txLog.Topics[1]) &&
			yearn.IsCorrectVault(txLog.Address, txLog.Topics[2]) {

			return true
		}

		// 3-1. check is WETH or not
		// 3-2. check is WETH deposit event or not, note that it doesn't check furucombo related address
		// 2-3. check is correct vault address or not
		if ethereum.IsToken("WETH", txLog.Address) &&
			ethereum.IsWETHDepositEvent(txLog.Topics[0]) &&
			yearn.IsCorrectVault(txLog.Address, txLog.Topics[1]) {
			return true
		}

	}

	return false
}

func findYearnDepositCube(txLog *types.Log) (*Cube, error) {
	if !isYearnDepositCube(txLog) {
		return nil, nil
	}

	// check token is listed or not
	if !IsTokenListed(txLog.Address) {
		log.Printf("Yearn Deposit: %s is not listed", txLog.Address.String())
		return nil, nil
	}

	cube := Cube{
		Name:         "Yearn Deposit",
		TokenAddress: txLog.Address,
		TokenAmount:  new(big.Int).SetBytes(txLog.Data),
	}

	return &cube, nil
}
