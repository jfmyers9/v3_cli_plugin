package commands

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
	AppName       string
	CliConnection plugin.CliConnection
}

func (c *CreateAppCommand) Perform() {
	util := utils.Utils{CliConnection: c.CliConnection}
	spaceGuid, err := util.GetTargetSpaceGuid()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	createAppPath := "/v3/apps"
	createAppBody := fmt.Sprintf(`{"name":"%s","space_guid":"%s"}`, c.AppName, spaceGuid)
	output, err := c.CliConnection.CliCommandWithoutTerminalOutput("curl", createAppPath, "-X", "POST", "-d", createAppBody)
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

	presenter := presenters.AppPresenter{App: app}
	fmt.Println(presenter.PresentApp())
}
