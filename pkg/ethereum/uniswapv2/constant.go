package uniswapv2

import "github.com/ethereum/go-ethereum/common"

var router02AddressHash = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D").Hash()

// IsRouter02Contract is router 02 contract
func IsRouter02Contract(topic common.Hash) bool {
	return topic == router02AddressHash
}
