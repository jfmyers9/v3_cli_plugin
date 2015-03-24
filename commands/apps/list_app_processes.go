package apps

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jfmyers9/v3_cli_plugin/presenters"
	"github.com/jfmyers9/v3_cli_plugin/utils"
)

type ListAppProcessesCommand struct {
	appName       string
	cliConnection plugin.CliConnection
}

func NewListAppProcessesCommand(appName string, cliConnection plugin.CliConnection) ListAppProcessesCommand {
	return ListAppProcessesCommand{
		appName:       appName,
		cliConnection: cliConnection,
	}
}

func (c *ListAppProcessesCommand) Perform() {
	util := utils.NewUtils(c.cliConnection)
	app, err := util.GetAppScopedToSpace(c.appName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	processes, err := util.ListProcessesScopedToApp(app.Guid)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	presenter := presenters.NewProcessesPresenter(processes)
	fmt.Println(presenter.Present())
}
