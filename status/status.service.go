package status

import (
	"errors"

	slack "github.com/jordangarrison/whats-my-status/slack"
	util "github.com/jordangarrison/whats-my-status/util"
)

// SetStatus sets the status of the service
func SetStatus(config util.Config) error {
	for _, workspace := range config.Workspaces {
		switch workspace.Type {
		case "slack":
			client := slack.GetClient(config.Status, workspace)
			err := client.SetStatus()
			if err != nil {
				return err
			}
		default:
			return errors.New("Unknown service: " + workspace.Type)
		}
	}
	return nil
}
