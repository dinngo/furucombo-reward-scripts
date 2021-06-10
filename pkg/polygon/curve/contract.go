package curve

import "github.com/ethereum/go-ethereum/common"

var swapAddresses []common.Address

// SwapAddresses return swap addresses
func SwapAddresses() []common.Address {
	if len(swapAddresses) == 0 {
		swapAddresses = append(swapAddresses, common.HexToAddress("0x445fe580ef8d70ff569ab36e80c647af338db351")) // aave
		swapAddresses = append(swapAddresses, common.HexToAddress("0xC2d95EEF97Ec6C17551d45e77B590dc1F9117C67")) // ren
	}

	return swapAddresses
}

// IsSwapAddress is swap address
func IsSwapAddress(address common.Address) bool {
	for _, v := range SwapAddresses() {
		if address == v {
			return true
		}
	}

	return false
}
