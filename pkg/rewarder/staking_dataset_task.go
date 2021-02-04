package rewarder

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum"
	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
)

// LoadStakingDatasetTask  load staking dataset task
type LoadStakingDatasetTask struct {
	filepath    string
	poolAddress common.Address
	startBlock  uint64

	endBlock uint64

	stakingDataset StakingDataset
}

// LoadFromFileOrCreate load from file or create
func (t *LoadStakingDatasetTask) LoadFromFileOrCreate() error {
	if _, err := os.Stat(t.filepath); err == nil {
		log.Printf("loading staking dataset from ./%s", t.filepath)

		data, err := ioutil.ReadFile(t.filepath)
		if err != nil {
			return err
		}

		if err := jsonex.Unmarshal(data, &t.stakingDataset); err != nil {
			return err
		}

		t.startBlock = t.stakingDataset.EndBlock + 1

		return nil
	}

	t.stakingDataset = StakingDataset{
		StartBlock: t.startBlock,
		EventMap:   StakingEventMap{},
	}

	return nil
}

// GetEndBlock get end block
func (t *LoadStakingDatasetTask) GetEndBlock() error {
	endBlock, err := ethereum.Client().BlockNumber(context.Background())
	if err != nil {
		return err
	}

	t.stakingDataset.EndBlock = endBlock
	t.endBlock = endBlock

	return nil
}

// UpdateDataset update dataset
func (t *LoadStakingDatasetTask) UpdateDataset() error {
	log.Printf("getting pool %s staking events", t.poolAddress)

	events, err := furucombo.GetStakingEvents(t.poolAddress, t.startBlock, t.endBlock)
	if err != nil {
		return err
	}

	log.Printf("found %d staking events from %d to %d", len(events), t.startBlock, t.endBlock)

	for _, event := range events {
		account := common.BytesToAddress(event.Topics[2].Bytes())
		amount := decimal.NewFromBigInt(new(big.Int).SetBytes(event.Data), -18)

		if event.Topics[0] == furucombo.StakingUnstakedEventSig {
			amount = amount.Neg()
		}

		t.stakingDataset.EventMap.Add(event.BlockNumber, account, amount)
	}

	return nil
}

// SaveToFile save to file
func (t *LoadStakingDatasetTask) SaveToFile() error {
	log.Printf("saving staking dataset to ./%s", t.filepath)

	data, err := json.MarshalIndent(t.stakingDataset, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(t.filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}

// Execute execute
func (t *LoadStakingDatasetTask) Execute() error {
	if err := t.LoadFromFileOrCreate(); err != nil {
		return err
	}

	if err := t.GetEndBlock(); err != nil {
		return err
	}

	if err := t.UpdateDataset(); err != nil {
		return err
	}

	if err := t.SaveToFile(); err != nil {
		return err
	}

	return nil
}
