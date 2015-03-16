package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jfmyers9/v3_cli_plugin/commands"
)

type V3Cli struct{}

func (c *V3Cli) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "V3Cli",
		Commands: []plugin.Command{
			{
				Name:     "create-v3-app",
				HelpText: "This command creates a v3 app.",
				UsageDetails: plugin.Usage{
					Usage: "cf create-app app-name",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(V3Cli))
}

func (c *V3Cli) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "create-v3-app" {
		appName := args[1]
		spaceGuid, err := getTargetSpaceGuid(cliConnection)
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
}

func getTargetSpaceGuid(cliConnection plugin.CliConnection) (string, error) {
	output, err := cliConnection.CliCommandWithoutTerminalOutput("target")
	if err != nil {
		return "", err
	}

	if len(output) < 5 {
		return "", errors.New("Space not targeted.")
	}

	if !strings.HasPrefix(output[4], "Space:") || strings.Contains(output[4], "No space targeted") {
		return "", errors.New("Space not targeted.")
	}

	spaceName := strings.TrimSpace(strings.TrimPrefix(output[4], "Space:"))

	spaceGuid, err := cliConnection.CliCommandWithoutTerminalOutput("space", spaceName, "--guid")
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(spaceGuid[0]), nil
}
