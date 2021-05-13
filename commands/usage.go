package commands

import (
	"flag"
	"log"
	"strings"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/rewarder"
)

// UsageCommand the command struct
type UsageCommand struct {
	ConfigPath string
}

// Synopsis the synopsis of command
func (c *UsageCommand) Synopsis() string {
	return "Generate usage reward merkle files"
}

// Help the help of command
func (c *UsageCommand) Help() string {
	helpText := `
Usage: rewarder usage
	Generate usage reward merkle files

Options:
	-c     The path of config file.
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *UsageCommand) Run(args []string) int {
	f := flag.NewFlagSet("usage", flag.ContinueOnError)
	f.StringVar(&c.ConfigPath, "c", "", "c")
	if err := f.Parse(args); err != nil {
		log.Println(err)
		return 1
	}

	if len(c.ConfigPath) == 0 {
		log.Println("config is required")
		return 1
	}

	config, err := rewarder.NewConfig(c.ConfigPath)
	if err != nil {
		log.Println(err)
		return 1
	}

	if err := config.LogToFile(); err != nil {
		log.Println(err)
		return 1
	}

	if config.EndBlock == 0 {
		if err := config.UpdateEndBlockToCurrentBlock(); err != nil {
			log.Println(err)
			return 1
		}
	}

	if config.MaticEndBlock == 0 {
		if err := config.UpdateMaticEndBlockToCurrentBlock(); err != nil {
			log.Println(err)
			return 1
		}
	}

	if err := config.Validates(rewarder.UsageRewarderRequiredFieldNames); err != nil {
		log.Println(err)
		return 1
	}

	if err := config.MakeRoundDir(); err != nil {
		log.Println(err)
		return 1
	}

	if err := config.GetBlockTimestamps(); err != nil {
		log.Println(err)
		return 1
	}

	if err := config.SaveBlocks(); err != nil {
		log.Println(err)
		return 1
	}

	usageRewarder := rewarder.NewUsageRewarder(config)

	if err := usageRewarder.LoadTradingGasCombo(); err != nil {
		log.Println(err)
		return 1
	}

	if err := usageRewarder.LoadNftCounts(); err != nil {
		log.Println(err)
		return 1
	}

	if err := usageRewarder.LoadGasRewards(); err != nil {
		log.Println(err)
		return 1
	}

	if err := usageRewarder.GenerateRewardsMerkleTree(); err != nil {
		log.Println(err)
		return 1
	}

	return 0
}
