package ethereum

import "github.com/ethereum/go-ethereum/common"

// IsZeroHash is zero hash
func IsZeroHash(hash common.Hash) bool {
	return hash == common.Hash{}
}
