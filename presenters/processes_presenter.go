package presenters

import (
	"fmt"

	"github.com/jfmyers9/v3_cli_plugin/resources"
)

type ProcessesPresenter struct {
	processes []resources.Process
}

func NewProcessesPresenter(processes []resources.Process) ProcessesPresenter {
	return ProcessesPresenter{processes: processes}
}

func (p *ProcessesPresenter) Present() string {
	var listOutput string
	for _, process := range p.processes {
		template := "Process Type: %s\tGuid: %s\tCommand: %s\n"
		listOutput += fmt.Sprintf(template, process.Type, process.Guid, process.Command)
	}
	return listOutput
}
