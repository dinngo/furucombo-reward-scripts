package rewarder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// SaveRewards save rewards to file
func (c *Config) SaveRewards() error {
	filepath := path.Join(c.RoundDir(), "rewards.json")

	if _, err := os.Stat(filepath); err == nil {
		return nil
	}

	log.Printf("saving rewards: ./%s", filepath)

	data, err := json.MarshalIndent(c.RewardMap, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath, append(data, '\n'), 0644); err != nil {
		return err
	}

	return nil
}
