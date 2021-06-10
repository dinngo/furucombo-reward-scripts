package commands

import (
	"flag"
	"log"
	"strings"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/rewarder"
)

// PolygonUsageCommand the command struct
type PolygonUsageCommand struct {
	ConfigPath string
}

// Synopsis the synopsis of command
func (c *PolygonUsageCommand) Synopsis() string {
	return "Generate polygon usage reward merkle files"
}

// Help the help of command
func (c *PolygonUsageCommand) Help() string {
	helpText := `
PolygonUsage: rewarder polygon usage
	Generate polygon usage reward merkle files

Options:
	-c     The path of config file.
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *PolygonUsageCommand) Run(args []string) int {
	f := flag.NewFlagSet("polygon-usage", flag.ContinueOnError)
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

	if config.PolygonEndBlock == 0 {
		if err := config.UpdatePolygonEndBlockToCurrentBlock(); err != nil {
			log.Println(err)
			return 1
		}
	}

	if err := config.Validates(rewarder.PolygonUsageRewarderRequiredFieldNames); err != nil {
		log.Println(err)
		return 1
	}

	if err := config.MakeRoundDir(); err != nil {
		log.Println(err)
		return 1
	}

	if err := config.SavePolygonBlocks(); err != nil {
		log.Println(err)
		return 1
	}

	if err := config.SaveTokenPrices(); err != nil {
		log.Println(err)
		return 1
	}

	polygonUsageRewarder := rewarder.NewPolygonUsageRewarder(config)

	if err := polygonUsageRewarder.LoadBridgeTxs(); err != nil {
		log.Println(err)
		return 1
	}

	if err := polygonUsageRewarder.LoadTradingLoyalty(); err != nil {
		log.Println(err)
		return 1
	}

	if err := polygonUsageRewarder.MakeUsageDir(); err != nil {
		log.Println(err)
		return 1
	}

	if err := polygonUsageRewarder.LoadRewards(); err != nil {
		log.Println(err)
		return 1
	}

	if err := polygonUsageRewarder.GenerateRewardsMerkleTree(); err != nil {
		log.Println(err)
		return 1
	}

	return 0
}
