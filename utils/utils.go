package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jfmyers9/v3_cli_plugin/resources"
)

type Utils struct {
	CliConnection plugin.CliConnection
}

func (u *Utils) GetTargetSpaceGuid() (string, error) {
	output, err := u.CliConnection.CliCommandWithoutTerminalOutput("target")
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

	spaceGuid, err := u.CliConnection.CliCommandWithoutTerminalOutput("space", spaceName, "--guid")
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(spaceGuid[0]), nil
}

func (u *Utils) GetAppScopedToSpace(appName string) (resources.V3App, error) {
	spaceGuid, err := u.GetTargetSpaceGuid()
	if err != nil {
		return resources.V3App{}, err
	}

	appPath := fmt.Sprintf("/v3/apps?names[]=%s&space_guids[]=%s", appName, spaceGuid)
	output, err := u.CliConnection.CliCommandWithoutTerminalOutput("curl", appPath)
	if err != nil {
		return resources.V3App{}, err
	}

	var response resources.V3AppResponse
	err = json.Unmarshal([]byte(output[0]), &response)
	if err != nil {
		return resources.V3App{}, err
	}

	if len(response.Resources) == 0 {
		return resources.V3App{}, errors.New("App Not Found")
	}
	return response.Resources[0], nil
}

func (u *Utils) ListAppsScopedToSpace() ([]resources.V3App, error) {
	spaceGuid, err := u.GetTargetSpaceGuid()
	if err != nil {
		return []resources.V3App{}, err
	}

	appPath := fmt.Sprintf("/v3/apps?space_guids[]=%s", spaceGuid)
	output, err := u.CliConnection.CliCommandWithoutTerminalOutput("curl", appPath)
	if err != nil {
		return []resources.V3App{}, err
	}

	var response resources.V3AppResponse
	err = json.Unmarshal([]byte(output[0]), &response)
	if err != nil {
		return []resources.V3App{}, err
	}

	return response.Resources, nil
}
