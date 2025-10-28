# AGENTS.md

This file provides guidance to AI coding assistants (Claude Code, Warp AI, Cursor, etc.) when working with code in this repository.

## Project Overview

What's My Status is a command line application for setting presence status across multiple platforms (Slack, GitHub) simultaneously. Built in Go with Cobra for the CLI.

## Organizing Your Work

For complex fixes or features, use the `docs/agents/` folder structure to organize your work:

- Create a descriptive subfolder (e.g., `docs/agents/fix-build-issue/`)
- Add `PLAN.md` to document the problem, approach, and implementation steps
- Add `TASKS.md` to track actionable tasks and their status
- See `docs/agents/README.md` for detailed guidelines and examples

This structure helps maintain context, track progress, and document decisions for both current work and future reference.

## Architecture

### Entry Point
- `wms/main.go`: CLI application entry (Cobra-based)

### Core Packages
- `status/`: Orchestration layer that dispatches to platform-specific services
- `slack/`: Slack API integration for status updates
- `github/`: GitHub API integration for status updates
- `util/`: Shared types (`Config`, `Workspace`, `Alias`, `Status`) and utilities

### Configuration System
Uses Viper for config management. Config file is `.wms.yaml` and is searched in:
1. Current directory (`.`)
2. User home directory (`~`)
3. `~/.config`

Config structure:
- `workspaces[]`: Platform accounts (slack/github) with tokens
- `aliases[]`: Predefined status shortcuts with message, emoji, time
- `status`: Current status being set

### Platform Services Pattern
Each platform (slack, github) implements a service with:
- `GetClient(status, workspace)`: Factory function returning configured client
- `SetStatus()`: Method to apply status to the platform

## Important: Use Taskfile for All Operations

**ALWAYS use the Taskfile commands for building, testing, and other operations.** Do not run `go build`, `go test`, or other go commands directly. The Taskfile ensures consistent builds across platforms and includes necessary steps like dependency management.

Examples:
- ✅ `task build` - Correct
- ❌ `go build ./...` - Don't do this
- ✅ `task test` - Correct
- ❌ `go test ./...` - Don't do this

## Development Commands

All development operations should be performed using the Taskfile. Run `task --list` to see all available commands.

### Building

**Always use `task build` for building**, not direct go commands.

Build for all platforms:
```bash
task build
```

This automatically:
- Runs `go mod tidy` to ensure dependencies are current
- Runs tests to verify code quality
- Builds binaries for all supported platforms (Linux amd64/arm/arm64, macOS amd64, Windows amd64)

Build for specific platform:
```bash
task build_linux_amd64
task build_darwin_amd64
task build_windows_amd64
task build_linux_arm64
task build_linux_arm
```

### Testing

**Always use `task test` for testing**, not `go test` directly.

```bash
task test
```

This runs:
- `go vet ./...` - Static analysis
- `go test -count=1 -v ./...` - All tests with verbose output

### Dependency Management

Update dependencies:
```bash
task mod
```

This runs `go mod tidy` to clean up and update the `go.mod` file.

### Installing Locally

Install to `~/.local/bin`:
```bash
task local_install
```

This performs a clean build and installs the appropriate binary for your platform.

### Cleaning

Clean build artifacts:
```bash
task clean
```

Removes the `build/` directory containing all built binaries.

### Clean and Build

Clean and rebuild everything:
```bash
task clean_build
```

This is equivalent to `task clean && task build`.

### Releasing

Create a draft GitHub release:
```bash
task release
```

This:
- Performs a clean build of all platform binaries
- Creates a draft GitHub release using `gh` CLI
- Attaches all binaries with descriptive labels
- Uses `changelog.md` for release notes

## Key Implementation Details

### Status Aliases
- Built-in alias: `clear` (clears status on all workspaces)
- Custom aliases defined in config with name, message, emoji, and duration
- Aliases are resolved in `wms/cmd/status.go` before dispatching to services

### Time Handling
- Accepts duration strings like "30m", "1h"
- Converts to both Unix epoch and ISO8601 format
- Helper functions in `util/utility.go`

### CLI Usage Pattern
```bash
wms status -e ":smile:" -t 30m "Working on something"
wms status <alias-name>
wms clear  # Built-in alias
```

### Configuration Example
See `.wms.example.yaml` for structure. Users copy to `~/.wms.yaml` with their tokens.

## Dependencies

- `github.com/spf13/cobra`: CLI framework
- `github.com/spf13/viper`: Configuration management
- `github.com/slack-go/slack`: Slack API client
- `golang.org/x/oauth2`: OAuth2 for GitHub
- `github.com/hako/durafmt`: Duration formatting

## Version Management

Version is defined in `Taskfile.yml` under `vars.VERSION`. Update this for releases.
