# CLI Reference

Complete reference for all vibe commands.

## Global Flags

| Flag | Description |
|------|-------------|
| `-h, --help` | Show help for any command |

## Commands

### vibe init

Initialize vibe workflow in a repository.

```bash
vibe init
```

Creates:
- `docs/seps/` - SEP documents directory
- `docs/seps/0000-sep-process.md` - Process documentation
- `docs/seps/SEP-TEMPLATE.md` - Template for new SEPs
- `.claude/commands/` - Claude Code custom commands

### vibe sep

Parent command for SEP management.

```bash
vibe sep [command]
```

#### vibe sep new

Create a new SEP.

```bash
vibe sep new "Feature Title"
```

**Arguments:**
- `title` - The title of the SEP (required)

**Flags:**
- `-d, --dir` - SEP directory (default: `docs/seps`)

**Example:**
```bash
vibe sep new "User Authentication"
# Created: docs/seps/0001-user-authentication.md
```

#### vibe sep list

List all SEPs grouped by status.

```bash
vibe sep list
```

**Flags:**
- `-d, --dir` - SEP directory (default: `docs/seps`)

**Example output:**
```
DRAFT:
  SEP-0001  User Authentication      (created 2025-01-15)
  SEP-0002  User Profile Management  (created 2025-01-15)

BLOCKED:
  SEP-0003  Payment Integration  (created 2025-01-14)

DONE:
  SEP-0000  SEP Process  (created 2025-01-01)

Total: 4 SEPs
```

#### vibe sep status

Show SEP progress and recommended next action.

```bash
vibe sep status
```

**Flags:**
- `-d, --dir` - SEP directory (default: `docs/seps`)

**Example output:**
```
SEP Status
========================================

DRAFT:
  - SEP-0001: User Authentication (created 2025-01-15)

BLOCKED:
  - SEP-0003: Payment Integration
    → Waiting for: SEP-0001

DONE:
  - SEP-0000: SEP Process

NEXT: Implement SEP-0001 (no dependencies)
```

#### vibe sep update

Update the status of a SEP.

```bash
vibe sep update <number> <status>
```

**Arguments:**
- `number` - SEP number (e.g., `0001`)
- `status` - New status: `DRAFT`, `ACCEPTED`, `BLOCKED`, `DONE`, `CANCELLED`

**Flags:**
- `-d, --dir` - SEP directory (default: `docs/seps`)

**Example:**
```bash
vibe sep update 0001 DONE
# Updated SEP-0001 status to DONE
```

#### vibe sep claim

Claim a SEP (assign + commit + push in one step).

```bash
vibe sep claim <number> <pilot>
```

**Arguments:**
- `number` - SEP number (e.g., `0001`)
- `pilot` - Pilot name/handle (e.g., `@alice`)

**Examples:**
```bash
vibe sep claim 0001 @alice
# ✓ SEP-0001 claimed by @alice and pushed

vibe sep claim 0001 ""
# ✓ SEP-0001 unclaimed and pushed
```

This command:
1. Assigns the pilot to the SEP
2. Commits the change
3. Pushes to remote

Other pilots will see the claim after running `vibe sep sync`.

#### vibe sep sync

Pull latest changes and show pipeline.

```bash
vibe sep sync
```

**Example:**
```bash
vibe sep sync
# Pulling latest changes...
# Already up to date.
#
# SEP Pipeline - Area Conflicts
# ...
```

#### vibe sep assign

Assign a pilot locally (without committing/pushing). Use `claim` instead to share with team.

```bash
vibe sep assign <number> <pilot>
```

**Arguments:**
- `number` - SEP number (e.g., `0001`)
- `pilot` - Pilot name/handle (e.g., `@alice`)

#### vibe sep pipeline

Show SEP pipeline with area conflicts and assignments for pilot coordination.

```bash
vibe sep pipeline
```

**Flags:**
- `-d, --dir` - SEP directory (default: `docs/seps`)

**Example output:**
```
SEP Pipeline - Area Conflicts
==================================================

ACCEPTED:
  SEP-0001: User Authentication [@alice] ⚠️  CONFLICT with SEP-0002
    areas: internal/auth/*, internal/user/*, api/routes/login.go
  SEP-0002: User Profile Management ⚠️  CONFLICT with SEP-0001
    areas: internal/user/*, api/routes/profile.go

DONE: 1 SEPs completed

--------------------------------------------------
⚠️  Conflicts detected:
  SEP-0001 ↔ SEP-0002: internal/user/* (@alice assigned)

→ Coordinate with assigned pilots or implement sequentially
```

## Feedback

### vibe feedback

Submit feedback about vibe or a specific SEP.

```bash
vibe feedback "Your feedback message"
vibe feedback --sep 0001 "Feedback about this SEP"
vibe feedback  # Interactive multi-line mode
```

**Flags:**
- `--sep` - Link feedback to a specific SEP number
- `--file` - Feedback log file (default: `docs/feedback.log`)

### vibe feedback list

View all recorded feedback.

```bash
vibe feedback list
```

**Example output:**
```
Feedback Log
==================================================
[2025-01-15 10:30] The claim command is really useful
[2025-01-15 14:22] SEP-0001: Acceptance criteria could be clearer
[2025-01-16 09:15] Would be nice to have notifications
```

### vibe feedback clear

Clear all feedback after reviewing.

```bash
vibe feedback clear
```

---

## Claude Commands

These commands are available in Claude Code after running `vibe init`.

| Command | Description |
|---------|-------------|
| `/sep-status` | Show current SEP status and next action |
| `/sep-new` | Create a new SEP interactively |
| `/sep-discuss XXXX` | Discuss and refine a SEP's requirements |
| `/sep-plan XXXX` | Create implementation plan based on current codebase |
| `/sep-implement XXXX` | Execute the plan and build the feature |
| `/sep-split XXXX` | Split a large SEP into smaller ones |
| `/sep-suggest` | Get AI suggestions for new SEPs |

### /sep-plan

Create an implementation plan before building. Claude will:

1. Read the SEP document
2. Explore the current codebase state
3. Analyze files in `areas` and related code
4. Create a concrete implementation plan
5. Write the plan to the SEP's "## Plan" section
6. Update `areas` if additional files are discovered

```
/sep-plan 0001
```

Run this before `/sep-implement` to ensure the plan accounts for recent codebase changes.

### /sep-implement

Execute the plan and build the feature. Claude will:

1. Read the SEP document
2. Check for an existing plan (suggests `/sep-plan` if missing)
3. Follow the plan step-by-step
4. Check off completed acceptance criteria
5. Add implementation notes
6. Update status to DONE

```
/sep-implement 0001
```

### /sep-discuss

Refine a SEP before planning. Claude will:

1. Review the SEP
2. Ask clarifying questions
3. Suggest improvements to acceptance criteria
4. Help define areas that will be modified

```
/sep-discuss 0001
```
