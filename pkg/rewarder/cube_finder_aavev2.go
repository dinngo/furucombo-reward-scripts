package rewarder

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/aavev2"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
)

func init() {
	cubeFinderMap["AaveV2Deposit"] = findAaveV2DepositCube
}

func isAaveV2DepositCube(txLog *types.Log) bool {
	// 1. check is lending pool address
	// 2. check is aave Deposit event or not
	// 3. check is to furucombo proxy or not
	if aavev2.IsLendingPoolAddress(txLog.Address) &&
		aavev2.IsDepositEvent(txLog.Topics[0]) &&
		furucombo.IsProxy(txLog.Topics[2]) {
		return true
	}

	return false
}

func findAaveV2DepositCube(txLog *types.Log) (*Cube, error) {
	if !isAaveV2DepositCube(txLog) {
		return nil, nil
	}

	tokenAddress := common.HexToAddress(txLog.Topics[1].Hex()) // reserve

	contractABI, err := abi.JSON(strings.NewReader(aavev2.LendingPoolContractABI))
	if err != nil {
		return nil, err
	}

	event := new(aavev2.LendingPoolContractDeposit)
	if err := contractABI.UnpackIntoInterface(event, "Deposit", txLog.Data); err != nil {
		return nil, err
	}

	cube := Cube{
		Name:         "Aave V2 Deposit",
		TokenAddress: tokenAddress,
		TokenAmount:  event.Amount,
	}

	return &cube, nil
}
