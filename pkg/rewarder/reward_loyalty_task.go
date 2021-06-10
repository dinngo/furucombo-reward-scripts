package rewarder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
)

// LoadLoyaltyRewardsTask load rewards task
type LoadLoyaltyRewardsTask struct {
	// Exterior
	rootpath          string
	baseReward        decimal.Decimal
	maxReward         decimal.Decimal
	tradingLoyaltyMap TradingLoyaltyMap

	// Intermedium
	filepath string

	// File
	rewardMap RewardMap
}

// LoadLoyaltyRewardsFromFile load loyalty rewards from file
func (t *LoadLoyaltyRewardsTask) LoadLoyaltyRewardsFromFile() error {
	t.filepath = path.Join(t.rootpath, "rewards.json")

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

// CalcLoyaltyRewards calc rewards by weight
func (t *LoadLoyaltyRewardsTask) CalcLoyaltyRewards() error {
	log.Println("calculating loyalty rewards")

	t.rewardMap = make(RewardMap)

	totalReward := decimal.Zero
	for user, loyalty := range t.tradingLoyaltyMap {
		reward := loyalty.Loyalty.Mul(t.baseReward).Truncate(furucombo.COMBODecimals)
		totalReward = totalReward.Add(reward)
		t.rewardMap[user] = reward
	}

	// Scale down rewards if exceeds max reward
	if totalReward.GreaterThan(t.maxReward) {
		newTotalReward := decimal.Zero
		newHighestReward := decimal.Zero
		winner := common.Address{}
		for user, reward := range t.rewardMap {
			newReward := reward.Mul(t.maxReward).Div(totalReward).Truncate(furucombo.COMBODecimals)
			newTotalReward = newTotalReward.Add(newReward)
			if newReward.GreaterThan(newHighestReward) {
				winner = user
				newHighestReward = newReward
			}
			t.rewardMap[user] = newReward
		}

		if dust := t.maxReward.Sub(newTotalReward); dust.GreaterThan(decimal.Zero) {
			t.rewardMap[winner] = t.rewardMap[winner].Add(dust)

			log.Printf("give %s the dust amount %s", winner.String(), dust)
		}
	}

	return nil
}

// SaveRewardsToFile save rewards to file
func (t *LoadLoyaltyRewardsTask) SaveRewardsToFile() error {
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

// Execute execute
func (t *LoadLoyaltyRewardsTask) Execute() error {
	if err := t.LoadLoyaltyRewardsFromFile(); err != nil {
		if err := t.CalcLoyaltyRewards(); err != nil {
			return err
		}

		if err := t.SaveRewardsToFile(); err != nil {
			return err
		}
	}

	return nil
}
