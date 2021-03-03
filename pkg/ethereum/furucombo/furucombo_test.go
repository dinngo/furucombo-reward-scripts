package furucombo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFurucomboAddresses(t *testing.T) {
	os.Setenv("FURUCOMBO_PROXY_ADDRESSES", "0x17e8Ca1b4798B97602895f63206afCd1Fc90Ca5f,0xA013AfbB9A92cEF49e898C87C060e6660E050569")
	proxyAddress := ProxyAddresses()
	assert.Equal(t, "0x17e8Ca1b4798B97602895f63206afCd1Fc90Ca5f", proxyAddress[0].String())
	assert.Equal(t, "0xA013AfbB9A92cEF49e898C87C060e6660E050569", proxyAddress[1].String())
	proxyAddressHash := ProxyAddressHashes()
	assert.Equal(t, "0x00000000000000000000000017e8ca1b4798b97602895f63206afcd1fc90ca5f", proxyAddressHash[0].String())
	assert.Equal(t, "0x000000000000000000000000a013afbb9a92cef49e898c87c060e6660e050569", proxyAddressHash[1].String())

	os.Setenv("FURUCOMBO_DSPROXY_ADDRESSES", "0xa47aC1dd1cF8513B2eB5ca8baD5B8DA93492bF02,0x8BeeDa1F5cD3b5068653a18DB656A5F0c6a1BEeA")
	dsProxyAddress := DSProxyAddresses()
	assert.Equal(t, "0xa47aC1dd1cF8513B2eB5ca8baD5B8DA93492bF02", dsProxyAddress[0].String())
	assert.Equal(t, "0x8BeeDa1F5cD3b5068653a18DB656A5F0c6a1BEeA", dsProxyAddress[1].String())
	dsProxyAddressHash := DSProxyAddressHashes()
	assert.Equal(t, "0x000000000000000000000000a47ac1dd1cf8513b2eb5ca8bad5b8da93492bf02", dsProxyAddressHash[0].String())
	assert.Equal(t, "0x0000000000000000000000008beeda1f5cd3b5068653a18db656a5f0c6a1beea", dsProxyAddressHash[1].String())
}
