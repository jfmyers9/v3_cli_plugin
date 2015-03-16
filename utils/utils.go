package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
)

type V3App struct {
	Guid string `json:"guid"`
}

type V3AppResponse struct {
	Resources []V3App `json:"resources"`
}

func GetTargetSpaceGuid(cliConnection plugin.CliConnection) (string, error) {
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

func GetAppGuid(cliConnection plugin.CliConnection, appName string) (string, error) {
	appPath := fmt.Sprintf("/v3/apps?names[]=%s", appName)
	output, err := cliConnection.CliCommandWithoutTerminalOutput("curl", appPath)
	if err != nil {
		return "", err
	}

	var response V3AppResponse
	json.Unmarshal([]byte(output[0]), &response)
	return response.Resources[0].Guid, nil
}
