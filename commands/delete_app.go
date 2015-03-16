package commands

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
)

type DeleteAppCommand struct {
	AppGuid       string
	CliConnection plugin.CliConnection
}

func (c *DeleteAppCommand) Perform() {
	deleteAppPath := fmt.Sprintf("/v3/apps/%s", c.AppGuid)
	_, err := c.CliConnection.CliCommandWithoutTerminalOutput("curl", deleteAppPath, "-X", "DELETE")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("OK")
}
