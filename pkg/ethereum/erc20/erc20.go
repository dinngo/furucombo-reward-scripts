package erc20

import "github.com/ethereum/go-ethereum/common"

var transferEventSig = common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")

// IsTransferEvent is "Transfer" event
func IsTransferEvent(topic common.Hash) bool {
	return topic == transferEventSig
}
