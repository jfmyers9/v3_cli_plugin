package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jfmyers9/v3_cli_plugin/commands/apps"
)

const (
	createAppString      = "create-v3-app"
	deleteAppString      = "delete-v3-app"
	getAppString         = "v3-app"
	listAppString        = "v3-apps"
	uploadProcfileString = "procfile"
	removeProcessString  = "remove-process"
	listProcessesString  = "processes"
)

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
			{
				Name:     getAppString,
				HelpText: "This command retrieves information a v3 app.",
				UsageDetails: plugin.Usage{
					Usage: fmt.Sprintf("cf %s app-name", getAppString),
				},
			},
			{
				Name:     listAppString,
				HelpText: "This command retrieves information v3 app in targeted space.",
				UsageDetails: plugin.Usage{
					Usage: fmt.Sprintf("cf %s app-name", listAppString),
				},
			},
			{
				Name:     uploadProcfileString,
				HelpText: "This command posts a procfile to create processes associated to the app.",
				UsageDetails: plugin.Usage{
					Usage: fmt.Sprintf("cf %s app-name path", uploadProcfileString),
				},
			},
			{
				Name:     removeProcessString,
				HelpText: "This command removes a process from an app.",
				UsageDetails: plugin.Usage{
					Usage: fmt.Sprintf("cf %s app-name process-type", removeProcessString),
				},
			},
			{
				Name:     listProcessesString,
				HelpText: "This command lists processes from an app.",
				UsageDetails: plugin.Usage{
					Usage: fmt.Sprintf("cf %s app-name", listProcessesString),
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
	} else if args[0] == getAppString && len(args) == 2 {
		appName := args[1]
		c.getApp(cliConnection, appName)
	} else if args[0] == listAppString && len(args) == 1 {
		c.listApps(cliConnection)
	} else if args[0] == uploadProcfileString && len(args) == 3 {
		appName := args[1]
		procfilePath := args[2]
		c.uploadProcfile(cliConnection, appName, procfilePath)
	} else if args[0] == removeProcessString && len(args) == 3 {
		appName := args[1]
		processType := args[2]
		c.removeProcess(cliConnection, appName, processType)
	} else if args[0] == listProcessesString && len(args) == 2 {
		appName := args[1]
		c.listProcesses(cliConnection, appName)
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
	createCommand := apps.NewCreateAppCommand(appName, cliConnection)
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

	deleteCommand := apps.NewDeleteAppCommand(appName, cliConnection)
	deleteCommand.Perform()
}

func (c *V3Cli) getApp(cliConnection plugin.CliConnection, appName string) {
	getCommand := apps.NewGetAppCommand(appName, cliConnection)
	getCommand.Perform()
}

func (c *V3Cli) listApps(cliConnection plugin.CliConnection) {
	listCommand := apps.NewListAppsCommand(cliConnection)
	listCommand.Perform()
}

func (c *V3Cli) uploadProcfile(cliConnection plugin.CliConnection, appName, procfilePath string) {
	uploadCommand := apps.NewUploadProcfileCommand(appName, procfilePath, cliConnection)
	uploadCommand.Perform()
}

func (c *V3Cli) removeProcess(cliConnection plugin.CliConnection, appName, processType string) {
	removeCommand := apps.NewRemoveProcessCommand(appName, processType, cliConnection)
	removeCommand.Perform()
}

func (c *V3Cli) listProcesses(cliConnection plugin.CliConnection, appName string) {
	listCommand := apps.NewListAppProcessesCommand(appName, cliConnection)
	listCommand.Perform()
}
