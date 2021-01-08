package rewarder

import (
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/compound"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
)

func init() {
	cubeFinderMap["CompoundSupply"] = findCompoundSupplyCube
	cubeFinderMap["CompoundRepay"] = findCompoundRepayCube
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

	// check token is listed or not
	if !IsTokenListed(tokenAddress) {
		log.Printf("Compound Supply: %s is not listed", tokenAddress.String())
		return nil, nil
	}

	contractABI, err := abi.JSON(strings.NewReader(compound.CTokenContractABI))
	if err != nil {
		return nil, err
	}

	event := new(compound.CTokenContractMint)
	if err := contractABI.Unpack(event, "Mint", txLog.Data); err != nil {
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

	// check token is listed or not
	if !IsTokenListed(tokenAddress) {
		log.Printf("Compound Repay: %s is not listed", tokenAddress.String())
		return nil, nil
	}

	contractABI, err := abi.JSON(strings.NewReader(compound.CTokenContractABI))
	if err != nil {
		return nil, err
	}

	event := new(compound.CTokenContractRepayBorrow)
	if err := contractABI.Unpack(event, "RepayBorrow", txLog.Data); err != nil {
		return nil, err
	}

	// 3. check is from furucombo proxy or not
	if !furucombo.IsProxyAddress(event.Payer) {
		return nil, nil
	}

	cube := Cube{
		Name:         "Compound Repay",
		TokenAddress: tokenAddress,
		TokenAmount:  event.RepayAmount,
	}

	return &cube, nil
}
