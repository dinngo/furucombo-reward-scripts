package commands

import (
	"flag"
	"log"
	"strings"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/rewarder"
)

// BonusCommand the command struct
type BonusCommand struct {
	ConfigPath string
}

// Synopsis the synopsis of command
func (c *BonusCommand) Synopsis() string {
	return "Generate one-off reward merkle files"
}

// Help the help of command
func (c *BonusCommand) Help() string {
	helpText := `
Usage: rewarder bonus
	Generate bonus reward merkle files

Options:
	-c     The path of config file.
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *BonusCommand) Run(args []string) int {
	f := flag.NewFlagSet("bonus", flag.ContinueOnError)
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

	if err := config.MakeRoundDir(); err != nil {
		log.Println(err)
		return 1
	}

	if err := config.Validates(rewarder.BonusRewarderRequiredFieldNames); err != nil {
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

	if err := config.SavePool(); err != nil {
		log.Println(err)
		return 1
	}

	bonusRewarder := rewarder.NewBonusRewarder(config)

	if err := bonusRewarder.LoadTxs(); err != nil {
		log.Println(err)
		return 1
	}

	if err := bonusRewarder.LoadTradingVolumes(); err != nil {
		log.Println(err)
		return 1
	}

	if err := bonusRewarder.LoadTradingRanks(); err != nil {
		log.Println(err)
		return 1
	}

	if err := bonusRewarder.LoadStakingDataset(); err != nil {
		log.Println(err)
		return 1
	}

	if err := bonusRewarder.LoadStakingStaked(); err != nil {
		log.Println(err)
		return 1
	}

	if err := bonusRewarder.LoadStakings(); err != nil {
		log.Println(err)
		return 1
	}

	if err := bonusRewarder.LoadRewards(); err != nil {
		log.Println(err)
		return 1
	}

	if err := bonusRewarder.GenerateRewardsMerkleTree(); err != nil {
		log.Println(err)
		return 1
	}

	return 0
}
