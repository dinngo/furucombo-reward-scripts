package ethereum

import "github.com/ethereum/go-ethereum/common"

var tokenAddressMap = map[string]common.Address{
	"DAI":  common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"),
	"WETH": common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
	"eETH": common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
}

// IsToken is token
func IsToken(symbol string, address common.Address) bool {
	return address == tokenAddressMap[symbol]
}

// GetToken get token
func GetToken(symbol string) common.Address {
	return tokenAddressMap[symbol]
}
