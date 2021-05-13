package rewarder

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
)

// LoadGasRewardsTask load rewards task
type LoadGasRewardsTask struct {
	rootpath           string
	rewardAmount       decimal.Decimal
	tradingGasComboMap TradingGasComboMap
	nftCountMap        NftCountMap
	nftBoost           float64
	nftMaxBoost        float64

	filepath  string
	rewardMap RewardMap
}

// LoadRewardsFromFile load rewards from file
func (t *LoadGasRewardsTask) LoadRewardsFromFile() error {
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

// CalcRewards calc rewards
func (t *LoadGasRewardsTask) CalcRewards() error {
	log.Println("calculating rewards")

	t.rewardMap = make(RewardMap)

	// Boost reward by nft
	sum := decimal.Zero
	for account, amount := range t.tradingGasComboMap {
		if amount.IsZero() {
			return errors.New("combo amount is 0")
		}

		var boost float64 = 1
		if count, ok := t.nftCountMap[account]; ok {
			boost = math.Pow(t.nftBoost, float64(count))
			if boost > t.nftMaxBoost {
				boost = t.nftMaxBoost
			}
		}

		// Update reward
		boostAmount := amount.Mul(decimal.NewFromFloat(boost)).Truncate(furucombo.COMBODecimals)
		sum = sum.Add(boostAmount)
		t.rewardMap[account] = boostAmount
	}
	log.Printf("print boost combo sum: %s", sum.String())

	// Decline reward if reward is not enough
	ratio := t.rewardAmount.Div(sum)
	if ratio.LessThan(decimal.New(1, 0)) {
		sum := decimal.Zero
		winner := common.Address{}
		highestAmount := decimal.Zero
		for account, amount := range t.rewardMap {
			// Update reward
			declineAmount := amount.Mul(ratio).Truncate(furucombo.COMBODecimals)
			sum = sum.Add(declineAmount)
			t.rewardMap[account] = declineAmount
			if declineAmount.GreaterThan(highestAmount) {
				winner = account
				highestAmount = declineAmount
			}
		}
		if dust := t.rewardAmount.Sub(sum); dust.GreaterThan(decimal.Zero) {
			t.rewardMap[winner] = t.rewardMap[winner].Add(dust)
			log.Printf("give %s the dust amount %s COMBO", winner.String(), dust)
		}
		log.Printf("print decline combo sum: %s", sum.String())
	}

	return nil
}

// SaveRewardsToFile save rewards to file
func (t *LoadGasRewardsTask) SaveRewardsToFile() error {
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
func (t *LoadGasRewardsTask) Execute() error {
	if err := t.LoadRewardsFromFile(); err != nil {
		if err := t.CalcRewards(); err != nil {
			return err
		}

		if err := t.SaveRewardsToFile(); err != nil {
			return err
		}
	}

	return nil
}
