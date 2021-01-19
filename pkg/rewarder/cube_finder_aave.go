package rewarder

import (
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/aave"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
)

func init() {
	cubeFinderMap["AaveDeposit"] = findAaveDepositCube
}

func isAaveDepositCube(txLog *types.Log) bool {
	// 1. check is supported token or not
	// 2. check is aave MintOnDeposit event or not
	// 3. check is to furucombo proxy or not
	if aave.IsSupportedToken(txLog.Address) &&
		aave.IsMintOnDepositEvent(txLog.Topics[0]) &&
		furucombo.IsProxy(txLog.Topics[1]) {
		return true
	}

	return false
}

func findAaveDepositCube(txLog *types.Log) (*Cube, error) {
	if !isAaveDepositCube(txLog) {
		return nil, nil
	}

	tokenAddress := aave.GetTokenAddress(txLog.Address)

	// check token is listed or not
	if !IsTokenListed(tokenAddress) {
		log.Printf("Aave Deposit: %s is not listed", tokenAddress.String())
		return nil, nil
	}

	contractABI, err := abi.JSON(strings.NewReader(aave.ATokenContractABI))
	if err != nil {
		return nil, err
	}

	event := new(aave.ATokenContractMintOnDeposit)
	if err := contractABI.UnpackIntoInterface(event, "MintOnDeposit", txLog.Data); err != nil {
		return nil, err
	}

	cube := Cube{
		Name:         "Aave Deposit",
		TokenAddress: tokenAddress,
		TokenAmount:  event.Value,
	}

	return &cube, nil
}
