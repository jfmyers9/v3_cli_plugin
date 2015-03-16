package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jfmyers9/v3_cli_plugin/commands"
	"github.com/jfmyers9/v3_cli_plugin/utils"
)

const createAppString = "create-v3-app"
const deleteAppString = "delete-v3-app"

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
			{
				Name:     deleteAppString,
				HelpText: "This command deletes a v3 app.",
				UsageDetails: plugin.Usage{
					Usage: fmt.Sprintf("cf %s [-f] app-name", deleteAppString),
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
	} else if args[0] == deleteAppString && (len(args) == 2 || len(args) == 3 && args[1] == "-f") {
		var appName string
		var forceFlag string
		if len(args) == 2 {
			appName = args[1]
		} else {
			appName = args[2]
			forceFlag = args[1]
		}
		c.deleteApp(cliConnection, forceFlag, appName)
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
		os.Exit(1)
	}

	createCommand := commands.CreateAppCommand{
		AppName:       appName,
		SpaceGuid:     spaceGuid,
		CliConnection: cliConnection,
	}
	createCommand.Perform()
}

func (c *V3Cli) deleteApp(cliConnection plugin.CliConnection, forceFlag string, appName string) {
	force := forceFlag == "-f"
	if !force {
		var confirmation string
		fmt.Printf("Are you sure? ")
		fmt.Scanf("%s", &confirmation)

		if confirmation != "yes" && confirmation != "y" {
			fmt.Println("Cancelled")
			os.Exit(1)
		}
	}

	appGuid, err := utils.GetAppGuid(cliConnection, appName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	deleteCommand := commands.DeleteAppCommand{
		AppGuid:       appGuid,
		CliConnection: cliConnection,
	}
	deleteCommand.Perform()
}
