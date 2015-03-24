package apps

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jfmyers9/v3_cli_plugin/utils"
)

type UploadProcfileCommand struct {
	appName       string
	procfilePath  string
	cliConnection plugin.CliConnection
}

func NewUploadProcfileCommand(appName, procfilePath string, cliConnection plugin.CliConnection) UploadProcfileCommand {
	return UploadProcfileCommand{
		appName:       appName,
		procfilePath:  procfilePath,
		cliConnection: cliConnection,
	}
}

func (c *UploadProcfileCommand) Perform() {
	b, err := ioutil.ReadFile(c.procfilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	util := utils.NewUtils(c.cliConnection)
	app, err := util.GetAppScopedToSpace(c.appName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	procfilePath := fmt.Sprintf("/v3/apps/%s/procfile", app.Guid)
	output, err := c.cliConnection.CliCommandWithoutTerminalOutput("curl", procfilePath, "-X", "PUT", "-d", string(b))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(output[0])
}
