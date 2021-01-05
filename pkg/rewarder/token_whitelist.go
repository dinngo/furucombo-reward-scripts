package rewarder

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/ethereum/go-ethereum/common"
)

var (
	tokenWhitelistFilepath = path.Join("config", "token_whitelist.json")
	tokenWhitelist         TokenMap
)

// GetTokenWhitelist get token whitelist
func GetTokenWhitelist() (TokenMap, error) {
	if len(tokenWhitelist) == 0 {
		if _, err := os.Stat(tokenWhitelistFilepath); err != nil {
			return nil, err
		}

		data, err := ioutil.ReadFile(tokenWhitelistFilepath)
		if err != nil {
			return nil, err
		}

		if err := jsonex.Unmarshal(data, &tokenWhitelist); err != nil {
			return nil, err
		}
	}

	return tokenWhitelist, nil
}

// IsTokenListed is token listed
func IsTokenListed(tokenAddress common.Address) bool {
	tokenMap, err := GetTokenWhitelist()
	if err != nil {
		return false
	}

	_, ok := tokenMap[tokenAddress]

	return ok
}
