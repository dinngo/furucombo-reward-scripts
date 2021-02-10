package rewarder

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/compound"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
)

func init() {
	cubeFinderMap["CompoundSupply"] = findCompoundSupplyCube
	cubeFinderMap["CompoundRepay"] = findCompoundRepayCube
	cubeFinderMap["CompoundBorrow"] = findCompoundBorrowCube
}

func isCompoundSupplyCube(txLog *types.Log) bool {
	// 1. check is supported token or not
	// 2. check is compound Mint event or not
	if compound.IsSupportedToken(txLog.Address) && compound.IsMintEvent(txLog.Topics[0]) {
		return true
	}

	return false
}

func findCompoundSupplyCube(txLog *types.Log) (*Cube, error) {
	if !isCompoundSupplyCube(txLog) {
		return nil, nil
	}

	tokenAddress := compound.GetTokenAddress(txLog.Address)

	contractABI, err := abi.JSON(strings.NewReader(compound.CTokenContractABI))
	if err != nil {
		return nil, err
	}

	event := new(compound.CTokenContractMint)
	if err := contractABI.UnpackIntoInterface(event, "Mint", txLog.Data); err != nil {
		return nil, err
	}

	// check is from furucombo proxy or not
	if !furucombo.IsProxyAddress(event.Minter) {
		return nil, nil
	}

	cube := Cube{
		Name:         "Compound Supply",
		TokenAddress: tokenAddress,
		TokenAmount:  event.MintAmount,
	}

	return &cube, nil
}

func isCompoundRepayCube(txLog *types.Log) bool {
	// 1. check is supported token or not
	// 2. check is compound RepayBorrow event
	if compound.IsSupportedToken(txLog.Address) && compound.IsRepayBorrowEvent(txLog.Topics[0]) {
		return true
	}

	return false
}

func findCompoundRepayCube(txLog *types.Log) (*Cube, error) {
	if !isCompoundRepayCube(txLog) {
		return nil, nil
	}

	tokenAddress := compound.GetTokenAddress(txLog.Address)

	contractABI, err := abi.JSON(strings.NewReader(compound.CTokenContractABI))
	if err != nil {
		return nil, err
	}

	event := new(compound.CTokenContractRepayBorrow)
	if err := contractABI.UnpackIntoInterface(event, "RepayBorrow", txLog.Data); err != nil {
		return nil, err
	}

	cube := Cube{
		Name:         "Compound Repay",
		TokenAddress: tokenAddress,
		TokenAmount:  event.RepayAmount,
	}

	return &cube, nil
}

func isCompoundBorrowCube(txLog *types.Log) bool {
	// 1. check is supported token or not
	// 2. check is compound Borrow event or not
	if compound.IsSupportedToken(txLog.Address) && compound.IsBorrowEvent(txLog.Topics[0]) {
		return true
	}

	return false
}

func findCompoundBorrowCube(txLog *types.Log) (*Cube, error) {
	if !isCompoundBorrowCube(txLog) {
		return nil, nil
	}

	tokenAddress := compound.GetTokenAddress(txLog.Address)

	contractABI, err := abi.JSON(strings.NewReader(compound.CTokenContractABI))
	if err != nil {
		return nil, err
	}

	event := new(compound.CTokenContractBorrow)
	if err := contractABI.UnpackIntoInterface(event, "Borrow", txLog.Data); err != nil {
		return nil, err
	}

	cube := Cube{
		Name:         "Compound Borrow",
		TokenAddress: tokenAddress,
		TokenAmount:  event.BorrowAmount,
	}

	return &cube, nil
}
