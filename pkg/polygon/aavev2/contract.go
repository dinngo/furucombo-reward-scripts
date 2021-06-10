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
		lendingPoolAddress = common.HexToAddress("0x8dFf5E27EA6b7AC08EbFdf9eB090F32ee9a30fcf")
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
