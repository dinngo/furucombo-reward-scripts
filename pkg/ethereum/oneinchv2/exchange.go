package oneinchv2

import "github.com/ethereum/go-ethereum/common"

var exchangeContractSwappedEventSig = common.HexToHash("0x76af224a143865a50b41496e1a73622698692c565c1214bc862f18e22d829c5e")

// IsSwappedEvent is exchange contract "Swapped" event
func IsSwappedEvent(topic common.Hash) bool {
	return topic == exchangeContractSwappedEventSig
}
