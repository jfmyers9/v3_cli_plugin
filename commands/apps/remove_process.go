package apps

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jfmyers9/v3_cli_plugin/utils"
)

type RemoveProcessCommand struct {
	appName       string
	processType   string
	cliConnection plugin.CliConnection
}

func NewRemoveProcessCommand(appName, processType string, cliConnection plugin.CliConnection) RemoveProcessCommand {
	return RemoveProcessCommand{
		appName:       appName,
		processType:   processType,
		cliConnection: cliConnection,
	}
}

func (c *RemoveProcessCommand) Perform() {
	util := utils.NewUtils(c.cliConnection)
	app, err := util.GetAppScopedToSpace(c.appName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	process, err := util.GetProcessScopedToApp(app.Guid, c.processType)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	removeProcessPath := fmt.Sprintf("/v3/apps/%s/processes", app.Guid)
	removeProcessBody := fmt.Sprintf(`{"process_guid":"%s"}`, process.Guid)

	_, err = c.cliConnection.CliCommandWithoutTerminalOutput("curl", removeProcessPath, "-X", "DELETE", "-d", removeProcessBody)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	deleteProcessPath := fmt.Sprintf("/v3/processes/%s", process.Guid)
	_, err = c.cliConnection.CliCommandWithoutTerminalOutput("curl", deleteProcessPath, "-X", "DELETE")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("OK")
}
