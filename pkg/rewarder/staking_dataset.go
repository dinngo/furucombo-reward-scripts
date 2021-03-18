package rewarder

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// StakingDataset struct
type StakingDataset struct {
	StartBlock uint64          `json:"startBlock"`
	EndBlock   uint64          `json:"endBlock"`
	EventMap   StakingEventMap `json:"events"`
}

// StakingEventMap represents user's staking data
type StakingEventMap map[uint64]map[common.Address][]decimal.Decimal

// Add add staking
func (m StakingEventMap) Add(block uint64, account common.Address, amount decimal.Decimal) {
	log.Printf("adding %s staking event", account.String())

	if _, ok := m[block]; ok {
		m[block][account] = append(m[block][account], amount)
	} else {
		m[block] = map[common.Address][]decimal.Decimal{
			account: {amount},
		}
	}

	log.Printf("block: %d account: %s amount: %s", block, account, amount)
}
