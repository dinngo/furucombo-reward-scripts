package compound

import "github.com/ethereum/go-ethereum/common"

// tokenAddressMap represents ctoken to underlying token address map
var tokenAddressMap = map[common.Address]common.Address{
	common.HexToAddress("0x6C8c6b02E7b2BE14d4fA6022Dfd6d75921D90E4E"): common.HexToAddress("0x0D8775F648430679A709E98d2b0Cb6250d2887EF"), // cBAT
	common.HexToAddress("0x70e36f6BF80a52b3B46b3aF8e106CC0ed743E8e4"): common.HexToAddress("0xc00e94Cb662C3520282E6f5717214004A7f26888"), // cCOMP
	common.HexToAddress("0x5d3a536E4D6DbD6114cc1Ead35777bAB948E3643"): common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"), // cDAI
	common.HexToAddress("0x4Ddc2D193948926D02f9B1fE9e1daa0718270ED5"): common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), // cETH
	common.HexToAddress("0x35A18000230DA775CAc24873d00Ff85BccdeD550"): common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"), // cUNI
	common.HexToAddress("0x39AA39c021dfbaE8faC545936693aC917d5E7563"): common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"), // cUSDC
	common.HexToAddress("0xf650C3d88D12dB855b8bf7D11Be6C55A4e07dCC9"): common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"), // cUSDT
	common.HexToAddress("0xC11b1268C1A384e55C48c2391d8d480264A3A7F4"): common.HexToAddress("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"), // cWBTC
	common.HexToAddress("0xB3319f5D18Bc0D84dD1b4825Dcde5d5f7266d407"): common.HexToAddress("0xE41d2489571d322189246DaFA5ebDe1F4699F498"), // cZRX
}

// IsSupportedToken is supported token
func IsSupportedToken(cTokenAddress common.Address) (ok bool) {
	_, ok = tokenAddressMap[cTokenAddress]
	return
}

// GetTokenAddress get token Address
func GetTokenAddress(cTokenAddress common.Address) common.Address {
	return tokenAddressMap[cTokenAddress]
}

var cTokenContractMintEventSig = common.HexToHash("0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f")

// IsMintEvent is cToken "Mint" event
func IsMintEvent(topic common.Hash) bool {
	return topic == cTokenContractMintEventSig
}

var cTokenContractRepayBorrowEventSig = common.HexToHash("0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1")

// IsRepayBorrowEvent is cToken "RepayBorrow" event
func IsRepayBorrowEvent(topic common.Hash) bool {
	return topic == cTokenContractRepayBorrowEventSig
}
