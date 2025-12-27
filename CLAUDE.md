# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build Commands

```bash
task build            # Build the CLI
task run -- <command> # Run without building
task check            # Run fmt, vet, test
task install          # Install to $GOPATH/bin
```

Run `task --list` for all available tasks. No tests exist yet. No linter configured.

## Architecture

Vibe is a Go CLI tool for managing Software Enhancement Proposals (SEPs) in AI-native workflows. It coordinates work between three roles: Authors (draft SEPs), Editors (review/approve), and Pilots (implement with AI).

### Package Structure

- `main.go` - Entry point that executes the CLI.
- `internal/cli/` - Cobra CLI commands. Each `sep_*.go` file is a subcommand. `RootCmd` is exported for use by main.go.
- `internal/sep/` - Core SEP parsing and manipulation. The `SEP` struct represents a proposal with YAML frontmatter.
- `internal/templates/` - Embedded templates via Go's `embed` package. Contains SEP templates and Claude Code custom commands.

### Key Concepts

**SEP Lifecycle**: DRAFT → ACCEPTED → DONE (with BLOCKED, CANCELLED branches)

**SEP Files**: Markdown with YAML frontmatter in `docs/seps/`. Format: `XXXX-title-slug.md`

**Frontmatter fields**: `title`, `status`, `created`, `depends_on`, `areas`, `assigned`

**Templates**: Embedded at build time from `internal/templates/`. The `vibe init` command copies these to the target repo.

### Adding New Commands

1. Create `internal/cli/<command>.go`
2. Define a `cobra.Command`
3. Register in `init()` with parent command (e.g., `sepCmd.AddCommand(...)` or `RootCmd.AddCommand(...)`)

### SEP Parsing

The `internal/sep` package parses SEP files by:
1. Extracting YAML frontmatter between `---` markers
2. Parsing sections like "What & Why", "Done When", "Plan"
3. Tracking checkbox status for acceptance criteria

To modify a SEP file, use `SEP.UpdateStatus()` or `SEP.Assign()` which properly marshal YAML back to the file.
