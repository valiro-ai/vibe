# SEP-0000: Enhancement Proposal Process

**Status**: DONE
**Created**: 2025-11-18

## What & Why

SEPs (Software Enhancement Proposals) are numbered documents that track features from idea to implementation. They provide a lightweight way to organize work and maintain commit history.

**Key principle**: Trust the AI to know how to build software. SEPs specify *what* to build, not *how*.

## Process

1. **Create**: `/sep-new Feature Name` creates a numbered SEP file
2. **Draft**: Fill in what you want and when it's done
3. **Implement**: `/sep-implement XXXX` builds it
4. **Track**: `/sep-status` shows progress

## States

- **DRAFT**: Not ready to build
- **BLOCKED**: Can't proceed (add reason in SEP)
- **CANCELLED**: Decided not to build
- **DONE**: Built and shipped

## File Structure

```
docs/
  seps/
    0000-sep-process.md       # This file
    0001-feature-name.md      # A feature
    0002-another-feature.md   # Another feature
    SEP-TEMPLATE.md           # Template for new SEPs
    README.md                 # Optional index
```

## SEP Template

```markdown
# SEP-XXXX: [Title]

**Status**: DRAFT | BLOCKED | CANCELLED | DONE
**Created**: YYYY-MM-DD
**Depends on**: SEP-XXXX (optional)

## What & Why

[Explain the feature and why you want it]

## Done When

- [ ] Acceptance criteria 1
- [ ] Acceptance criteria 2
```

## Commit Messages

Reference the SEP number in every commit:

```
SEP-0004: Add notification model
SEP-0004: Implement API endpoints
SEP-0004: Add tests
```

## Rules

1. **Sequential numbering**: 0001, 0002, 0003...
2. **No building from DRAFT**: Only implement when ready
3. **One feature per SEP**: Keep scope focused
4. **Commit references**: Always use `SEP-XXXX:` prefix
5. **Dependencies**: Use "Depends on" field to track ordering

## Quick Reference

```bash
/sep-suggest               # AI suggests new SEP ideas
/sep-new "Feature Name"    # Create new SEP
/sep-discuss XXXX          # Refine requirements
/sep-split XXXX            # Break up large SEP
/sep-implement XXXX        # Build it
/sep-status                # Check progress
```