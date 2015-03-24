package apps

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jfmyers9/v3_cli_plugin/presenters"
	"github.com/jfmyers9/v3_cli_plugin/utils"
)

type ListAppsCommand struct {
	cliConnection plugin.CliConnection
}

func NewListAppsCommand(cliConnection plugin.CliConnection) ListAppsCommand {
	return ListAppsCommand{cliConnection: cliConnection}
}

func (c *ListAppsCommand) Perform() {
	util := utils.NewUtils(c.cliConnection)
	apps, err := util.ListAppsScopedToSpace()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	presenter := presenters.AppsPresenter{Apps: apps}
	fmt.Println(presenter.PresentApps())
}
