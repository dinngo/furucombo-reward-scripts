package furucombo

import (
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

var (
	proxyAddresses     []common.Address
	proxyAddressHashes []common.Hash
)

// ProxyAddresses return proxy addresses
func ProxyAddresses() []common.Address {
	if len(proxyAddresses) == 0 {
		key := os.Getenv("FURUCOMBO_PROXY_ADDRESSES")
		if len(key) == 0 {
			log.Fatalln("env FURUCOMBO_PROXY_ADDRESSES can't be blank")
		}

		addresses := strings.Split(key, ",")
		for _, v := range addresses {
			proxyAddresses = append(proxyAddresses, common.HexToAddress(v))
		}
	}

	return proxyAddresses
}

// ProxyAddressHashes return proxy address hashes
func ProxyAddressHashes() []common.Hash {
	if len(proxyAddressHashes) == 0 {
		for _, v := range ProxyAddresses() {
			proxyAddressHashes = append(proxyAddressHashes, v.Hash())
		}
	}

	return proxyAddressHashes
}

// IsProxyAddress is proxy address
func IsProxyAddress(address common.Address) bool {
	for _, v := range ProxyAddresses() {
		if address == v {
			return true
		}
	}

	return false
}

// IsProxy is proxy
func IsProxy(topic common.Hash) bool {
	for _, v := range ProxyAddressHashes() {
		if topic == v {
			return true
		}
	}

	return false
}

var (
	dsProxyAddresses     []common.Address
	dsProxyAddressHashes []common.Hash
)

// DSProxyAddresses return ds proxy addresses
func DSProxyAddresses() []common.Address {
	if len(dsProxyAddresses) == 0 {
		key := os.Getenv("FURUCOMBO_DSPROXY_ADDRESSES")
		if len(key) == 0 {
			log.Fatalln("env FURUCOMBO_DSPROXY_ADDRESSES can't be blank")
		}

		addresses := strings.Split(key, ",")
		for _, v := range addresses {
			dsProxyAddresses = append(dsProxyAddresses, common.HexToAddress(v))
		}
	}

	return dsProxyAddresses
}

// DSProxyAddressHashes return ds proxy address hashes
func DSProxyAddressHashes() []common.Hash {
	if len(dsProxyAddressHashes) == 0 {
		for _, v := range DSProxyAddresses() {
			dsProxyAddressHashes = append(dsProxyAddressHashes, v.Hash())
		}
	}

	return dsProxyAddressHashes
}

// IsDSProxyAddress is ds proxy address
func IsDSProxyAddress(address common.Address) bool {
	for _, v := range DSProxyAddresses() {
		if address == v {
			return true
		}
	}

	return false
}

// IsDSProxy is ds proxy
func IsDSProxy(topic common.Hash) bool {
	for _, v := range DSProxyAddressHashes() {
		if topic == v {
			return true
		}
	}

	return false
}
