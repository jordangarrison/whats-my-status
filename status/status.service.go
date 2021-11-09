package status

import (
	"errors"
	"fmt"

	github "github.com/jordangarrison/whats-my-status/github"
	slack "github.com/jordangarrison/whats-my-status/slack"
	util "github.com/jordangarrison/whats-my-status/util"
)

// SetStatus sets the status of the service
func SetStatus(config util.Config) error {
	for _, workspace := range config.Workspaces {
		switch workspace.Type {
		case "slack":
			fmt.Println("Setting slack status for " + workspace.Name)
			client := slack.GetClient(config.Status, workspace)
			err := client.SetStatus()
			if err != nil {
				return err
			}
		case "github":
			fmt.Println("Setting github status for " + workspace.Name)
			client := github.GetClient(config.Status, workspace)
			client.SetStatus()
		default:
			return errors.New("Unknown service: " + workspace.Type)
		}
	}
	return nil
}
