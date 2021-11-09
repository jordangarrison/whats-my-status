package utility

type Config struct {
	// Status
	Status Status `json:"status"`
	// Workspaces
	Workspaces []Workspace `json:"workspaces"`
	// Aliases
	Aliases []Alias `json:"aliases"`
}

type Workspace struct {
	// Name of the workspace, this is the actual name
	Name string `json:"name"`
	// Token used for authentication
	Token string `json:"token"`
	// Type of the workspace
	Type string `json:"type"`
}

type Alias struct {
	// Name of the alias
	Name   string `json:"name"`
	Status Status `json:"status"`
}

type Status struct {
	// Status to set
	StatusMessage string `json:"message"`
	// Emoji to set
	Emoji string `json:"emoji"`
	// Time to set
	Time string `json:"time"`
	// Epoch to set
	Epoch int64 `json:"epoch"`
	// ISO8601 to set
	ISO8601 string `json:"iso8601"`
}
