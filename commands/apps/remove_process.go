package apps

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jfmyers9/v3_cli_plugin/utils"
)

type RemoveProcessCommand struct {
	AppName       string
	ProcessType   string
	CliConnection plugin.CliConnection
}

func (c *RemoveProcessCommand) Perform() {
	util := utils.Utils{CliConnection: c.CliConnection}
	app, err := util.GetAppScopedToSpace(c.AppName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	process, err := util.GetProcessScopedToApp(app.Guid, c.ProcessType)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	removeProcessPath := fmt.Sprintf("/v3/apps/%s/processes", app.Guid)
	removeProcessBody := fmt.Sprintf(`{"process_guid":"%s"}`, process.Guid)

	_, err = c.CliConnection.CliCommandWithoutTerminalOutput("curl", removeProcessPath, "-X", "DELETE", "-d", removeProcessBody)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	deleteProcessPath := fmt.Sprintf("/v3/processes/%s", process.Guid)
	_, err = c.CliConnection.CliCommandWithoutTerminalOutput("curl", deleteProcessPath, "-X", "DELETE")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("OK")
}
