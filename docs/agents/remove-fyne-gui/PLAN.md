# Remove Fyne GUI Implementation

## Problem Statement

The Fyne-based GUI implementation (`main.go` and `ui/` package) was never fully functional and represents unnecessary bloat in the repository. The CLI application (`wms/`) is the primary and working interface for the application. The GUI components should be completely removed to:

- Reduce maintenance burden
- Remove unused dependencies
- Simplify the codebase
- Focus on the working CLI interface

## Current State Analysis

### GUI Components Identified

**Source Files:**
- `main.go` - GUI entry point using Fyne app framework
- `ui/about.go` - About tab UI component
- `ui/alias.go` - Alias management UI component
- `ui/status.go` - Status setting UI component
- `ui/tabs.go` - Tab container setup

**Assets:**
- `Icon.png` - Application icon for GUI

**Dependencies (go.mod):**
- `fyne.io/fyne/v2 v2.1.1` - Primary GUI framework
- Multiple indirect Fyne dependencies (GL, graphics, etc.)

**Build System (Taskfile.yml):**
- `build_gui` - Builds all GUI targets
- `build_gui_linux`, `build_gui_freebsd`, `build_gui_darwin`, `build_gui_windows`, `build_gui_android`, `build_gui_ios` - Platform-specific GUI builds
- `build_gui_darwin_amd64`, `build_gui_darwin_arm64` - macOS architecture variants
- `zip_gui_darwin_*` - macOS app packaging
- `clean_gui` - Clean GUI build artifacts
- `clean_build_gui` - Clean and rebuild GUI
- GUI binary references in `release` task

**Documentation:**
- README.md lines 7-33 - GUI installation instructions
- CLAUDE.md (AGENTS.md) - References to dual-interface architecture

**Build Artifacts:**
- `fyne-cross/` directory - Cross-compilation outputs
- `*.syso` files - Windows resource files

### CLI Components (Unchanged)

The working CLI implementation in `wms/` directory is completely independent:
- `wms/main.go` - CLI entry point
- `wms/cmd/root.go` - Cobra root command
- `wms/cmd/status.go` - Status command implementation

Shared business logic packages remain unaffected:
- `status/` - Status service orchestration
- `slack/` - Slack API integration
- `github/` - GitHub API integration
- `util/` - Configuration and utilities

## Proposed Approach

### 1. Delete GUI Source Code
Remove all GUI-specific source files and assets that are no longer needed.

### 2. Update Build System
Clean up Taskfile.yml to remove all GUI-related build tasks and update the release process to only build/package CLI binaries.

### 3. Update Dependencies
Remove Fyne from go.mod and clean up all indirect dependencies that were only needed for the GUI.

### 4. Update Documentation
Update README.md and CLAUDE.md to reflect the CLI-only nature of the project.

### 5. Clean Build Artifacts
Remove any GUI build output directories.

## Implementation Steps

### Phase 1: Source Code Removal
1. Delete `main.go` (GUI entry point)
2. Delete entire `ui/` directory and all files within it
3. Delete `Icon.png` asset file

### Phase 2: Build System Updates
1. Remove all `build_gui*` tasks from Taskfile.yml
2. Remove `clean_gui` and `clean_build_gui` tasks
3. Update `clean_build` task to only depend on CLI build
4. Update `release` task to remove GUI binary references

### Phase 3: Dependency Cleanup
1. Remove `fyne.io/fyne/v2` from go.mod requires section
2. Run `go mod tidy` to automatically remove unused indirect dependencies
3. Verify the CLI still builds correctly

### Phase 4: Documentation Updates
1. README.md:
   - Remove "GUI" section (lines 7-27)
   - Keep only "CLI" installation section
   - Update project description to be CLI-focused
2. CLAUDE.md (AGENTS.md):
   - Update "Architecture" section to remove dual-interface mention
   - Remove "Dual Entry Points" subsection
   - Update to show single CLI entry point
   - Remove references to `ui/` package

### Phase 5: Artifact Cleanup
1. Remove `fyne-cross/` directory if it exists
2. Optionally update .gitignore to remove GUI-specific entries (or leave them as harmless)

## Acceptance Criteria

- [ ] No GUI source files remain in repository
- [ ] No Fyne dependencies in go.mod
- [ ] CLI builds successfully on all target platforms
- [ ] Taskfile.yml contains only CLI build tasks
- [ ] README.md documents only CLI usage
- [ ] CLAUDE.md reflects CLI-only architecture
- [ ] All tests pass
- [ ] `go mod tidy` shows no warnings

## Testing Plan

1. Run `task clean` to remove all build artifacts
2. Run `task mod` to tidy dependencies
3. Run `task test` to verify tests still pass
4. Run `task build` to verify CLI builds for all platforms
5. Manually test built CLI binary with sample config

## Risk Assessment

**Low Risk:**
- GUI and CLI are completely separate entry points
- No shared code dependencies from GUI to CLI
- Business logic packages (`status/`, `slack/`, `github/`, `util/`) are used by CLI and remain untouched
- Changes are purely deletions, not modifications

**Validation:**
- CLI implementation will not be affected as it has its own entry point in `wms/main.go`
- All platform-specific services remain unchanged
- Configuration system (Viper) remains unchanged
