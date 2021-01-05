package furucombo

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

var (
	proxyAddress     common.Address
	proxyAddressHash common.Hash
)

// ProxyAddress return proxy address
func ProxyAddress() common.Address {
	if proxyAddress == (common.Address{}) {
		proxyAddress = common.HexToAddress(os.Getenv("FURUCOMBO_PROXY_ADDRESS"))

		if proxyAddress == (common.Address{}) {
			log.Fatalln("env FURUCOMBO_PROXY_ADDRESS can't be blank")
		}
	}

	return proxyAddress
}

// ProxyAddressHash return proxy address hash
func ProxyAddressHash() common.Hash {
	if proxyAddressHash == (common.Hash{}) {
		proxyAddressHash = ProxyAddress().Hash()
	}

	return proxyAddressHash
}

// IsProxyAddress is proxy address
func IsProxyAddress(address common.Address) bool {
	return address == ProxyAddress()
}

// IsProxy is proxy
func IsProxy(topic common.Hash) bool {
	return topic == ProxyAddressHash()
}

var (
	dsProxyAddress     common.Address
	dsProxyAddressHash common.Hash
)

// DSProxyAddress return ds proxy address
func DSProxyAddress() common.Address {
	if dsProxyAddress == (common.Address{}) {
		dsProxyAddress = common.HexToAddress(os.Getenv("FURUCOMBO_DSPROXY_ADDRESS"))

		if dsProxyAddress == (common.Address{}) {
			log.Fatalln("env FURUCOMBO_DSPROXY_ADDRESS can't be blank")
		}
	}

	return dsProxyAddress
}

// DSProxyAddressHash return ds proxy address hash
func DSProxyAddressHash() common.Hash {
	if dsProxyAddressHash == (common.Hash{}) {
		dsProxyAddressHash = DSProxyAddress().Hash()
	}

	return dsProxyAddressHash
}

// IsDSProxyAddress is ds proxy address
func IsDSProxyAddress(address common.Address) bool {
	return address == DSProxyAddress()
}

// IsDSProxy is ds proxy
func IsDSProxy(topic common.Hash) bool {
	return topic == DSProxyAddressHash()
}
