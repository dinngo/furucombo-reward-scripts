package curve

import "github.com/ethereum/go-ethereum/common"

// event sig
var (
	TokenExchangeEventSig = common.HexToHash("0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140")
)

// IsTokenExchangeEvent is exchange contract "TokenExchange" event
func IsTokenExchangeEvent(topic common.Hash) bool {
	return topic == TokenExchangeEventSig
}

// swapContractAddressMap represents all curve swap contracts except eth
var swapContractAddressMap = map[common.Hash]struct{}{
	common.HexToHash("0xA2B47E3D5c44877cca798226B7B8118F9BFb7A56"): {}, // compound
	common.HexToHash("0x52EA46506B9CC5Ef470C5bf89f17Dc28bB35D85C"): {}, // usdt
	common.HexToHash("0x45F783CCE6B7FF23B2ab2D70e416cdb7D6055f51"): {}, // y
	common.HexToHash("0x79a8C46DeA5aDa233ABaFFD40F3A0A2B1e5A4F27"): {}, // busd
	common.HexToHash("0xA5407eAE9Ba41422680e2e00537571bcC53efBfD"): {}, // susdv2
	common.HexToHash("0x06364f10B501e868329afBc005b3492902d6C763"): {}, // pax
	common.HexToHash("0x93054188d876f558f4a66B2EF1d97d16eDf0895B"): {}, // ren
	common.HexToHash("0x7fC77b5c7614E1533320Ea6DDc2Eb61fa00A9714"): {}, // sbtc
	common.HexToHash("0x4CA9b3063Ec5866A4B82E437059D2C43d1be596F"): {}, // hbtc
	common.HexToHash("0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7"): {}, // 3pool
	common.HexToHash("0x4f062658EaAF2C1ccf8C8e36D6824CDf41167956"): {}, // gusd
	common.HexToHash("0x3eF6A01A0f81D6046290f3e2A8c5b843e738E604"): {}, // husd
	common.HexToHash("0x3E01dD8a5E1fb3481F0F589056b428Fc308AF0Fb"): {}, // usdk
	common.HexToHash("0x0f9cb53Ebe405d49A0bbdBD291A65Ff571bC83e1"): {}, // usdn
	common.HexToHash("0xE7a24EF0C5e95Ffb0f6684b813A78F2a3AD7D171"): {}, // linkusd
	common.HexToHash("0x8474DdbE98F5aA3179B3B3F5942D724aFcdec9f6"): {}, // musd
	common.HexToHash("0xC18cC39da8b11dA8c3541C598eE022258F9744da"): {}, // rsv
	common.HexToHash("0xC25099792E9349C7DD09759744ea681C7de2cb66"): {}, // tbtc
	common.HexToHash("0x8038C01A0390a8c547446a0b2c18fc9aEFEcc10c"): {}, // dusd
	common.HexToHash("0x7F55DDe206dbAD629C080068923b36fe9D6bDBeF"): {}, // pbtc
	common.HexToHash("0x071c661B4DeefB59E2a3DdB20Db036821eeE8F4b"): {}, // bbtc
	common.HexToHash("0xd81dA8D904b52208541Bade1bD6595D8a251F8dd"): {}, // obtc
	common.HexToHash("0x890f4e345B1dAED0367A877a1612f86A1f86985f"): {}, // ust
	common.HexToHash("0x0Ce6a5fF5217e38315f87032CF90686C96627CAA"): {}, // eurs
	common.HexToHash("0xDeBF20617708857ebe4F679508E7b7863a8A8EeE"): {}, // aave
	common.HexToHash("0xEB16Ae0052ed37f479f7fe63849198Df1765a733"): {}, // saave
	common.HexToHash("0xA96A65c051bF88B4095Ee1f2451C2A9d43F53Ae2"): {}, // ankreth
	common.HexToHash("0x2dded6Da1BF5DBdF597C45fcFaa3194e53EcfeAF"): {}, // ironbank
	common.HexToHash("0xF178C0b5Bb7e7aBF4e12A4838C7b7c5bA2C623c0"): {}, // link
}

// IsSwapContract is swap contract
func IsSwapContract(topic common.Hash) (ok bool) {
	_, ok = swapContractAddressMap[topic]
	return
}

// ethSwapContractAddressMap represents curve eth swap contracts
var ethSwapContractAddressMap = map[common.Address]struct{}{
	common.HexToAddress("0xc5424B857f758E906013F3555Dad202e4bdB4567"): {}, // seth
	common.HexToAddress("0xDC24316b9AE028F1497c275EB9192a3Ea0f67022"): {}, // steth
}

// IsEthSwapContract is eth swap contract
func IsEthSwapContract(topic common.Address) (ok bool) {
	_, ok = ethSwapContractAddressMap[topic]
	return
}

var oneSplitContractAddressHash = common.HexToHash("0xC586BeF4a0992C495Cf22e1aeEE4E446CECDee0E")

// IsOneSplitContract is one splic contract
func IsOneSplitContract(topic common.Hash) bool {
	return topic == oneSplitContractAddressHash
}
