package resources

type V3App struct {
	Guid         string `json:"guid"`
	Name         string `json:"name"`
	DesiredState string `json:"desired_state"`
}

type V3AppResponse struct {
	Resources []V3App `json:"resources"`
}
