package ethereum

import "github.com/ethereum/go-ethereum/common"

var wethDepositEventSit = common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c")

// IsWETHDepositEvent is WETH "Deposit" event
func IsWETHDepositEvent(topic common.Hash) bool {
	return topic == wethDepositEventSit
}
