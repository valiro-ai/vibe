# Vibe

CLI tool for managing Software Enhancement Proposals (SEPs) in AI-native development workflows.

## Quick Start

```bash
# Build
go build -o vibe ./cmd/vibe/

# Initialize in your repo
./vibe init

# Create a SEP
./vibe sep new "User Authentication"

# Check status
./vibe sep status
```

## Workflow

```
Author          Editor            Pilot
  │               │                 │
  ├─ Create SEP   │                 │
  ├─ Draft        │                 │
  │               ├─ Review         │
  │               ├─ Approve        │
  │               │   (ACCEPTED)    │
  │               │                 ├─ Claim
  │               │                 ├─ /sep-plan
  │               │                 ├─ /sep-implement
  │               │                 └─ Done
```

1. **Author**: `vibe sep new "Feature"` → Edit SEP (what, why, done-when, areas)
2. **Editor**: Review → `vibe sep update XXXX ACCEPTED`
3. **Pilot**: `vibe sep claim XXXX @name` → `/sep-plan` → `/sep-implement`

## Roles

| Role | Responsibility |
|------|----------------|
| **Author** | Draft SEPs - anyone on the team |
| **Editor** | Review, give feedback, approve SEPs |
| **Pilot** | Implement approved SEPs with AI agents |

## Commands

### CLI

```bash
vibe init                    # Set up SEP workflow
vibe sep new "Feature"       # Create new SEP
vibe sep list                # List all SEPs
vibe sep status              # Show progress + next action
vibe sep sync                # Pull latest + show pipeline
vibe sep claim XXXX @pilot   # Claim SEP (assign + commit + push)
vibe sep pipeline            # Show areas + conflicts + assignments
vibe sep update XXXX STATUS  # Update status (DRAFT/ACCEPTED/DONE/BLOCKED)
vibe feedback "message"      # Submit feedback
```

### Claude Code

```
/sep-discuss XXXX     # Refine requirements
/sep-plan XXXX        # Create implementation plan
/sep-implement XXXX   # Execute the plan
/sep-split XXXX       # Break up large SEP
/sep-suggest          # AI suggests new SEPs
```

## Documentation

See [docs/](docs/) for full documentation:

- [Getting Started](docs/getting-started.md)
- [CLI Reference](docs/cli-reference.md)
- [Workflow Guide](docs/workflow.md)
