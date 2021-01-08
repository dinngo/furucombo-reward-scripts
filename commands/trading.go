package commands

import (
	"flag"
	"log"
	"strings"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/rewarder"
)

// TradingCommand the command struct
type TradingCommand struct {
	ConfigPath string
}

// Synopsis the synopsis of command
func (c *TradingCommand) Synopsis() string {
	return "Generate trading reward merkle files"
}

// Help the help of command
func (c *TradingCommand) Help() string {
	helpText := `
Usage: rewarder trading
	Generate trading reward merkle files

Options:
	-c     The path of config file.
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *TradingCommand) Run(args []string) int {
	f := flag.NewFlagSet("trading", flag.ContinueOnError)
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

	if err := config.Validates(rewarder.TradingRewarderRequiredFieldNames); err != nil {
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

	if err := config.LoadCubeFinders(); err != nil {
		log.Println(err)
		return 1
	}

	if err := config.SaveBlocks(); err != nil {
		log.Println(err)
		return 1
	}

	tradingRewarder := rewarder.NewTradingRewarder(config)

	if err := tradingRewarder.LoadTokens(); err != nil {
		log.Println(err)
		return 1
	}

	if err := tradingRewarder.LoadTxs(); err != nil {
		log.Println(err)
		return 1
	}

	if err := tradingRewarder.LoadTradings(); err != nil {
		log.Println(err)
		return 1
	}

	if err := tradingRewarder.LoadRewards(); err != nil {
		log.Println(err)
		return 1
	}

	if err := tradingRewarder.GenerateRewardsMerkleTree(); err != nil {
		log.Println(err)
		return 1
	}

	return 0
}
