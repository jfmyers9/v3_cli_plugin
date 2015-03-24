package presenters

import (
	"fmt"

	"github.com/jfmyers9/v3_cli_plugin/resources"
)

type AppsPresenter struct {
	apps []resources.V3App
}

func NewAppsPresenter(apps []resources.V3App) AppsPresenter {
	return AppsPresenter{apps: apps}
}

func (p *AppsPresenter) Present() string {
	var listOutput string
	for _, app := range p.apps {
		template := "App Name: %s\tGuid: %s\tState: %s\n"
		listOutput += fmt.Sprintf(template, app.Name, app.Guid, app.DesiredState)
	}
	return listOutput
}
