# What's My Status?

What's my status? is a command line utility that allows you to set your status on multiple platforms at once.

## Installation

To install, download the binary for your platform from the [latest release](https://github.com/jordangarrison/whats-my-status/releases) and place it in your `$PATH`.

You can of course install from the go command, however, currently this creates the binary as `whats-my-status` instead of `wms`.

```sh
go install github.com/jordangarrison/whats-my-status@latest
```

## Setup

Copy the `.wms.example.yaml` file to `~/.wms.yaml` and edit it to your liking.

The general stucture of the file is as follows:

```yaml
# workspaces are the places you want to set your status
workspaces:
  - name: <your-slack-org>
    type: slack
    token: <your-slack-token>
  - name: <your-github-user>
    type: github
    token: <your-github-user-scope-token>
aliases:
  - name: myalias
    status:
      StatusMessage: "I'm doing something"
      Emoji: :smile:
      Time: "30m"
```

## Aliases

You can set up aliases for your status commands. For example, if you want to set your status to `"Focus time"`, you can do the following:

```yaml
aliases:
  - name: focus
    status:
      StatusMessage: "Focus Time"
      Emoji: ":compute:"
      Time: "1h"
```

You can run this alias simply with the following command:

```sh
wms status focus
```

The `clear` alias is a preset alias which will clear your status on all your workspaces.

## Tokens

You will need to generate tokens for access to your workspaces.

### Slack

For now, slack uses an old style token. You can generate it following these steps from the [Emacs Slack Repo (yuya373/emacs-slack)](https://github.com/yuya373/emacs-slack#how-to-get-token):

- Navigate to your Slack workspace customization portal at `https://[your-workspace].slack.com/customize`
- Log in with your credentials for the workspace if needed
- Open the console (`Ctrl+Shift+J` on Linux/Windows, `Cmd+Opt+J` on Mac)
- Run the following command: `window.prompt('your api token is: ', TS.boot_data.api_token)`
- Copy token to your `~/.wms.yaml` file

### GitHub

For github navigate to your [profile tokens settings](https://github.com/settings/tokens) and generate a new token with the following scopes:
- `user`

## Platform Support

- [x] Slack
- [x] GitHub
- [ ] Discord
- [ ] Google Hangouts
- [ ] Microsoft Teams
- [ ] Nextcloud
- [ ] Matrix (via Riot)