package presenters

import (
	"fmt"

	"github.com/jfmyers9/v3_cli_plugin/resources"
)

type ProcessPresenter struct {
	process resources.Process
}

func NewProcessPresenter(process resources.Process) ProcessPresenter {
	return ProcessPresenter{process: process}
}

func (p *ProcessPresenter) Present() string {
	processOutput := "Process Type: %s\nProcess Guid: %s\nCommand: %s"
	return fmt.Sprintf(processOutput, p.process.Type, p.process.Guid, p.process.Command)
}
