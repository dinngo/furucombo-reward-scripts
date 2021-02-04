package ethereum

import (
	"strings"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var tokenAddressMap = map[string]common.Address{
	"DAI":  common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"),
	"WETH": common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
	"eETH": common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
}

// Token token info
type Token struct {
	Address  common.Address `json:"address"`
	Decimals uint8          `json:"decimals"`
	Symbol   string         `json:"symbol"`
}

// IsToken is token
func IsToken(symbol string, address common.Address) bool {
	return address == tokenAddressMap[symbol]
}

// GetTokenAddress get token
func GetTokenAddress(symbol string) common.Address {
	return tokenAddressMap[symbol]
}

// GetToken get token info
func GetToken(tokenAddress common.Address) (*Token, error) {
	parsed, err := abi.JSON(strings.NewReader(erc20.ERC20ContractABI))
	if err != nil {
		return nil, err
	}

	decimalsCallData, err := parsed.Pack("decimals")
	if err != nil {
		return nil, err
	}

	symbolCallData, err := parsed.Pack("symbol")
	if err != nil {
		return nil, err
	}

	calls := []Struct0{
		{Target: tokenAddress, CallData: decimalsCallData},
		{Target: tokenAddress, CallData: symbolCallData},
	}

	contract, err := NewMulticallContract(MulticallContractAddress, Client())
	if err != nil {
		return nil, err
	}

	contractRaw := MulticallContractRaw{Contract: contract}

	var result []interface{}
	if err := contractRaw.Call(nil, &result, "aggregate", calls); err != nil {
		return nil, err
	}

	returnData := result[1].([][]byte)

	decimalsDeocded, err := parsed.Methods["decimals"].Outputs.Unpack(returnData[0])
	if err != nil {
		return nil, err
	}

	symbolDecoded, err := parsed.Methods["symbol"].Outputs.Unpack(returnData[1])
	if err != nil {
		return nil, err
	}

	return &Token{
		Address:  tokenAddress,
		Decimals: decimalsDeocded[0].(uint8),
		Symbol:   symbolDecoded[0].(string),
	}, nil
}
