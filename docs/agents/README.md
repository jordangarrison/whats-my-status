# Agent Workspace

This folder contains organized workspaces for AI coding assistants to work through fixes and features in a structured manner.

## Purpose

Each subfolder represents a specific fix, feature, or improvement being worked on. This organization helps maintain clarity and track progress on individual items.

## Structure

Each subfolder should follow this convention:

```
docs/agents/
├── <descriptive-folder-name>/
│   ├── PLAN.md     # Problem description, approach, and implementation plan
│   └── TASKS.md    # Actionable task list with status tracking
└── README.md       # This file
```

## Usage

### Creating a New Workspace

1. Create a descriptive subfolder (e.g., `fix-build-issue`, `add-discord-support`, `refactor-status-service`)
2. Add a `PLAN.md` file describing:
   - Problem statement
   - Proposed approach
   - Implementation steps
   - Acceptance criteria
3. Add a `TASKS.md` file with:
   - Actionable tasks
   - Status tracking (todo, in-progress, done)
   - Dependencies between tasks

### Example

```
docs/agents/fix-cli-config-loading/
├── PLAN.md     # Documents the config loading issue and fix approach
└── TASKS.md    # Lists specific tasks like "Update viper config paths", "Add tests"
```

## Benefits

- **Isolation**: Each fix/feature has its own workspace
- **Documentation**: Plans and decisions are recorded
- **Progress Tracking**: Task lists show what's done and what's remaining
- **Context**: Easy for AI assistants to understand current work
- **History**: Past work is preserved for reference
