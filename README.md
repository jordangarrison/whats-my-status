# What's My Status?

What's my status? is a command line utility that allows you to set your status on multiple platforms at once.

## Installation

```sh
go install github.com/jordangarrison/whatsmystatus@latest
```

## Setup

Copy the `.wms.example.yaml` file to `~/.wms.yaml` and edit it to your liking.

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

## Platform Support

- [x] Slack
- [ ] GitHub
- [ ] Discord
- [ ] Google Hangouts
- [ ] Microsoft Teams
- [ ] Nextcloud
- [ ] Matrix (via Riot)