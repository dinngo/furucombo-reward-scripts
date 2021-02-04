package rewarder

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// StakingStakedMap represents user's staked amount
type StakingStakedMap map[common.Address]decimal.Decimal

// Add add staked amoount
func (m StakingStakedMap) Add(account common.Address, amount decimal.Decimal) {
	log.Printf("adding %s staked amount: %s", account.String(), amount)

	if _, ok := m[account]; ok {
		m[account] = m[account].Add(amount)
	} else {
		m[account] = amount
	}
}

// OmitZero omit zero
func (m StakingStakedMap) OmitZero() {
	for account, amount := range m {
		if amount.IsZero() {
			delete(m, account)
		}
	}
}
