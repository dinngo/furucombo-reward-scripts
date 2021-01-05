package rewarder

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"garage.dinngo.co/furucombo/mining-scripts/pkg/ethereum/furucombo"
)

// LoadStakingsTask  load stakings task
type LoadStakingsTask struct {
	filepath    string
	round       int
	duration    decimal.Decimal
	tradingMap  TradingMap
	poolAddress common.Address
	baseAmount  decimal.Decimal
	startBlock  uint64
	endBlock    uint64

	stakingMap StakingMap
}

// LoadStakingsFromFile load stakings from file
func (t *LoadStakingsTask) LoadStakingsFromFile() error {
	log.Printf("loading stakings from ./%s", t.filepath)

	if _, err := os.Stat(t.filepath); err != nil {
		log.Println("stakings file not found")
		return err
	}

	data, err := ioutil.ReadFile(t.filepath)
	if err != nil {
		return err
	}

	if err := jsonex.Unmarshal(data, &t.stakingMap); err != nil {
		return err
	}

	return nil
}

// MakeStakingPoolDir make staking pool dir
func (t *LoadStakingsTask) MakeStakingPoolDir() error {
	poolDir := path.Dir(t.filepath)

	log.Printf("making staking pool dir: ./%s/", poolDir)

	if err := os.MkdirAll(poolDir, os.ModePerm); err != nil {
		return err
	}

	return nil
}

// InitStakings init staking
func (t *LoadStakingsTask) InitStakings() error {
	t.stakingMap = make(StakingMap)

	// init with pre round stakings
	if t.round > 0 {
		preRoundFilepath := strings.Replace(t.filepath, fmt.Sprintf("/%d/", t.round), fmt.Sprintf("/%d/", t.round-1), 1)

		log.Printf("init stakings with pre-round stakings: ./%s", preRoundFilepath)

		if _, err := os.Stat(preRoundFilepath); err != nil {
			return err
		}

		data, err := ioutil.ReadFile(preRoundFilepath)
		if err != nil {
			return err
		}

		var preStakingMap StakingMap
		if err := jsonex.Unmarshal(data, &preStakingMap); err != nil {
			return err
		}

		for account, staking := range preStakingMap {
			t.stakingMap.Add(account, t.duration, t.baseAmount, staking)
		}
	}

	// init with tradings
	log.Printf("init stakings with tradings")

	for account := range t.tradingMap {
		if _, ok := t.stakingMap[account]; !ok {
			t.stakingMap.Add(account, t.duration, t.baseAmount, nil)
		}
	}

	return nil
}

// UpdateStakingsByStakingEvents set stakings by staking events
func (t *LoadStakingsTask) UpdateStakingsByStakingEvents() error {
	log.Printf("getting pool %s staking events", t.poolAddress.String())

	events, err := furucombo.GetStakingEvents(t.poolAddress, t.startBlock, t.endBlock)
	if err != nil {
		return err
	}

	log.Printf("found %d staking events", len(events))

	for _, event := range events {
		account := common.BytesToAddress(event.Topics[2].Bytes())

		log.Printf("found %s staking event at %d block", account.String(), event.BlockNumber)

		if _, ok := t.stakingMap[account]; !ok {
			t.stakingMap.Add(account, t.duration, t.baseAmount, nil)
		}

		duration := decimal.NewFromInt(int64(t.endBlock - event.BlockNumber))
		amount := decimal.NewFromBigInt(new(big.Int).SetBytes(event.Data), -18)
		area := duration.Mul(amount)

		var eventName string
		switch event.Topics[0] {
		case furucombo.StakingStakedEventSig:
			eventName = "Staked"

			t.stakingMap[account].EndAmount = t.stakingMap[account].EndAmount.Add(amount)
			t.stakingMap[account].Area = t.stakingMap[account].Area.Add(area)

		case furucombo.StakingUnstakedEventSig:
			eventName = "Unstaked"

			if area.GreaterThan(t.stakingMap[account].Area) {
				return errors.New("unexpected staking area < 0 ")
			}

			t.stakingMap[account].EndAmount = t.stakingMap[account].EndAmount.Sub(amount)
			t.stakingMap[account].Area = t.stakingMap[account].Area.Sub(area)

		default:
			return fmt.Errorf("unexpected staking event: %s", event.Topics[0].String())
		}

		log.Printf("event: %s duration: %s amount: %s area: %s", eventName, duration, amount, area)
	}

	return nil
}

// CalcStakingsWeightWithTradingRank calc stakings weight with trading rank
func (t *LoadStakingsTask) CalcStakingsWeightWithTradingRank() error {
	log.Println("calculating stakings weight")

	totalRankArea := decimal.Zero
	for account, trading := range t.tradingMap {
		t.stakingMap[account].RankArea = t.stakingMap[account].Area.Mul(decimal.NewFromInt(int64(trading.Rank)))
		totalRankArea = totalRankArea.Add(t.stakingMap[account].RankArea)
	}

	if totalRankArea.IsZero() {
		return errors.New("unexpected 0 total rank area")
	}

	for account := range t.stakingMap {
		weight, _ := t.stakingMap[account].RankArea.QuoRem(totalRankArea, 27)
		t.stakingMap[account].Weight = weight
	}

	return nil
}

// SaveStakingsToFile save stakings to file
func (t *LoadStakingsTask) SaveStakingsToFile() error {
	log.Printf("saving stakings to ./%s", t.filepath)

	data, err := json.MarshalIndent(t.stakingMap, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(t.filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}
