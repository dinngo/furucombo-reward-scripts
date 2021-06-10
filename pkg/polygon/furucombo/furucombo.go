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
		key := os.Getenv("POLYGON_FURUCOMBO_PROXY_ADDRESSES")
		if len(key) == 0 {
			log.Fatalln("env POLYGON_FURUCOMBO_PROXY_ADDRESSES can't be blank")
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
