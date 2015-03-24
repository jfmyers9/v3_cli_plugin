package presenters

import (
	"fmt"

	"github.com/jfmyers9/v3_cli_plugin/resources"
)

type AppPresenter struct {
	app resources.V3App
}

func NewAppPresenter(app resources.V3App) AppPresenter {
	return AppPresenter{app: app}
}

func (p *AppPresenter) Present() string {
	appOutput := "App Name: %s\nApp Guid: %s\nApp State: %s"
	return fmt.Sprintf(appOutput, p.app.Name, p.app.Guid, p.app.DesiredState)
}
