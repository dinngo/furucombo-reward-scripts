package furucombo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFurucomboAddresses(t *testing.T) {
	os.Setenv("FURUCOMBO_PROXY_ADDRESS", "0x17e8Ca1b4798B97602895f63206afCd1Fc90Ca5f")
	proxyAddress := ProxyAddress()
	assert.Equal(t, "0x17e8Ca1b4798B97602895f63206afCd1Fc90Ca5f", proxyAddress.String())
	proxyAddressHash := ProxyAddressHash()
	assert.Equal(t, "0x00000000000000000000000017e8ca1b4798b97602895f63206afcd1fc90ca5f", proxyAddressHash.String())

	os.Setenv("FURUCOMBO_DSPROXY_ADDRESS", "0xa47aC1dd1cF8513B2eB5ca8baD5B8DA93492bF02")
	dsProxyAddress := DSProxyAddress()
	assert.Equal(t, "0xa47aC1dd1cF8513B2eB5ca8baD5B8DA93492bF02", dsProxyAddress.String())
	dsProxyAddressHash := DSProxyAddressHash()
	assert.Equal(t, "0x000000000000000000000000a47ac1dd1cf8513b2eb5ca8bad5b8da93492bf02", dsProxyAddressHash.String())
}
