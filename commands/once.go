package commands

import (
	"flag"
	"log"
	"strings"

	"garage.dinngo.co/furucombo/mining-scripts/pkg/rewarder"
)

// OnceCommand the command struct
type OnceCommand struct {
	ConfigPath string
}

// Synopsis the synopsis of command
func (c *OnceCommand) Synopsis() string {
	return "Generate one-off reward merkle files"
}

// Help the help of command
func (c *OnceCommand) Help() string {
	helpText := `
Usage: rewarder once
	Generate one-off reward merkle files

Options:
	-c     The path of config file.
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *OnceCommand) Run(args []string) int {
	f := flag.NewFlagSet("once", flag.ContinueOnError)
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

	if err := config.Validates(rewarder.OnceRewarderRequiredFieldNames); err != nil {
		log.Println(err)
		return 1
	}

	if err := config.MakeRoundDir(); err != nil {
		log.Println(err)
		return 1
	}

	if err := config.SaveRewards(); err != nil {
		log.Println(err)
		return 1
	}

	onceRewarder := rewarder.NewOnceRewarder(config)

	if err := onceRewarder.GenerateRewardMerkleTree(); err != nil {
		log.Println(err)
		return 1
	}

	return 0
}
