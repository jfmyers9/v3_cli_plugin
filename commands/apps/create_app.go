package apps

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jfmyers9/v3_cli_plugin/presenters"
	"github.com/jfmyers9/v3_cli_plugin/resources"
	"github.com/jfmyers9/v3_cli_plugin/utils"
)

type CreateAppCommand struct {
	appName       string
	cliConnection plugin.CliConnection
}

func NewCreateAppCommand(appName string, cliConnection plugin.CliConnection) CreateAppCommand {
	return CreateAppCommand{
		appName:       appName,
		cliConnection: cliConnection,
	}
}

func (c *CreateAppCommand) Perform() {
	util := utils.NewUtils(c.cliConnection)
	spaceGuid, err := util.GetTargetSpaceGuid()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	createAppPath := "/v3/apps"
	createAppBody := fmt.Sprintf(`{"name":"%s","space_guid":"%s"}`, c.appName, spaceGuid)
	output, err := c.cliConnection.CliCommandWithoutTerminalOutput("curl", createAppPath, "-X", "POST", "-d", createAppBody)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var app resources.V3App
	err = json.Unmarshal([]byte(output[0]), &app)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	presenter := presenters.NewAppPresenter(app)
	fmt.Println(presenter.Present())
}
