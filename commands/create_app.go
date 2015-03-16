package commands

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
)

type CreateAppCommand struct {
	AppName       string
	SpaceGuid     string
	CliConnection plugin.CliConnection
}

func (c *CreateAppCommand) Perform() {
	createAppPath := "/v3/apps"
	createAppBody := fmt.Sprintf(`{"name":"%s","space_guid":"%s"}`, c.AppName, c.SpaceGuid)
	output, err := c.CliConnection.CliCommandWithoutTerminalOutput("curl", createAppPath, "-X", "POST", "-d", createAppBody)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(output[0])
}
