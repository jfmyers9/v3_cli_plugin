package commands

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jfmyers9/v3_cli_plugin/presenters"
	"github.com/jfmyers9/v3_cli_plugin/utils"
)

type ListAppsCommand struct {
	CliConnection plugin.CliConnection
}

func (c *ListAppsCommand) Perform() {
	util := utils.Utils{CliConnection: c.CliConnection}
	apps, err := util.ListAppsScopedToSpace()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	presenter := presenters.AppsPresenter{Apps: apps}
	fmt.Println(presenter.PresentApps())
}
