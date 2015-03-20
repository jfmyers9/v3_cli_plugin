package processes

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jfmyers9/v3_cli_plugin/presenters"
	"github.com/jfmyers9/v3_cli_plugin/resources"
	"github.com/jfmyers9/v3_cli_plugin/utils"
)

type CreateProcessCommand struct {
	ProcessType   string
	CliConnection plugin.CliConnection
}

func (c *CreateProcessCommand) Perform() {
	util := utils.Utils{CliConnection: c.CliConnection}
	spaceGuid, err := util.GetTargetSpaceGuid()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	createProcessPath := "/v3/processes"
	createProcessBody := fmt.Sprintf(`{"type":"%s","space_guid":"%s"}`, c.ProcessType, spaceGuid)
	output, err := c.CliConnection.CliCommandWithoutTerminalOutput("curl", createProcessPath, "-X", "POST", "-d", createProcessBody)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var process resources.Process
	err = json.Unmarshal([]byte(output[0]), &process)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	presenter := presenters.ProcessPresenter{Process: process}
	fmt.Println(presenter.Present())
}
