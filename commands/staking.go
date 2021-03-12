package commands

import (
	"flag"
	"log"
	"strings"

	"github.com/dinngodev/furucombo-reward-scripts/pkg/rewarder"
)

// StakingCommand the command struct
type StakingCommand struct {
	ConfigPath string
}

// Synopsis the synopsis of command
func (c *StakingCommand) Synopsis() string {
	return "Generate staking reward merkle files"
}

// Help the help of command
func (c *StakingCommand) Help() string {
	helpText := `
Usage: rewarder staking
	Generate staking reward merkle files

Options:
	-c     The path of config file.
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *StakingCommand) Run(args []string) int {
	f := flag.NewFlagSet("staking", flag.ContinueOnError)
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

	if err := config.Validates(rewarder.StakingRewarderRequiredFieldNames); err != nil {
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

	if err := config.SavePools(); err != nil {
		log.Println(err)
		return 1
	}

	if err := config.SavePoolPrices(); err != nil {
		log.Println(err)
		return 1
	}

	stakingRewarder := rewarder.NewStakingRewarder(config)

	if err := stakingRewarder.MakeStakingPoolDir(); err != nil {
		log.Println(err)
		return 1
	}

	if err := stakingRewarder.LoadTxs(); err != nil {
		log.Println(err)
		return 1
	}

	if err := stakingRewarder.LoadTradingVolumes(); err != nil {
		log.Println(err)
		return 1
	}

	if err := stakingRewarder.LoadTradingRanks(); err != nil {
		log.Println(err)
		return 1
	}

	if err := stakingRewarder.LoadStakingsDataset(); err != nil {
		log.Println(err)
		return 1
	}

	if err := stakingRewarder.LoadStakingsStaked(); err != nil {
		log.Println(err)
		return 1
	}

	if err := stakingRewarder.LoadStakings(); err != nil {
		log.Println(err)
		return 1
	}

	if err := stakingRewarder.LoadRewards(); err != nil {
		log.Println(err)
		return 1
	}

	if err := stakingRewarder.GenerateRewardsMerkleTree(); err != nil {
		log.Println(err)
		return 1
	}

	return 0
}
