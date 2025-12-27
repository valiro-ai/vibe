# Vibe - AI-Native Development Workflow

Vibe is a CLI tool for managing Software Enhancement Proposals (SEPs) in an AI-native development workflow. It coordinates work between authors, editors, and pilots.

## Quick Start

```bash
# Initialize vibe in your repository
./vibe init

# Create your first SEP
./vibe sep new "User Authentication"

# Check status
./vibe sep status
```

## Documentation

- [Getting Started](getting-started.md) - Installation and setup
- [CLI Reference](cli-reference.md) - All vibe commands
- [Workflow Guide](workflow.md) - How authors, editors, and pilots work together

## Core Concepts

### What is a SEP?

A **Software Enhancement Proposal** is a numbered document that tracks a feature from idea to implementation. SEPs specify *what* to build, not *how* - the AI handles implementation details.

### Roles

| Role | Responsibility |
|------|----------------|
| **Author** | Draft SEPs - anyone on the team can propose features |
| **Editor** | Review SEPs, give feedback, approve for implementation |
| **Pilot** | Implement approved SEPs using AI coding agents |

### SEP Lifecycle

```
DRAFT → ACCEPTED → DONE
  ↓        ↓
BLOCKED  BLOCKED
  ↓
CANCELLED
```

- **DRAFT**: Being written, awaiting editor review
- **ACCEPTED**: Editor approved, ready for implementation
- **BLOCKED**: Can't proceed (dependency, question, etc.)
- **DONE**: Built and shipped
- **CANCELLED**: Decided not to build

### Workflow

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

## Why Vibe?

1. **Lightweight** - Simple markdown files, no complex tooling
2. **AI-Native** - Designed for AI agents to read and implement
3. **Coordination** - Detect conflicts when multiple pilots work in parallel
4. **Traceability** - Every commit links back to a SEP
