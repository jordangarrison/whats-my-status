package slack

import (
	cfg "github.com/jordangarrison/whats-my-status/config"
	"github.com/slack-go/slack"
)

type SlackClient struct {
	client    *slack.Client
	status    cfg.Status
	workspace cfg.Workspace
}

func GetClient(status cfg.Status, workspace cfg.Workspace) SlackClient {
	slackClient := SlackClient{
		client:    nil,
		status:    status,
		workspace: workspace,
	}
	slackClient.client = slack.New(slackClient.workspace.Token)
	return slackClient
}

func (slackClient *SlackClient) SetStatus() error {
	err := slackClient.client.SetUserCustomStatus(slackClient.status.StatusMessage, slackClient.status.Emoji, 1234567890)
	if err != nil {
		return err
	}
	return nil
}
