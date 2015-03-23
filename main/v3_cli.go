package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jfmyers9/v3_cli_plugin/commands/apps"
	"github.com/jfmyers9/v3_cli_plugin/commands/processes"
)

const (
	createAppString      = "create-v3-app"
	deleteAppString      = "delete-v3-app"
	getAppString         = "v3-app"
	listAppString        = "v3-apps"
	uploadProcfileString = "procfile"
	createProcessString  = "create-process"
	removeProcessString  = "remove-process"
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
				Name:     createProcessString,
				HelpText: "This command posts a procfile to create processes associated to the app.",
				UsageDetails: plugin.Usage{
					Usage: fmt.Sprintf("cf %s process-type", createProcessString),
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
	} else if args[0] == createProcessString && len(args) == 2 {
		processType := args[1]
		c.createProcess(cliConnection, processType)
	} else if args[0] == removeProcessString && len(args) == 3 {
		appName := args[1]
		processType := args[2]
		c.removeProcess(cliConnection, appName, processType)
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
	createCommand := apps.CreateAppCommand{
		AppName:       appName,
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

	deleteCommand := apps.DeleteAppCommand{
		AppName:       appName,
		CliConnection: cliConnection,
	}
	deleteCommand.Perform()
}

func (c *V3Cli) getApp(cliConnection plugin.CliConnection, appName string) {
	getCommand := apps.GetAppCommand{
		AppName:       appName,
		CliConnection: cliConnection,
	}
	getCommand.Perform()
}

func (c *V3Cli) listApps(cliConnection plugin.CliConnection) {
	listCommand := apps.ListAppsCommand{
		CliConnection: cliConnection,
	}
	listCommand.Perform()
}

func (c *V3Cli) uploadProcfile(cliConnection plugin.CliConnection, appName string, procfilePath string) {
	uploadCommand := apps.UploadProcfileCommand{
		AppName:       appName,
		ProcfilePath:  procfilePath,
		CliConnection: cliConnection,
	}
	uploadCommand.Perform()
}

func (c *V3Cli) createProcess(cliConnection plugin.CliConnection, processType string) {
	createCommand := processes.CreateProcessCommand{
		ProcessType:   processType,
		CliConnection: cliConnection,
	}
	createCommand.Perform()
}

func (c *V3Cli) removeProcess(cliConnection plugin.CliConnection, appName string, processType string) {
	removeCommand := apps.RemoveProcessCommand{
		AppName:       appName,
		ProcessType:   processType,
		CliConnection: cliConnection,
	}
	removeCommand.Perform()
}
