package apps

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jfmyers9/v3_cli_plugin/utils"
)

type DeleteAppCommand struct {
	AppName       string
	CliConnection plugin.CliConnection
}

func (c *DeleteAppCommand) Perform() {
	util := utils.Utils{CliConnection: c.CliConnection}
	app, err := util.GetAppScopedToSpace(c.AppName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	deleteAppPath := fmt.Sprintf("/v3/apps/%s", app.Guid)
	_, err = c.CliConnection.CliCommandWithoutTerminalOutput("curl", deleteAppPath, "-X", "DELETE")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("OK")
}
