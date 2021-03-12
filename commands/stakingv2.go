package commands

import (
	"flag"
	"log"
	"strings"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/rewarder"
)

// StakingV2Command the command struct
type StakingV2Command struct {
	ConfigPath string
}

// Synopsis the synopsis of command
func (c *StakingV2Command) Synopsis() string {
	return "Generate staking v2 reward merkle files"
}

// Help the help of command
func (c *StakingV2Command) Help() string {
	helpText := `
Usage: rewarder staking v2
	Generate staking v2 reward merkle files

Options:
	-c     The path of config file.
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *StakingV2Command) Run(args []string) int {
	f := flag.NewFlagSet("stakingv2", flag.ContinueOnError)
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

	if err := config.Validates(rewarder.StakingV2RewarderRequiredFieldNames); err != nil {
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

	if err := config.SavePools(); err != nil {
		log.Println(err)
		return 1
	}

	if err := config.SavePoolPrices(); err != nil {
		log.Println(err)
		return 1
	}

	stakingV2Rewarder := rewarder.NewStakingV2Rewarder(config)

	if err := stakingV2Rewarder.MakeStakingPoolDir(); err != nil {
		log.Println(err)
		return 1
	}

	if err := stakingV2Rewarder.LoadTradingCount(); err != nil {
		log.Println(err)
		return 1
	}

	if err := stakingV2Rewarder.LoadStakingsDataset(); err != nil {
		log.Println(err)
		return 1
	}

	if err := stakingV2Rewarder.LoadStakingsStaked(); err != nil {
		log.Println(err)
		return 1
	}

	if err := stakingV2Rewarder.LoadStakings(); err != nil {
		log.Println(err)
		return 1
	}

	if err := stakingV2Rewarder.LoadRewards(); err != nil {
		log.Println(err)
		return 1
	}

	if err := stakingV2Rewarder.GenerateRewardsMerkleTree(); err != nil {
		log.Println(err)
		return 1
	}

	return 0
}
