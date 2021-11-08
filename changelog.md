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