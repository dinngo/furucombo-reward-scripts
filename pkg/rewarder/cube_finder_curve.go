package rewarder

import (
	"log"
	"math/big"

	"garage.dinngo.co/furucombo/mining-scripts/pkg/ethereum/curve"
	"garage.dinngo.co/furucombo/mining-scripts/pkg/ethereum/erc20"
	"garage.dinngo.co/furucombo/mining-scripts/pkg/ethereum/furucombo"
	"github.com/ethereum/go-ethereum/core/types"
)

func init() {
	cubeFinderMap["CurveSwap"] = findCurveSwapCube
}

func isCurveSwapCube(txLog *types.Log) bool {
	// 1. check is supported token or not
	// 2. check is erc20 Transfer event or not
	// 3. check is to furucombo proxy or not
	// 4. check is swap / one split contract or not
	if curve.IsSupportedToken(txLog.Address) &&
		erc20.IsTransferEvent(txLog.Topics[0]) &&
		furucombo.IsProxy(txLog.Topics[2]) &&
		(curve.IsSwapContract(txLog.Topics[1]) || curve.IsOneSplitContract(txLog.Topics[1])) {

		return true
	}

	return false
}

func findCurveSwapCube(txLog *types.Log) (*Cube, error) {
	if !isCurveSwapCube(txLog) {
		return nil, nil
	}

	// check token is listed or not
	if !IsTokenListed(txLog.Address) {
		log.Printf("Curve Swap: %s is not listed", txLog.Address.String())
		return nil, nil
	}

	cube := Cube{
		Name:         "Curve Swap",
		TokenAddress: txLog.Address,
		TokenAmount:  new(big.Int).SetBytes(txLog.Data),
	}

	return &cube, nil
}
