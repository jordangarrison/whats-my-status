package status

import (
	cfg "github.com/jordangarrison/whats-my-status/config"
	slack "github.com/jordangarrison/whats-my-status/slack"
)

// SetStatus sets the status of the service
func SetStatus(config cfg.Config) error {
	for _, workspace := range config.Workspaces {
		switch workspace.Type {
		case "slack":
			client := slack.GetClient(config.Status, workspace)
			err := client.SetStatus()
			if err != nil {
				return err
			}
		default:
			panic("Unknown service: " + workspace.Type)
		}
	}
	return nil
}
