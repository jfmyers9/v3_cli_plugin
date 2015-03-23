package resources

type V3App struct {
	Guid         string `json:"guid"`
	Name         string `json:"name"`
	DesiredState string `json:"desired_state"`
}

type V3AppResponse struct {
	Resources []V3App `json:"resources"`
}

type Process struct {
	Guid    string `json:"guid"`
	Type    string `json:"type"`
	Command string `json:"command"`
}

type ProcessesResponse struct {
	Resources []Process `json:"resources"`
}
