package aave

import "github.com/ethereum/go-ethereum/common"

// tokenAddressMap represents aToken to underlying token address map
var tokenAddressMap = map[common.Address]common.Address{
	common.HexToAddress("0x3a3A65aAb0dd2A17E3F1947bA16138cd37d08c04"): common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), // aETH
	common.HexToAddress("0xfC1E690f61EFd961294b3e1Ce3313fBD8aa4f85d"): common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"), // aDAI
	common.HexToAddress("0x9bA00D6856a4eDF4665BcA2C2309936572473B7E"): common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"), // aUSDC
	common.HexToAddress("0x625aE63000f46200499120B906716420bd059240"): common.HexToAddress("0x57Ab1ec28D129707052df4dF418D58a2D46d5f51"), // aSUSD
	common.HexToAddress("0x4DA9b813057D04BAef4e5800E36083717b4a0341"): common.HexToAddress("0x0000000000085d4780B73119b644AE5ecd22b376"), // aTUSD
	common.HexToAddress("0x71fc860F7D3A592A4a98740e39dB31d25db65ae8"): common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"), // aUSDT
	common.HexToAddress("0x6Ee0f7BB50a54AB5253dA0667B0Dc2ee526C30a8"): common.HexToAddress("0x4Fabb145d64652a948d72533023f6E7A623C7C53"), // aBUSD
	common.HexToAddress("0xE1BA0FB44CCb0D11b80F92f4f8Ed94CA3fF51D00"): common.HexToAddress("0x0D8775F648430679A709E98d2b0Cb6250d2887EF"), // aBAT
	common.HexToAddress("0x712DB54daA836B53Ef1EcBb9c6ba3b9Efb073F40"): common.HexToAddress("0xF629cBd94d3791C9250152BD8dfBDF380E2a3B9c"), // aENJ
	common.HexToAddress("0x9D91BE44C06d373a8a226E1f3b146956083803eB"): common.HexToAddress("0xdd974D5C2e2928deA5F71b9825b8b646686BD200"), // aKNC
	common.HexToAddress("0xba3D9687Cf50fE253cd2e1cFeEdE1d6787344Ed5"): common.HexToAddress("0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9"), // aAAVE
	common.HexToAddress("0xA64BD6C70Cb9051F6A9ba1F163Fdc07E0DfB5F84"): common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA"), // aLINK
	common.HexToAddress("0x6FCE4A401B6B80ACe52baAefE4421Bd188e76F6f"): common.HexToAddress("0x0F5D2fB29fb7d3CFeE444a200298f468908cC942"), // aMANA
	common.HexToAddress("0x7deB5e830be29F91E298ba5FF1356BB7f8146998"): common.HexToAddress("0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2"), // aMKR
	common.HexToAddress("0x69948cC03f478B95283F7dbf1CE764d0fc7EC54C"): common.HexToAddress("0x408e41876cCCDC0F92210600ef50372656052a38"), // aREN
	common.HexToAddress("0x328C4c80BC7aCa0834Db37e6600A6c49E12Da4DE"): common.HexToAddress("0xC011a73ee8576Fb46F5E1c5751cA3B9Fe0af2a6F"), // aSNX
	common.HexToAddress("0xB124541127A0A657f056D9Dd06188c4F1b0e5aab"): common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"), // aUNI
	common.HexToAddress("0xFC4B8ED459e00e5400be803A9BB3954234FD50e3"): common.HexToAddress("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"), // aWBTC
	common.HexToAddress("0x12e51E77DAAA58aA0E9247db7510Ea4B46F9bEAd"): common.HexToAddress("0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e"), // aYFI
	common.HexToAddress("0x6Fb0855c404E09c47C3fBCA25f08d4E41f9F062f"): common.HexToAddress("0xE41d2489571d322189246DaFA5ebDe1F4699F498"), // aZRX
}

// IsSupportedToken is supported token
func IsSupportedToken(aTokenAddress common.Address) (ok bool) {
	_, ok = tokenAddressMap[aTokenAddress]
	return
}

// GetTokenAddress get token Address
func GetTokenAddress(aTokenAddress common.Address) common.Address {
	return tokenAddressMap[aTokenAddress]
}

var aTokenContractMintOnDepositEventSig = common.HexToHash("0xbe7799898ca2d813ff902b487c1b434ab45b47edd8f38b77ad5e99aae8341b7a")

// IsMintOnDepositEvent is aToken "MintOnDeposit" event
func IsMintOnDepositEvent(topic common.Hash) bool {
	return topic == aTokenContractMintOnDepositEventSig
}
