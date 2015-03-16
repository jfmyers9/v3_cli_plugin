package commands

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jfmyers9/v3_cli_plugin/presenters"
	"github.com/jfmyers9/v3_cli_plugin/utils"
)

type GetAppCommand struct {
	AppName       string
	CliConnection plugin.CliConnection
}

func (c *GetAppCommand) Perform() {
	util := utils.Utils{CliConnection: c.CliConnection}
	app, err := util.GetAppScopedToSpace(c.AppName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	presenter := presenters.AppPresenter{App: app}
	fmt.Println(presenter.PresentApp())
}
