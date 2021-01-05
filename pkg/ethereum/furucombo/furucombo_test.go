package furucombo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFurucomboAddresses(t *testing.T) {
	os.Setenv("FURUCOMBO_PROXY_ADDRESS", "0xEEbeC7874f46C4452610A11FA6bE0264f7f0557F")
	proxyAddress := ProxyAddress()
	assert.Equal(t, "0xEEbeC7874f46C4452610A11FA6bE0264f7f0557F", proxyAddress.String())
	proxyAddressHash := ProxyAddressHash()
	assert.Equal(t, "0x000000000000000000000000eebec7874f46c4452610a11fa6be0264f7f0557f", proxyAddressHash.String())

	os.Setenv("FURUCOMBO_DSPROXY_ADDRESS", "0xa5c1D410bA1a2839591A9672CC4e0e0281b23b1c")
	dsProxyAddress := DSProxyAddress()
	assert.Equal(t, "0xa5c1D410bA1a2839591A9672CC4e0e0281b23b1c", dsProxyAddress.String())
	dsProxyAddressHash := DSProxyAddressHash()
	assert.Equal(t, "0x000000000000000000000000a5c1d410ba1a2839591a9672cc4e0e0281b23b1c", dsProxyAddressHash.String())
}
