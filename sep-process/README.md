# SEP (Software Enhancement Proposals)

A lightweight, AI-friendly system for organizing feature development through numbered documents that track work from idea to implementation.

## Philosophy

**Trust the AI to know how to build software.** SEPs specify *what* to build and *when it's done*, not *how* to build it. The AI figures out the implementation details.

## Quick Start

### Initialize in a New Project

```bash
# Start fresh
/sep-init

# Or import an existing PRD as SEP-0001
/sep-init docs/product-requirements.md
```

### Create Your First Feature

```bash
# Create a new SEP
/sep-new "User Authentication"

# Discuss and refine requirements
/sep-discuss 0001

# When ready, implement it
/sep-implement 0001

# Check progress anytime
/sep-status
```

## Available Commands

### `/sep-init` - Bootstrap SEP System

Initialize the SEP system in your project, optionally importing an existing PRD.

**Usage:**
```
/sep-init
/sep-init path/to/prd.md
```

**What it does:**
- Creates `docs/seps/` directory structure
- Copies process documentation (`0000-sep-process.md`)
- Sets up SEP template and README
- Installs SEP commands to `.claude/commands/`
- Optionally converts existing PRD to SEP-0001

**Example:**
```
/sep-init docs/mvp-requirements.md
```

---

### `/sep-new` - Create New SEP

Create a new SEP from template with automatic numbering.

**Usage:**
```
/sep-new "Feature Name"
```

**What it does:**
- Finds next available SEP number
- Creates file: `docs/seps/XXXX-feature-name.md`
- Fills in metadata (number, title, date)
- Sets status to DRAFT
- Updates SEP index if it exists

**Example:**
```
/sep-new "Email Notifications"
→ Creates SEP-0007: Email Notifications
→ File: docs/seps/0007-email-notifications.md
```

---

### `/sep-suggest` - AI Feature Suggestions

Get AI-generated SEP suggestions based on your current project state.

**Usage:**
```
/sep-suggest
```

**What it does:**
- Analyzes completed and active SEPs
- Reviews codebase and architecture
- Identifies gaps and natural next steps
- Generates 3-5 focused feature suggestions
- Estimates complexity and dependencies
- Saves suggestions to `docs/seps/SUGGESTIONS.md`

**Interactive workflow:**
- Shows suggested features with rationale
- You pick one or suggest your own
- Automatically creates new SEP for selected suggestion

---

### `/sep-discuss` - Refine Requirements

Collaboratively review and refine a DRAFT SEP's requirements.

**Usage:**
```
/sep-discuss XXXX
```

**What it does:**
- Shows current SEP content (What & Why, Done When)
- Analyzes scope and identifies issues
- Warns if scope is too large
- Identifies missing acceptance criteria
- Updates SEP based on discussion
- Suggests `/sep-split` if needed

**Example:**
```
/sep-discuss 0003
→ Shows analysis and asks what to discuss
→ Iteratively refines requirements
→ Updates "Done When" criteria
```

---

### `/sep-split` - Break Up Large SEPs

Split an unfocused or oversized DRAFT SEP into multiple focused SEPs.

**Usage:**
```
/sep-split XXXX
```

**What it does:**
- Analyzes SEP scope
- Proposes 2-3 focused splits
- Shows suggested dependency order
- Creates new SEP files from split
- Marks original as superseded
- Updates index

**Example:**
```
/sep-split 0005
→ Proposes split into 3 focused SEPs
→ Creates SEP-0008, SEP-0009, SEP-0010
→ Marks SEP-0005 as DONE (split)
```

---

### `/sep-implement` - Build the Feature

Execute an ACCEPTED SEP step-by-step.

**Usage:**
```
/sep-implement XXXX
```

**What it does:**
- Reads SEP requirements (What & Why, Done When)
- Implements feature using AI judgment
- Checks off "Done When" items as completed
- Runs tests, type checks, linting
- Updates status: DRAFT → DONE
- Commits with `SEP-XXXX:` prefix
- Adds implementation notes to SEP

**Example:**
```
/sep-implement 0004
→ Builds the feature
→ Marks all "Done When" items complete
→ Updates SEP status to DONE
```

**Note:** If blocked or unclear, stops to ask for clarification instead of guessing.

---

### `/sep-status` - View All SEPs

Show current state of all SEPs and recommend next action.

**Usage:**
```
/sep-status
```

**What it shows:**
- All SEPs grouped by status
- Creation dates
- Dependencies
- Blocked reasons
- Recommended next action

**Example output:**
```
SEP Status
==========

DRAFT:
- SEP-0003: User Profiles (created 2025-11-18)
- SEP-0002: API Auth (created 2025-11-15) [depends on: SEP-0001]

BLOCKED:
- SEP-0004: Payments (created 2025-11-16) - waiting for API keys

DONE:
- SEP-0001: Database Setup (completed 2025-11-10)
- SEP-0000: SEP Process (completed 2025-11-18)

NEXT: Implement SEP-0003 (no dependencies)
```

---

## SEP States

| State | Meaning |
|-------|---------|
| **DRAFT** | Requirements defined, ready to discuss or implement |
| **BLOCKED** | Cannot proceed (add reason in SEP file) |
| **CANCELLED** | Decided not to build |
| **DONE** | Built, tested, and shipped |

## SEP File Structure

```
docs/
  seps/
    0000-sep-process.md          # Process documentation
    0001-user-authentication.md  # Feature SEP
    0002-api-endpoints.md        # Another feature
    SEP-TEMPLATE.md              # Template for new SEPs
    README.md                    # SEP index (optional)
    SUGGESTIONS.md               # AI suggestions (generated)
```

## SEP Template Structure

```markdown
# SEP-XXXX: [Title]

**Status**: DRAFT | BLOCKED | CANCELLED | DONE
**Created**: YYYY-MM-DD
**Depends on**: SEP-XXXX (optional)

## What & Why

[Explain the feature and why you want it. 1-3 paragraphs.]

## Done When

- [ ] Acceptance criteria 1
- [ ] Acceptance criteria 2
- [ ] Acceptance criteria 3

---

## Implementation Notes

[Add notes as you build - what worked, gotchas, etc.]
```

## Commit Convention

**Always reference the SEP number in commits:**

```
SEP-0004: Add notification model
SEP-0004: Implement email service
SEP-0004: Add notification tests
SEP-0004: Update API documentation
```

This creates a clear history trail from requirements to implementation.

## Common Workflows

### Starting a New Feature

```bash
/sep-new "Real-time Chat"        # Create SEP-0012
# Fill in What & Why and Done When sections
/sep-discuss 0012                # Refine requirements
/sep-implement 0012              # Build it
```

### Exploring What to Build Next

```bash
/sep-status                      # See what's done and active
/sep-suggest                     # Get AI feature ideas
# Pick suggestion #2
# AI creates new SEP automatically
/sep-discuss 0013                # Refine the new SEP
```

### Handling Large Features

```bash
/sep-new "Complete Admin Dashboard"
/sep-discuss 0015                # AI warns scope is too large
/sep-split 0015                  # Split into 3 focused SEPs
/sep-implement 0016              # Build first piece
/sep-implement 0017              # Build second piece
/sep-implement 0018              # Build third piece
```

### Managing Blocked Work

```bash
/sep-status                      # See SEP-0010 is BLOCKED
# Edit SEP-0010: Add reason to BLOCKED status
# Status: BLOCKED - waiting for API access from vendor
# When unblocked:
/sep-implement 0010              # Resume implementation
```

### Import Existing Requirements

```bash
# You have docs/product-spec.md with requirements
/sep-init docs/product-spec.md   # Converts to SEP-0001
/sep-discuss 0001                # Refine into proper SEP
/sep-split 0001                  # Break into focused pieces
# Now have SEP-0002, 0003, 0004 ready to implement
```

## Best Practices

### Keep SEPs Focused
- ✅ One clear feature per SEP
- ✅ Can be completed in a reasonable timeframe
- ❌ Don't combine multiple unrelated features
- Use `/sep-split` if scope grows too large

### Write Clear Acceptance Criteria
- ✅ Specific, testable "Done When" items
- ✅ Focus on outcomes, not implementation
- Example: "User can reset password via email" not "Add reset_password() function"

### Use Dependencies
- Link SEPs that must be built in order
- Example: "SEP-0005: User Profiles" depends on "SEP-0003: Authentication"
- `/sep-status` shows the build order

### Iterate on Requirements
- Use `/sep-discuss` liberally while in DRAFT
- Better to refine requirements upfront
- Implementation goes smoother with clear specs

### Trust the AI
- Don't specify *how* to build, specify *what* to build
- Let `/sep-implement` make technical decisions
- Focus SEP content on business requirements

## Rules

1. **Sequential numbering**: 0001, 0002, 0003... (zero-padded to 4 digits)
2. **Only implement DRAFT**: Don't implement BLOCKED or CANCELLED SEPs
3. **One feature per SEP**: Keep scope focused, split if needed
4. **Always reference SEP in commits**: Use `SEP-XXXX:` prefix
5. **Use dependencies**: Track build order explicitly
6. **Update status**: Keep SEP status current (DRAFT → DONE, etc.)

## Why SEPs?

### For Solo Developers
- Clear record of what you wanted to build
- Easy to resume work after breaks
- Git history tied to requirements
- Prevents scope creep

### For AI Collaboration
- AI understands what to build without micro-management
- Clear acceptance criteria for validation
- Structured format AI can parse and work with
- History of decisions and implementation

### For Teams
- Shared understanding of features
- Clear ownership and status tracking
- Dependencies visible upfront
- Documentation generated as you build

## Tips

- Run `/sep-status` regularly to see progress
- Use `/sep-suggest` when out of ideas
- Don't hesitate to `/sep-split` large features
- Keep "Done When" criteria specific and testable
- Mark SEPs as BLOCKED with clear reasons
- Commit early and often with SEP references

## Installation

These commands are Claude Code custom commands. The `/sep-init` command will automatically set them up in your project's `.claude/commands/` directory, or you can manually copy them there.

## Learn More

Read `docs/seps/0000-sep-process.md` for the complete process documentation and philosophy.