package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jfmyers9/v3_cli_plugin/utils"
)

type UploadProcfileCommand struct {
	AppName       string
	ProcfilePath  string
	CliConnection plugin.CliConnection
}

func (c *UploadProcfileCommand) Perform() {
	b, err := ioutil.ReadFile(c.ProcfilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	util := utils.Utils{CliConnection: c.CliConnection}
	app, err := util.GetAppScopedToSpace(c.AppName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	procfilePath := fmt.Sprintf("/v3/apps/%s/procfile", app.Guid)
	output, err := c.CliConnection.CliCommandWithoutTerminalOutput("curl", procfilePath, "-X", "PUT", "-d", string(b))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(output[0])
}
