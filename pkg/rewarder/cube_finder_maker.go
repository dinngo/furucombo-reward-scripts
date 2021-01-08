package rewarder

import (
	"log"
	"math/big"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/erc20"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/maker"
	"github.com/ethereum/go-ethereum/core/types"
)

func init() {
	cubeFinderMap["MakerDeposit"] = findMakerDepositCube
	cubeFinderMap["MakerGenerate"] = findMakerGenerateCube
	cubeFinderMap["MakerPayBack"] = findMakerPayBackCube
}

func isMakerDepositCube(txLog *types.Log) bool {
	// 1. check is supported token or not
	// 2. check is erc20 Transfer event or not
	// 3. check is to furucombo proxy or not
	// 4. check is correct join address or not
	if maker.IsSupportedToken(txLog.Address) &&
		erc20.IsTransferEvent(txLog.Topics[0]) &&
		furucombo.IsDSProxy(txLog.Topics[1]) &&
		maker.IsCorrectJoin(txLog.Address, txLog.Topics[2]) {

		return true
	}

	return false
}

func findMakerDepositCube(txLog *types.Log) (*Cube, error) {
	if !isMakerDepositCube(txLog) {
		return nil, nil
	}

	// check token is listed or not
	if !IsTokenListed(txLog.Address) {
		log.Printf("Maker Deposit: %s is not listed", txLog.Address.String())
		return nil, nil
	}

	cube := Cube{
		Name:         "Maker Deposit",
		TokenAddress: txLog.Address,
		TokenAmount:  new(big.Int).SetBytes(txLog.Data),
	}

	return &cube, nil
}

func isMakerGenerateCube(txLog *types.Log) bool {
	// 1. check is DAI or not
	// 2. check is erc20 Transfer event or not
	// 3. check is zero address or not
	// 4. check is to furucombo proxy or not
	if ethereum.IsToken("DAI", txLog.Address) &&
		erc20.IsTransferEvent(txLog.Topics[0]) &&
		ethereum.IsZeroHash(txLog.Topics[1]) &&
		furucombo.IsProxy(txLog.Topics[2]) {

		return true
	}

	return false
}

func findMakerGenerateCube(txLog *types.Log) (*Cube, error) {
	if !isMakerGenerateCube(txLog) {
		return nil, nil
	}

	// check token is listed or not
	if !IsTokenListed(txLog.Address) {
		log.Printf("Maker Generate: %s is not listed", txLog.Address.String())
		return nil, nil
	}

	cube := Cube{
		Name:         "Maker Generate",
		TokenAddress: txLog.Address,
		TokenAmount:  new(big.Int).SetBytes(txLog.Data),
	}

	return &cube, nil
}

func isMakerPayBackCube(txLog *types.Log) bool {
	// 1. check is DAI or not
	// 2. check is erc20 Transfer event or not
	// 3. check is zero address or not
	// 4. check is to furucombo proxy or not
	if ethereum.IsToken("DAI", txLog.Address) &&
		erc20.IsTransferEvent(txLog.Topics[0]) &&
		furucombo.IsDSProxy(txLog.Topics[1]) &&
		ethereum.IsZeroHash(txLog.Topics[2]) {

		return true
	}

	return false
}

func findMakerPayBackCube(txLog *types.Log) (*Cube, error) {
	if !isMakerPayBackCube(txLog) {
		return nil, nil
	}

	// check token is listed or not
	if !IsTokenListed(txLog.Address) {
		log.Printf("Maker PayBack: %s is not listed", txLog.Address.String())
		return nil, nil
	}

	cube := Cube{
		Name:         "Maker PayBack",
		TokenAddress: txLog.Address,
		TokenAmount:  new(big.Int).SetBytes(txLog.Data),
	}

	return &cube, nil
}
