package curve

import "github.com/ethereum/go-ethereum/common"

// tokenAddressMap represents listed output tokens
var tokenAddressMap = map[common.Address]struct{}{
	common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"): {}, // DAI
	common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"): {}, // USDC
	common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"): {}, // USDT
	common.HexToAddress("0x57Ab1ec28D129707052df4dF418D58a2D46d5f51"): {}, // sUSD
	common.HexToAddress("0x0000000000085d4780B73119b644AE5ecd22b376"): {}, // TUSD
	common.HexToAddress("0x8E870D67F660D95d5be530380D0eC0bd388289E1"): {}, // PAX
	common.HexToAddress("0x4Fabb145d64652a948d72533023f6E7A623C7C53"): {}, // BUSD
	common.HexToAddress("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"): {}, // WBTC
	common.HexToAddress("0xEB4C2781e4ebA804CE9a9803C67d0893436bB27D"): {}, // renBTC
	common.HexToAddress("0xfE18be6b3Bd88A2D2A7f928d00292E7a9963CfC6"): {}, // sBTC
	common.HexToAddress("0x0316EB71485b0Ab14103307bf65a021042c6d380"): {}, // HBTC
	common.HexToAddress("0x056Fd409E1d7A124BD7017459dFEa2F387b6d5Cd"): {}, // GUSD
	common.HexToAddress("0xdF574c24545E5FfEcb9a659c229253D4111d87e1"): {}, // HUSD
	common.HexToAddress("0x1c48f86ae57291F7686349F12601910BD8D470bb"): {}, // USDK
	common.HexToAddress("0x674C6Ad92Fd080e4004b2312b45f796a192D27a0"): {}, // USDN
	common.HexToAddress("0x0E2EC54fC0B509F445631Bf4b91AB8168230C752"): {}, // LINKUSD
	common.HexToAddress("0xe2f2a5C287993345a840Db3B0845fbC70f5935a5"): {}, // mUSD
	common.HexToAddress("0x196f4727526eA7FB1e17b2071B3d8eAA38486988"): {}, // RSV
	common.HexToAddress("0x8dAEBADE922dF735c38C80C7eBD708Af50815fAa"): {}, // TBTC
	common.HexToAddress("0x5BC25f649fc4e26069dDF4cF4010F9f706c23831"): {}, // DUSD
	common.HexToAddress("0x5228a22e72ccC52d415EcFd199F99D0665E7733b"): {}, // pBTC
	common.HexToAddress("0x8064d9Ae6cDf087b1bcd5BDf3531bD5d8C537a68"): {}, // oBTC
}

// IsSupportedToken is supported token
func IsSupportedToken(tokenAddress common.Address) (ok bool) {
	_, ok = tokenAddressMap[tokenAddress]
	return
}
