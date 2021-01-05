package rewarder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"garage.dinngo.co/furucombo/mining-scripts/pkg/ethereum/furucombo"
)

// LoadRewardsTask load rewards task
type LoadRewardsTask struct {
	filepath        string
	rewardAmount    decimal.Decimal
	rewardWeightMap RewardWeightMap

	rewardMap RewardMap
}

// LoadRewardsFromFile load rewards from file
func (t *LoadRewardsTask) LoadRewardsFromFile() error {
	log.Printf("loading rewards from ./%s", t.filepath)

	if _, err := os.Stat(t.filepath); err != nil {
		log.Println("rewards file not found")
		return err
	}

	data, err := ioutil.ReadFile(t.filepath)
	if err != nil {
		return err
	}

	if err := jsonex.Unmarshal(data, &t.rewardMap); err != nil {
		return err
	}

	return nil
}

// CalcRewardsByWeight calc rewards by weight
func (t *LoadRewardsTask) CalcRewardsByWeight() error {
	log.Println("calculating rewards by weight")

	t.rewardMap = make(RewardMap)

	totalRewardAmount := decimal.Zero
	winner := common.Address{}
	highestRewardAmount := decimal.Zero

	for account, weighted := range t.rewardWeightMap {
		weight := weighted.GetWeight()

		if weight.IsZero() {
			continue
		}

		rewardAmount := weight.Mul(t.rewardAmount).Truncate(furucombo.COMBODecimals)
		if rewardAmount.GreaterThan(highestRewardAmount) {
			winner = account
			highestRewardAmount = rewardAmount
		}

		totalRewardAmount = totalRewardAmount.Add(rewardAmount)
		t.rewardMap[account] = rewardAmount
	}

	if dust := t.rewardAmount.Sub(totalRewardAmount); dust.GreaterThan(decimal.Zero) {
		t.rewardMap[winner] = t.rewardMap[winner].Add(dust)

		log.Printf("give %s the dust amount %s COMBO", winner.String(), dust)
	}

	return nil
}

// SaveRewardsToFile save rewards to file
func (t *LoadRewardsTask) SaveRewardsToFile() error {
	log.Printf("saving rewards to ./%s", t.filepath)

	data, err := json.MarshalIndent(t.rewardMap, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(t.filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}
