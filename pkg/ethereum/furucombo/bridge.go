package furucombo

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	ethereumex "github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
)

// event sigs
var BridgeDepositEventSig = common.HexToHash("0x3dad4cff57d5936bbc3bf46df565c9f1545c580c0efef238c7363402b42ba864")

// GetTokenBridgeEvents get token bridge events
func GetTokenBridgeEvents(from, to uint64) ([]types.Log, error) {
	events := make([]types.Log, 0)
	for _, address := range ProxyAddresses() {
		// new filter query
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(from)),
			ToBlock:   big.NewInt(int64(to)),
			Addresses: []common.Address{address},
			Topics: [][]common.Hash{
				{BridgeDepositEventSig}, // event sig
				{},                      // sender
				{},                      // token
			},
		}

		logs, err := ethereumex.Client().FilterLogs(context.Background(), query)
		if err != nil {
			return nil, err
		}
		events = append(events, logs...)
	}

	return events, nil
}
