package commands

import (
	"fmt"

	"github.com/digitalonus/venera/localnet"
)

type InfoCommand struct {
	Subcommand string
}

func (c *InfoCommand) Help() string {
	return `
		Usage: venera info address 
		will get as much info from the current node and its connections
	`
}

func (c *InfoCommand) Run(args []string) int {
	if len(args) < 1 {
		fmt.Println(c.Help())
		return 0
	}

	switch args[0] {
	case "address":
		addresses, err := localnet.GetLocalAddresses()
		if err != nil {
			fmt.Println(err)
			return 1
		}
		for _, add := range addresses {
			fmt.Println(add)
		}
	default:
		c.Help()
	}
	return 0
}

func (c *InfoCommand) Synopsis() string {
	return "returns info of the node"
}
