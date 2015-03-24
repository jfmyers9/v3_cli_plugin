package apps

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jfmyers9/v3_cli_plugin/utils"
)

type DeleteAppCommand struct {
	appName       string
	cliConnection plugin.CliConnection
}

func NewDeleteAppCommand(appName string, cliConnection plugin.CliConnection) DeleteAppCommand {
	return DeleteAppCommand{
		appName:       appName,
		cliConnection: cliConnection,
	}
}

func (c *DeleteAppCommand) Perform() {
	util := utils.NewUtils(c.cliConnection)
	app, err := util.GetAppScopedToSpace(c.appName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	deleteAppPath := fmt.Sprintf("/v3/apps/%s", app.Guid)
	_, err = c.cliConnection.CliCommandWithoutTerminalOutput("curl", deleteAppPath, "-X", "DELETE")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("OK")
}
