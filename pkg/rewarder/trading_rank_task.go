package rewarder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/shopspring/decimal"
)

// LoadTradingRanksTask load trading ranks task
type LoadTradingRanksTask struct {
	rootpath         string
	volumeCap        decimal.Decimal
	tradingVolumeMap TradingVolumeMap

	filepath       string
	tradingRankMap TradingRankMap
}

// LoadFromFile load from file
func (t *LoadTradingRanksTask) LoadFromFile() error {
	t.filepath = path.Join(t.rootpath, "trading_ranks.json")

	log.Printf("loading trading ranks from ./%s", t.filepath)

	if _, err := os.Stat(t.filepath); err != nil {
		log.Println("trading ranks file not found")
		return err
	}

	data, err := ioutil.ReadFile(t.filepath)
	if err != nil {
		return err
	}

	if err := jsonex.Unmarshal(data, &t.tradingRankMap); err != nil {
		return err
	}

	return nil
}

// RankTradingVolumes rank trading volumes
func (t *LoadTradingRanksTask) RankTradingVolumes() {
	t.tradingRankMap = make(TradingRankMap)

	for account, curVolume := range t.tradingVolumeMap {
		// set volume cap if exceeds
		if curVolume.GreaterThan(t.volumeCap) {
			curVolume = t.volumeCap
		}

		rank := 1
		for _, volume := range t.tradingVolumeMap {
			// set volume cap if exceeds
			if volume.GreaterThan(t.volumeCap) {
				volume = t.volumeCap
			}

			if curVolume.GreaterThan(volume) {
				rank++
			}
		}

		t.tradingRankMap[account] = rank
	}
}

// SaveToFile save trading ranks to file
func (t *LoadTradingRanksTask) SaveToFile() error {
	log.Printf("saving trading ranks to ./%s", t.filepath)

	data, err := json.MarshalIndent(t.tradingRankMap, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(t.filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}

// Execute execute
func (t *LoadTradingRanksTask) Execute() error {
	if err := t.LoadFromFile(); err != nil {
		t.RankTradingVolumes()

		if err := t.SaveToFile(); err != nil {
			return err
		}
	}

	return nil
}
