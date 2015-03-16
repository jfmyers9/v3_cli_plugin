package presenters

import (
	"fmt"

	"github.com/jfmyers9/v3_cli_plugin/resources"
)

type AppPresenter struct {
	App resources.V3App
}

func (p *AppPresenter) PresentApp() string {
	appOutput := "App Name: %s\nApp Guid: %s\nApp State: %s"
	return fmt.Sprintf(appOutput, p.App.Name, p.App.Guid, p.App.DesiredState)
}
