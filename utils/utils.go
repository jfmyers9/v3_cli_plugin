package utils

import (
	"errors"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
)

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
