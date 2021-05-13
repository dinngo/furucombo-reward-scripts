package rewarder

import "github.com/ethereum/go-ethereum/common"

// NftConfig nft config
type NftConfig struct {
	Ethereum []common.Address `json:"ethereum"`
	Matic    []common.Address `json:"matic"`
	Boost    float64          `json:"boost"`
	MaxBoost float64          `json:"maxBoost"`
}
