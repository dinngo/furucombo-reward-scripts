package quickswap

import "github.com/ethereum/go-ethereum/common"

var router02Address common.Address

// Router02Address return router02 address
func Router02Address() common.Address {
	if router02Address == (common.Address{}) {
		router02Address = common.HexToAddress("0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff")
	}

	return router02Address
}

// IsRouter02Address is router02 address
func IsRouter02Address(address common.Address) bool {
	return address == Router02Address()
}
