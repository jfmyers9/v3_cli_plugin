package apps

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jfmyers9/v3_cli_plugin/presenters"
	"github.com/jfmyers9/v3_cli_plugin/utils"
)

type GetAppCommand struct {
	appName       string
	cliConnection plugin.CliConnection
}

func NewGetAppCommand(appName string, cliConnection plugin.CliConnection) GetAppCommand {
	return GetAppCommand{
		appName:       appName,
		cliConnection: cliConnection,
	}
}

func (c *GetAppCommand) Perform() {
	util := utils.NewUtils(c.cliConnection)
	app, err := util.GetAppScopedToSpace(c.appName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	presenter := presenters.AppPresenter{App: app}
	fmt.Println(presenter.PresentApp())
}
