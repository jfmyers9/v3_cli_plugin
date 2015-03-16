package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jfmyers9/v3_cli_plugin/commands"
	"github.com/jfmyers9/v3_cli_plugin/utils"
)

const createAppString = "create-v3-app"

type V3Cli struct{}

func (c *V3Cli) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "V3Cli",
		Commands: []plugin.Command{
			{
				Name:     createAppString,
				HelpText: "This command creates a v3 app.",
				UsageDetails: plugin.Usage{
					Usage: fmt.Sprintf("cf %s app-name", createAppString),
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(V3Cli))
}

func (c *V3Cli) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == createAppString && len(args) == 2 {
		appName := args[1]
		c.createApp(cliConnection, appName)
	} else {
		c.showUsage(args)
	}
}

func (c *V3Cli) showUsage(args []string) {
	for _, cmd := range c.GetMetadata().Commands {
		if cmd.Name == args[0] {
			fmt.Println(fmt.Sprintf("Usage: %s", cmd.UsageDetails.Usage))
		}
	}
}

func (c *V3Cli) createApp(cliConnection plugin.CliConnection, appName string) {
	spaceGuid, err := utils.GetTargetSpaceGuid(cliConnection)
	if err != nil {
		fmt.Println(err)
		return
	}

	createCommand := commands.CreateAppCommand{
		AppName:       appName,
		SpaceGuid:     spaceGuid,
		CliConnection: cliConnection,
	}
	createCommand.Perform()
}
