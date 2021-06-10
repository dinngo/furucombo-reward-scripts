package wmatic

import (
	"context"
	"log"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
	"github.com/ethereum/go-ethereum/common"
)

var (
	wmaticAddress     common.Address
	wmaticAddressHash common.Hash
)

// WmaticAddress return wmatic address
func WmaticAddress() common.Address {
	if wmaticAddress == (common.Address{}) {
		wmaticAddress = common.HexToAddress("0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270")
	}

	return wmaticAddress
}

// WmaticAddressHash return wmatic address hash
func WmaticAddressHash() common.Hash {
	if wmaticAddressHash == (common.Hash{}) {
		wmaticAddressHash = WmaticAddress().Hash()
	}

	return wmaticAddressHash
}

// IsWmaticAddress is wmatic address
func IsWmaticAddress(address common.Address) bool {
	return address == WmaticAddress()
}

// IsWmatic is wmatic
func IsWmatic(topic common.Hash) bool {
	return topic == WmaticAddressHash()
}

// HasDepositOrWithdrawalEvent has deposit or withdrawal event
func HasDepositOrWithdrawalEvent(txHash common.Hash) bool {
	// get receipt
	receipt, err := ethereum.ClientPolygon().TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Println(err)
		return false
	}

	for _, txLog := range receipt.Logs {
		if txLog.Topics[0] == common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c") || // Deposit
			txLog.Topics[0] == common.HexToHash("0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65") { // Withdrawal
			return true
		}
	}

	return false
}
