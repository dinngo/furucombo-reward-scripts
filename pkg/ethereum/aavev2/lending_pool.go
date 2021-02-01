package aavev2

import (
	"github.com/ethereum/go-ethereum/common"
)

var (
	lendingPoolAddress     common.Address
	lendingPoolAddressHash common.Hash
)

// LendingPoolAddress return lending pool address
func LendingPoolAddress() common.Address {
	if lendingPoolAddress == (common.Address{}) {
		lendingPoolAddress = common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9")
	}

	return lendingPoolAddress
}

// LendingPoolAddressHash return lendingPool address hash
func LendingPoolAddressHash() common.Hash {
	if lendingPoolAddressHash == (common.Hash{}) {
		lendingPoolAddressHash = LendingPoolAddress().Hash()
	}

	return lendingPoolAddressHash
}

// IsLendingPoolAddress is lending pool address
func IsLendingPoolAddress(address common.Address) bool {
	return address == LendingPoolAddress()
}

// IsLendingPool is lendingPool
func IsLendingPool(topic common.Hash) bool {
	return topic == LendingPoolAddressHash()
}

var lendingPoolContractDepositEventSig = common.HexToHash("0xde6857219544bb5b7746f48ed30be6386fefc61b2f864cacf559893bf50fd951")

// IsDepositEvent is lending pool contract "Deposit" event
func IsDepositEvent(topic common.Hash) bool {
	return topic == lendingPoolContractDepositEventSig
}
