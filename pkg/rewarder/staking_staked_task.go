package rewarder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// LoadStakingStakedTask load staking staked task
type LoadStakingStakedTask struct {
	filepath        string
	stakingEventMap StakingEventMap
	startBlock      uint64

	stakingStakedMap StakingStakedMap
}

// LoadFromFile load from file
func (t *LoadStakingStakedTask) LoadFromFile() error {
	log.Printf("loading staking staked from ./%s", t.filepath)

	if _, err := os.Stat(t.filepath); err != nil {
		log.Println("staking staked file not found")
		return err
	}

	data, err := ioutil.ReadFile(t.filepath)
	if err != nil {
		return err
	}

	if err := jsonex.Unmarshal(data, &t.stakingStakedMap); err != nil {
		return err
	}

	return nil
}

// LoadFromDataset load from dataset
func (t *LoadStakingStakedTask) LoadFromDataset() error {
	log.Printf("loading staked from dataset")

	t.stakingStakedMap = StakingStakedMap{}

	for blockNumber, stakedAmountsMap := range t.stakingEventMap {
		if blockNumber < t.startBlock {
			for account, amounts := range stakedAmountsMap {
				for _, amount := range amounts {
					t.stakingStakedMap.Add(account, amount)
				}
			}
		}
	}

	t.stakingStakedMap.OmitZero()

	return nil
}

// SaveToFile save to file
func (t *LoadStakingStakedTask) SaveToFile() error {
	log.Printf("saving staked to ./%s", t.filepath)

	data, err := json.MarshalIndent(t.stakingStakedMap, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(t.filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}

// Execute execute
func (t *LoadStakingStakedTask) Execute() error {
	if err := t.LoadFromFile(); err != nil {
		if err := t.LoadFromDataset(); err != nil {
			return err
		}

		if err := t.SaveToFile(); err != nil {
			return err
		}
	}

	return nil
}
