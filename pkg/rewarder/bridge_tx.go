package rewarder

import (
	"github.com/ethereum/go-ethereum/common"
)

// TradingCountMap represents address to trading count map
type BridgeTxMap map[common.Address][]common.Hash
