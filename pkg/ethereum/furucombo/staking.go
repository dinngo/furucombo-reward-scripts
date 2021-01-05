package furucombo

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	ethereumex "garage.dinngo.co/furucombo/mining-scripts/pkg/ethereum"
)

// event sigs
var (
	StakingStakedEventSig   = common.HexToHash("0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7")
	StakingUnstakedEventSig = common.HexToHash("0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3")
)

// GetStakingEvents get staking events
func GetStakingEvents(stakingAddress common.Address, from, to uint64) ([]types.Log, error) {
	// new filter query
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from)),
		ToBlock:   big.NewInt(int64(to)),
		Addresses: []common.Address{stakingAddress},
		Topics: [][]common.Hash{
			{StakingStakedEventSig, StakingUnstakedEventSig}, // event sig
			{}, // sender
			{}, // onBehalfOf
		},
	}

	events, err := ethereumex.Client().FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}

	return events, nil
}
