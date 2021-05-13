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

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/ethereum/furucombo"
)

// LoadStakingsTask  load stakings task
type LoadStakingsTask struct {
	rootpath         string
	round            string
	duration         decimal.Decimal
	baseAmount       decimal.Decimal
	stakingStakedMap StakingStakedMap
	tradingRankMap   TradingRankMap  // consider trading rank in season 1
	tradingCountMap  TradingCountMap // consider trading count in season 2
	poolAddress      common.Address
	startBlock       uint64
	endBlock         uint64

	filepath   string
	stakingMap StakingMap
}

// LoadStakingsFromFile load stakings from file
func (t *LoadStakingsTask) LoadStakingsFromFile() error {
	t.filepath = path.Join(t.rootpath, "stakings.json")

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

// InitStakings init staking
func (t *LoadStakingsTask) InitStakings() error {
	t.stakingMap = make(StakingMap)

	log.Printf("init stakings with staked")

	for account, amount := range t.stakingStakedMap {
		t.stakingMap.Add(account, t.duration, t.baseAmount, amount)
	}

	// init with tradings
	log.Printf("init stakings with tradings")

	var tradingMap map[common.Address]int
	if len(t.tradingRankMap) > 0 { // season 1
		tradingMap = t.tradingRankMap
	} else if len(t.tradingCountMap) > 0 { // season 2
		tradingMap = t.tradingCountMap
	} else { // season 3
		return nil
	}

	for account := range tradingMap {
		if _, ok := t.stakingMap[account]; !ok {
			t.stakingMap.Add(account, t.duration, t.baseAmount, decimal.Zero)
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
			t.stakingMap.Add(account, t.duration, t.baseAmount, decimal.Zero)
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

// CalcStakingsWeightWithTrading calc stakings weight with trading
func (t *LoadStakingsTask) CalcStakingsWeightWithTrading() error {
	log.Println("calculating stakings weight")

	totalRankArea := decimal.Zero

	if len(t.tradingRankMap) > 0 { // season 1
		for account, rank := range t.tradingRankMap {
			t.stakingMap[account].RankArea = t.stakingMap[account].Area.Mul(decimal.NewFromInt(int64(rank)))
			totalRankArea = totalRankArea.Add(t.stakingMap[account].RankArea)
		}
	} else if len(t.tradingCountMap) > 0 { // season 2
		for account, count := range t.tradingCountMap {
			if count < 1 {
				return errors.New("unexpected trading count < 1")
			}
			// all accounts have the same rank if they did at least one tx
			t.stakingMap[account].RankArea = t.stakingMap[account].Area
			totalRankArea = totalRankArea.Add(t.stakingMap[account].RankArea)
		}
	} else { // season 3
		for account := range t.stakingMap {
			t.stakingMap[account].RankArea = t.stakingMap[account].Area
			totalRankArea = totalRankArea.Add(t.stakingMap[account].RankArea)
		}
	}

	if totalRankArea.IsZero() {
		log.Println("unexpected 0 total rank area")
		return nil
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

// Execute execute
func (t *LoadStakingsTask) Execute() error {
	if err := t.LoadStakingsFromFile(); err != nil {
		if err := t.InitStakings(); err != nil {
			return err
		}

		if err := t.UpdateStakingsByStakingEvents(); err != nil {
			return err
		}

		if err := t.CalcStakingsWeightWithTrading(); err != nil {
			return err
		}

		if err := t.SaveStakingsToFile(); err != nil {
			return err
		}
	}

	return nil
}
