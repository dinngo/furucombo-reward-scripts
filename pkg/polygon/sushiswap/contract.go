package sushiswap

import "github.com/ethereum/go-ethereum/common"

var router02Address common.Address

// Router02Address return router02 address
func Router02Address() common.Address {
	if router02Address == (common.Address{}) {
		router02Address = common.HexToAddress("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506")
	}

	return router02Address
}

// IsRouter02Address is router02 address
func IsRouter02Address(address common.Address) bool {
	return address == Router02Address()
}
