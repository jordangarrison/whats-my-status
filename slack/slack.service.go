package slack

import (
	util "github.com/jordangarrison/whats-my-status/util"
	"github.com/slack-go/slack"
)

type SlackClient struct {
	client    *slack.Client
	status    util.Status
	workspace util.Workspace
}

func GetClient(status util.Status, workspace util.Workspace) *SlackClient {
	return &SlackClient{
		client:    slack.New(workspace.Token),
		status:    status,
		workspace: workspace,
	}
}

func (slackClient *SlackClient) SetStatus() error {
	err := slackClient.client.SetUserCustomStatus(slackClient.status.StatusMessage, slackClient.status.Emoji, slackClient.status.Epoch)
	if err != nil {
		return err
	}
	return nil
}
