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
	cliConnection plugin.CliConnection
}

func NewUtils(cliConnection plugin.CliConnection) Utils {
	return Utils{cliConnection: cliConnection}
}

func (u *Utils) GetTargetSpaceGuid() (string, error) {
	output, err := u.cliConnection.CliCommandWithoutTerminalOutput("target")
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

	spaceGuid, err := u.cliConnection.CliCommandWithoutTerminalOutput("space", spaceName, "--guid")
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
	output, err := u.cliConnection.CliCommandWithoutTerminalOutput("curl", appPath)
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
	output, err := u.cliConnection.CliCommandWithoutTerminalOutput("curl", appPath)
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

func (u *Utils) ListProcessesScopedToApp(appGuid string) ([]resources.Process, error) {
	processesPath := fmt.Sprintf("/v3/apps/%s/processes", appGuid)
	output, err := u.cliConnection.CliCommandWithoutTerminalOutput("curl", processesPath)
	if err != nil {
		return []resources.Process{}, err
	}

	var response resources.ProcessesResponse
	err = json.Unmarshal([]byte(output[0]), &response)
	if err != nil {
		return []resources.Process{}, err
	}

	return response.Resources, nil
}

func (u *Utils) GetProcessScopedToApp(appGuid string, processType string) (resources.Process, error) {
	processes, err := u.ListProcessesScopedToApp(appGuid)
	if err != nil {
		return resources.Process{}, err
	}

	for _, process := range processes {
		if process.Type == processType {
			return process, nil
		}
	}

	return resources.Process{}, errors.New("Process Not Found")
}
