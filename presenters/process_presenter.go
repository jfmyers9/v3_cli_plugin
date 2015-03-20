package presenters

import (
	"fmt"

	"github.com/jfmyers9/v3_cli_plugin/resources"
)

type ProcessPresenter struct {
	Process resources.Process
}

func (p *ProcessPresenter) Present() string {
	processOutput := "Process Type: %s\nProcess Guid: %s\nCommand: %s"
	return fmt.Sprintf(processOutput, p.Process.Type, p.Process.Guid, p.Process.Command)
}
