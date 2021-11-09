package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	util "github.com/jordangarrison/whats-my-status/util"
	"golang.org/x/oauth2"
)

type GitHubClient struct {
	client    *http.Client
	status    util.Status
	workspace util.Workspace
}

type GraphQLRequest struct {
	Query     string `json:"query"`
	Variables string `json:"variables"`
}

func GetClient(status util.Status, workspace util.Workspace) *GitHubClient {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: workspace.Token},
	)
	client := oauth2.NewClient(ctx, ts)
	return &GitHubClient{
		client:    client,
		status:    status,
		workspace: workspace,
	}
}

func (gitHubClient *GitHubClient) SetStatus() error {
	apiUrl := "https://api.github.com/graphql"
	mutation := `
			mutation ($status: ChangeUserStatusInput!) {
				changeUserStatus(input: $status) {
					status {
						emoji
						expiresAt
						limitedAvailability: indicatesLimitedAvailability
						message
					}
				}
			}
		`
	variables := fmt.Sprintf(`
		{
			"status": {
				"emoji": "%s",
				"message": "%s",
				"expiresAt": "%s"
			}
		}
	`, gitHubClient.status.Emoji, gitHubClient.status.StatusMessage, gitHubClient.status.ISO8601)

	gqlMarshalled, err := json.Marshal(GraphQLRequest{
		Query:     mutation,
		Variables: variables,
	})
	if err != nil {
		return err
	}

	resp, err := gitHubClient.client.Post(apiUrl, "application/json",
		strings.NewReader(string(gqlMarshalled)))
	if err != nil {
		return err
	}
	_, err = httputil.DumpResponse(resp, true)
	if err != nil {
		return err
	}
	return nil
}
