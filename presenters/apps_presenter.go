package presenters

import (
	"fmt"

	"github.com/jfmyers9/v3_cli_plugin/resources"
)

type AppsPresenter struct {
	Apps []resources.V3App
}

func (p *AppsPresenter) PresentApps() string {
	var listOutput string
	for _, app := range p.Apps {
		template := "App Name: %s\tGuid: %s\tState: %s\n"
		listOutput += fmt.Sprintf(template, app.Name, app.Guid, app.DesiredState)
	}
	return listOutput
}
