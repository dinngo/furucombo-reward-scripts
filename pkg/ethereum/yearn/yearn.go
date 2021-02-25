package yearn

import "github.com/ethereum/go-ethereum/common"

// vaultAddressHashMap represents token to vault address hash map
var vaultAddressHashMap = map[common.Address]common.Hash{
	common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"): common.HexToHash("0x597aD1e0c13Bfe8025993D9e79C69E1c0233522e"), // yUSDC
	common.HexToAddress("0xdF5e0e81Dff6FAF3A7e52BA697820c5e32D806A8"): common.HexToHash("0x5dbcF33D8c2E976c6b560249878e6F1491Bca25c"), // yUSD
	common.HexToAddress("0x0000000000085d4780B73119b644AE5ecd22b376"): common.HexToHash("0x37d19d1c4E1fa9DC47bD1eA12f742a0887eDa74a"), // yTUSD
	common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"): common.HexToHash("0xACd43E627e64355f1861cEC6d3a6688B31a6F952"), // yDAI
	common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"): common.HexToHash("0x2f08119C6f07c006695E079AAFc638b8789FAf18"), // yUSDT
	common.HexToAddress("0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e"): common.HexToHash("0xBA2E7Fed597fd0E3e70f5130BcDbbFE06bB94fe1"), // yvYFI
	common.HexToAddress("0x3B3Ac5386837Dc563660FB6a0937DFAa5924333B"): common.HexToHash("0x2994529C0652D127b7842094103715ec5299bBed"), // yvcrvBUSD
	common.HexToAddress("0x075b1bb99792c9E1041bA13afEf80C91a1e70fB3"): common.HexToHash("0x7Ff566E1d69DEfF32a7b244aE7276b9f90e9D0f6"), // ycrvRenWSBTC
	common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"): common.HexToHash("0xe1237aA7f535b0CC33Fd973D66cBf830354D16c7"), // yWETH
	common.HexToAddress("0x6c3F90f043a72FA612cbac8115EE7e52BDe6E490"): common.HexToHash("0x9cA85572E6A3EbF24dEDd195623F188735A5179f"), // yv3Crv
	common.HexToAddress("0x514910771AF9Ca656af840dff83E8264EcF986CA"): common.HexToHash("0x881b06da56BB5675c54E4Ed311c21E54C5025298"), // yLINK
	common.HexToAddress("0xA64BD6C70Cb9051F6A9ba1F163Fdc07E0DfB5F84"): common.HexToHash("0x29E240CFD7946BA20895a7a02eDb25C210f9f324"), // yaLINK-V1
	common.HexToAddress("0x845838DF265Dcd2c412A1Dc9e959c7d08537f8a2"): common.HexToHash("0x629c759D1E83eFbF63d84eb3868B564d9521C129"), // yvcDAI+cUSDC
	common.HexToAddress("0xD2967f45c4f384DEEa880F807Be904762a3DeA07"): common.HexToHash("0xcC7E70A958917cCe67B4B87a8C30E6297451aE98"), // yvgusd3CRV
	common.HexToAddress("0x1AEf73d49Dedc4b1778d0706583995958Dc862e6"): common.HexToHash("0x0FCDAeDFb8A7DfDa2e9838564c5A1665d856AFDF"), // yvmusd3CRV
	common.HexToAddress("0xe2f2a5C287993345a840Db3B0845fbC70f5935a5"): common.HexToHash("0xE0db48B4F71752C4bEf16De1DBD042B82976b8C7"), // yvmUSD
	common.HexToAddress("0xFd2a8fA60Abd58Efe3EeE34dd494cD491dC14900"): common.HexToHash("0x03403154afc09Ce8e44C3B185C82C6aD5f86b9ab"), // yva3CRV
	common.HexToAddress("0x3a664Ab939FD8482048609f652f9a0B0677337B9"): common.HexToHash("0x8e6741b456a074F0Bc45B8b82A755d4aF7E965dF"), // yvdusd3CRV
	common.HexToAddress("0x4f3E8F405CF5aFC05D68142F3783bDfE13811522"): common.HexToHash("0xFe39Ce91437C76178665D64d7a2694B0f6f17fE3"), // yvusdn3CRV
	common.HexToAddress("0x94e131324b6054c0D789b190b2dAC504e4361b53"): common.HexToHash("0xF6C9E9AF314982A4b38366f4AbfAa00595C5A6fC"), // yvust3CRV
	common.HexToAddress("0x2fE94ea3d5d4a175184081439753DE15AeF9d614"): common.HexToHash("0x7F83935EcFe4729c4Ea592Ab2bC1A32588409797"), // ycrvoBTC
	common.HexToAddress("0xC25a3A3b969415c80451098fa907EC722572917F"): common.HexToHash("0x5533ed0a3b83F70c3c4a1f69Ef5546D3D4713E44"), // yvcrvSUSD
	common.HexToAddress("0xb19059ebb43466C323583928285a49f558E572Fd"): common.HexToHash("0x46AFc2dfBd1ea0c0760CAD8262A5838e803A37e5"), // yvhCRV
	common.HexToAddress("0xDE5331AC4B3630f94853Ff322B66407e0D6331E8"): common.HexToHash("0x123964EbE096A920dae00Fb795FFBfA0c9Ff4675"), // ycrvpBTC
	common.HexToAddress("0x194eBd173F6cDacE046C53eACcE9B953F28411d1"): common.HexToHash("0x98B058b2CBacF5E99bC7012DF757ea7CFEbd35BC"), // yveursCRV
	common.HexToAddress("0x64eda51d3Ad40D56b9dFc5554E06F94e1Dd786Fd"): common.HexToHash("0x07FB4756f67bD46B748b16119E802F1f880fb2CC"), // ycrvtBTC
	common.HexToAddress("0x5B5CFE992AdAC0C9D48E05854B2d91C73a003858"): common.HexToHash("0x39546945695DCb1c037C836925B355262f551f55"), // yvhusd3CRV
	common.HexToAddress("0x410e3E86ef427e30B9235497143881f717d93c2A"): common.HexToHash("0xA8B1Cb4ed612ee179BDeA16CCa6Ba596321AE52D"), // ycrvbBTC
	common.HexToAddress("0xaA17A236F2bAdc98DDc0Cf999AbB47D47Fc0A6Cf"): common.HexToHash("0xE625F5923303f1CE7A43ACFEFd11fd12f30DbcA4"), // yvankrCRV
	common.HexToAddress("0x410e3E86ef427e30B9235497143881f717d93c2A"): common.HexToHash("0xA8B1Cb4ed612ee179BDeA16CCa6Ba596321AE52D"), // ycrvbBTC
	common.HexToAddress("0x49849C98ae39Fff122806C06791Fa73784FB3675"): common.HexToHash("0x5334e150B938dd2b6bd040D9c4a03Cff0cED3765"), // yvcrvRenWBTC
}

// vaultV2AddressHashMap represents token to vault v2 address hash map
var vaultV2AddressHashMap = map[common.Address]common.Hash{
	common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"): common.HexToHash("0x5f18C75AbDAe578b483E5F43f12a39cF75b973a9"), // yvUSDC
	common.HexToAddress("0x584bC13c7D411c00c01A62e8019472dE68768430"): common.HexToHash("0xe11ba472F74869176652C35D30dB89854b5ae84D"), // yvHEGIC
	common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"): common.HexToHash("0x19D3364A399d251E894aC732651be8B0E4e85001"), // yvDAI
	common.HexToAddress("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"): common.HexToHash("0xcB550A6D4C8e3517A939BC79d0c7093eb7cF56B5"), // yvWBTC
	common.HexToAddress("0xA3D87FffcE63B53E0d54fAa1cc983B7eB0b74A9c"): common.HexToHash("0x986b4AFF588a109c09B50A03f42E4110E29D353F"), // yveCRV
	common.HexToAddress("0x06325440D014e39736583c165C2963BA99fAf14E"): common.HexToHash("0xdCD90C7f6324cfa40d7169ef80b12031770B4325"), // yvsteCRV
}

// IsSupportedToken is supported token
func IsSupportedToken(tokenAddress common.Address) bool {
	_, ok1 := vaultAddressHashMap[tokenAddress]
	_, ok2 := vaultV2AddressHashMap[tokenAddress]
	return ok1 || ok2
}

// IsCorrectVault is correct vault
func IsCorrectVault(tokenAddress common.Address, vaultAddressHash common.Hash) bool {
	return vaultAddressHash == vaultAddressHashMap[tokenAddress] ||
		vaultAddressHash == vaultV2AddressHashMap[tokenAddress]
}
