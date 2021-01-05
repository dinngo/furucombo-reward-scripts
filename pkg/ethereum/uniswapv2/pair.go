package uniswapv2

import "github.com/ethereum/go-ethereum/common"

var pairContractSwapEventSig = common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822")

// IsSwapEvent is pair contract "Swap" event
func IsSwapEvent(topic common.Hash) bool {
	return topic == pairContractSwapEventSig
}
