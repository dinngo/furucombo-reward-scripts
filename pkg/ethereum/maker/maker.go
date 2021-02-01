package maker

import "github.com/ethereum/go-ethereum/common"

// joinAddressHashMap represents token to join address hash map
var joinAddressHashMap = map[common.Address]common.Hash{
	common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"): common.HexToHash("0x2F0b23f53734252Bda2277357e97e1517d6B042A"), // WETH
	common.HexToAddress("0x0D8775F648430679A709E98d2b0Cb6250d2887EF"): common.HexToHash("0x3D0B1912B66114d4096F48A8CEe3A56C231772cA"), // BAT
	common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"): common.HexToHash("0xA191e578a6736167326d05c119CE0c90849E84B7"), // USDC
	common.HexToAddress("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"): common.HexToHash("0xBF72Da2Bd84c5170618Fbe5914B0ECA9638d5eb5"), // WBTC
	common.HexToAddress("0xdd974D5C2e2928deA5F71b9825b8b646686BD200"): common.HexToHash("0x475F1a89C1ED844A08E8f6C50A00228b5E59E4A9"), // KNC
	common.HexToAddress("0xE41d2489571d322189246DaFA5ebDe1F4699F498"): common.HexToHash("0xc7e8Cd72BDEe38865b4F5615956eF47ce1a7e5D0"), // ZRX
	common.HexToAddress("0x0F5D2fB29fb7d3CFeE444a200298f468908cC942"): common.HexToHash("0xA6EA3b9C04b8a38Ff5e224E7c3D6937ca44C0ef9"), // MANA
	common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"): common.HexToHash("0x0Ac6A1D74E84C2dF9063bDDc31699FF2a2BB22A2"), // USDT
	common.HexToAddress("0x8E870D67F660D95d5be530380D0eC0bd388289E1"): common.HexToHash("0x7e62B7E279DFC78DEB656E34D6a435cC08a44666"), // PAX
	common.HexToAddress("0x0000000000085d4780B73119b644AE5ecd22b376"): common.HexToHash("0x4454aF7C8bb9463203b66C816220D41ED7837f44"), // TUSD
	common.HexToAddress("0xBBbbCA6A901c926F240b89EacB641d8Aec7AEafD"): common.HexToHash("0x6C186404A7A238D3d6027C0299D1822c1cf5d8f1"), // LRC
	common.HexToAddress("0xc00e94Cb662C3520282E6f5717214004A7f26888"): common.HexToHash("0xBEa7cDfB4b49EC154Ae1c0D731E4DC773A3265aA"), // COMP
	common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA"): common.HexToHash("0xdFccAf8fDbD2F4805C174f856a317765B49E4a50"), // LINK
	common.HexToAddress("0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e"): common.HexToHash("0x3ff33d9162aD47660083D7DC4bC02Fb231c81677"), // YFI
	common.HexToAddress("0xba100000625a3754423978a60c9317c58a424e3D"): common.HexToHash("0x4a03Aa7fb3973d8f0221B466EefB53D0aC195f55"), // BAL
	common.HexToAddress("0x056Fd409E1d7A124BD7017459dFEa2F387b6d5Cd"): common.HexToHash("0xe29A14bcDeA40d83675aa43B72dF07f649738C8b"), // GUSD
	common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"): common.HexToHash("0x3BC3A58b4FC1CbE7e98bB4aB7c99535e8bA9b8F1"), // UNI
	common.HexToAddress("0xEB4C2781e4ebA804CE9a9803C67d0893436bB27D"): common.HexToHash("0xFD5608515A47C37afbA68960c1916b79af9491D0"), // RENBTC
	common.HexToAddress("0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9"): common.HexToHash("0x24e459F61cEAa7b1cE70Dbaea938940A7c5aD46e"), // AAVE
	common.HexToAddress("0xA478c2975Ab1Ea89e8196811F51A7B7Ade33eB11"): common.HexToHash("0x2502F65D77cA13f183850b5f9272270454094A08"), // UNIV2DAIETH
	common.HexToAddress("0xBb2b8038a1640196FbE3e38816F3e67Cba72D940"): common.HexToHash("0xDc26C9b7a8fe4F5dF648E314eC3E6Dc3694e6Dd2"), // UNIV2WBTCETH
	common.HexToAddress("0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc"): common.HexToHash("0x03Ae53B33FeeAc1222C3f372f32D37Ba95f0F099"), // UNIV2USDCETH
}

// IsSupportedToken is supported token
func IsSupportedToken(tokenAddress common.Address) (ok bool) {
	_, ok = joinAddressHashMap[tokenAddress]
	return
}

// IsCorrectJoin is correct join
func IsCorrectJoin(tokenAddress common.Address, joinAddressHash common.Hash) bool {
	return joinAddressHash == joinAddressHashMap[tokenAddress]
}
