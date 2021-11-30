# v0.3.1

Documentation updates around GUI installation and fix MacOS app release upload script.

# v0.3.0

Gui feature added. We now have a gui application build with fyne. You can check out the installations in the releases tab.
- Added alias shortcuts to the gui and cli

# v0.2.1

- Updated location of the cli for easy installation with `go install github.com/jordangarrison/whats-my-status/wms@latest`
- Updated `Taskfile.yml` to use the new directory for building the application

# v0.2.0

Added the ability to set the status of GitHub users. You need the following for this to work correctly:
- Your GitHub username
- An access token with `user` scope
- An entry in the config

Also updated the documentation to reflect the new GitHub features.
## Configuration update

```yaml
workspaces:
  - name: [your-gihtub-username]
    type: github
    token: [ghp_your-generated-user-scope-access-token]
```

 # v0.1.0

 This is the first release of the project. We have basic functionality in the status command.

 ```sh
 wms status --emoji [:emoji:] --time [hrtime] Message goes here
 ```

 ## Aliases

 You can now configure aliases for the status command in the ~/.wms.yaml file.

 ```yaml
aliases:
  - name: my-alias
    status:
      emoji: :emoji:
      time: [hrtime]
      message: Message goes here
```