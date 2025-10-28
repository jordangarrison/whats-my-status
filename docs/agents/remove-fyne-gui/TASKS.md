# Tasks: Remove Fyne GUI

## Status Legend
- [ ] Todo
- [→] In Progress
- [x] Done

## Phase 1: Source Code Removal

- [x] Delete `main.go` (GUI entry point)
- [x] Delete `ui/about.go`
- [x] Delete `ui/alias.go`
- [x] Delete `ui/status.go`
- [x] Delete `ui/tabs.go`
- [x] Remove empty `ui/` directory
- [x] Delete `Icon.png`

## Phase 2: Build System Updates

- [x] Remove `build_gui` task from Taskfile.yml
- [x] Remove `build_gui_linux` task from Taskfile.yml
- [x] Remove `build_gui_freebsd` task from Taskfile.yml
- [x] Remove `build_gui_android` task from Taskfile.yml
- [x] Remove `build_gui_ios` task from Taskfile.yml
- [x] Remove `build_gui_windows` task from Taskfile.yml
- [x] Remove `build_gui_darwin` task from Taskfile.yml
- [x] Remove `build_gui_darwin_amd64` task from Taskfile.yml
- [x] Remove `build_gui_darwin_arm64` task from Taskfile.yml
- [x] Remove `zip_gui_darwin_amd64` task from Taskfile.yml
- [x] Remove `zip_gui_darwin_arm64` task from Taskfile.yml
- [x] Remove `clean_gui` task from Taskfile.yml
- [x] Remove `clean_build_gui` task from Taskfile.yml
- [x] Update `clean` task to only call `clean_cli`
- [x] Update `clean_build` task to only build CLI
- [x] Update `release` task to remove GUI binary references (lines 179-181)

## Phase 3: Dependency Cleanup

- [x] Edit go.mod to remove `fyne.io/fyne/v2` from requires
- [x] Run `go mod tidy` to clean up indirect dependencies
- [x] Verify no Fyne imports remain: `grep -r "fyne.io" --include="*.go"`
- [x] Run `task test` to verify all tests pass
- [x] Run `task build` to verify CLI builds successfully

## Phase 4: Documentation Updates

### README.md
- [x] Remove GUI section header (line 7)
- [x] Remove GUI installation intro (lines 8-9)
- [x] Remove MacOS GUI instructions (lines 11-13)
- [x] Remove Windows GUI instructions (lines 15-17)
- [x] Remove Linux GUI instructions (lines 19-27)
- [x] Update project description to mention CLI-only (line 3)

### CLAUDE.md (AGENTS.md)
- [x] Update "Architecture" section introduction
- [x] Remove "Dual Entry Points" subsection (lines ~14-16)
- [x] Update to show single CLI entry point: `wms/main.go`
- [x] Remove `ui/` from "Core Packages" list (line ~22)
- [x] Update development commands to remove GUI build references

## Phase 5: Artifact Cleanup

- [x] Remove `fyne-cross/` directory if it exists: `rm -rf fyne-cross/`
- [x] Verify `build/` directory only contains CLI binaries

## Verification

- [x] No GUI source files exist: `find . -name "main.go" -not -path "./wms/*"`
- [x] No Fyne in go.mod: `grep fyne go.mod`
- [x] CLI builds: `task build`
- [x] Tests pass: `task test`
- [x] No Fyne imports: `grep -r "fyne" --include="*.go"`
- [x] Manually test CLI: `./build/wms-linux-amd64-* status --help`

## Dependencies

- Phase 2 can be done in parallel with Phase 1
- Phase 3 must wait for Phase 1 to complete (all source deletions)
- Phase 4 can be done in parallel with Phase 3
- Phase 5 can be done at any time
- Verification must be done after all other phases complete
