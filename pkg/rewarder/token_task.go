package rewarder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// LoadTokensTask load tokens task
type LoadTokensTask struct {
	filepath       string
	startTimestamp uint64
	endTimestamp   uint64

	tokenMap TokenMap
}

// LoadTokensFromFile load tokens from file
func (t *LoadTokensTask) LoadTokensFromFile() error {
	log.Printf("loading tokens from ./%s", t.filepath)

	if _, err := os.Stat(t.filepath); err != nil {
		log.Println("tokens file not found")
		return err
	}

	data, err := ioutil.ReadFile(t.filepath)
	if err != nil {
		return err
	}

	if err := jsonex.Unmarshal(data, &t.tokenMap); err != nil {
		return err
	}

	return nil
}

// LoadTokensWithPrices load tokens with prices
func (t *LoadTokensTask) LoadTokensWithPrices() error {
	log.Printf("getting tokens with whitelist")

	tokenWhitelist, err := GetTokenWhitelist()
	if err != nil {
		return err
	}

	if err := tokenWhitelist.GetTokenPrices(t.startTimestamp, t.endTimestamp); err != nil {
		return err
	}

	t.tokenMap = tokenWhitelist

	return nil
}

// SaveTokensToFile save tokens to file
func (t *LoadTokensTask) SaveTokensToFile() error {
	log.Printf("saving tokens to ./%s", t.filepath)

	data, err := json.MarshalIndent(t.tokenMap, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(t.filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}
