package yearn

import "github.com/ethereum/go-ethereum/common"

// vaultAddressHashMap represents token to vault address hash map
var vaultAddressHashMap = map[common.Address]common.Hash{
	common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"): common.HexToHash("0x597aD1e0c13Bfe8025993D9e79C69E1c0233522e"), // yUSDC
	common.HexToAddress("0xdF5e0e81Dff6FAF3A7e52BA697820c5e32D806A8"): common.HexToHash("0x5dbcF33D8c2E976c6b560249878e6F1491Bca25c"), // yUSD
	common.HexToAddress("0x0000000000085d4780B73119b644AE5ecd22b376"): common.HexToHash("0x37d19d1c4E1fa9DC47bD1eA12f742a0887eDa74a"), // yTUSD
	common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"): common.HexToHash("0xACd43E627e64355f1861cEC6d3a6688B31a6F952"), // yDAI
	common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"): common.HexToHash("0x2f08119C6f07c006695E079AAFc638b8789FAf18"), // yUSDT
	common.HexToAddress("0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e"): common.HexToHash("0xBA2E7Fed597fd0E3e70f5130BcDbbFE06bB94fe1"), // yYFI
	common.HexToAddress("0x3B3Ac5386837Dc563660FB6a0937DFAa5924333B"): common.HexToHash("0x2994529C0652D127b7842094103715ec5299bBed"), // ycrvBUSD
	common.HexToAddress("0x075b1bb99792c9E1041bA13afEf80C91a1e70fB3"): common.HexToHash("0x7Ff566E1d69DEfF32a7b244aE7276b9f90e9D0f6"), // ycrvBTC
	common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"): common.HexToHash("0xe1237aA7f535b0CC33Fd973D66cBf830354D16c7"), // yWETH
	common.HexToAddress("0x6c3F90f043a72FA612cbac8115EE7e52BDe6E490"): common.HexToHash("0x9cA85572E6A3EbF24dEDd195623F188735A5179f"), // y3Crv
	common.HexToAddress("0x056Fd409E1d7A124BD7017459dFEa2F387b6d5Cd"): common.HexToHash("0xec0d8D3ED5477106c6D4ea27D90a60e594693C90"), // yGUSD
	common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA"): common.HexToHash("0x881b06da56BB5675c54E4Ed311c21E54C5025298"), // yLINK
	common.HexToAddress("0xA64BD6C70Cb9051F6A9ba1F163Fdc07E0DfB5F84"): common.HexToHash("0x29E240CFD7946BA20895a7a02eDb25C210f9f324"), // yaLINK
}

// IsSupportedToken is supported token
func IsSupportedToken(tokenAddress common.Address) (ok bool) {
	_, ok = vaultAddressHashMap[tokenAddress]
	return
}

// IsCorrectVault is correct vault
func IsCorrectVault(tokenAddress common.Address, vaultAddressHash common.Hash) bool {
	return vaultAddressHash == vaultAddressHashMap[tokenAddress]
}
