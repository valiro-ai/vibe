# Workflow Guide

How authors, editors, and pilots work together using vibe.

## Roles

### Authors

Authors draft SEPs. Anyone on the team can be an author:

- Propose new features
- Write initial SEP drafts
- Define what they want and why

Authors don't need to know how to code. They focus on what would be valuable.

### Editors

Editors review and approve SEPs:

- Give feedback on SEP drafts
- Ensure acceptance criteria are clear
- Approve SEPs for implementation
- Review completed work

Editors bridge business needs and technical feasibility.

### Pilots

Pilots implement approved SEPs:

- Run `/sep-plan` and `/sep-implement`
- Coordinate with other pilots to avoid conflicts
- Update SEP status as work progresses
- Add implementation notes

Pilots understand the codebase and guide AI through complex implementations.

## Author Workflow

### 1. Create a SEP

```bash
vibe sep new "Feature Name"
```

### 2. Define the Feature

Edit the SEP to include:

**What & Why** - Explain the feature in plain language:
```markdown
## What & Why

Users should be able to reset their password via email. Currently,
if they forget their password, they have to contact support. This
causes 50+ support tickets per week.
```

**Done When** - List specific acceptance criteria:
```markdown
## Done When

- [ ] User can request password reset from login page
- [ ] Reset link sent via email expires after 1 hour
- [ ] User can set new password using reset link
- [ ] User receives confirmation email after reset
- [ ] Failed reset attempts are logged for security
```

**Areas** - List files/directories that will change:
```markdown
areas:
  - internal/auth/*
  - api/routes/password.go
  - templates/email/password-reset.html
```

### 3. Submit for Review

Hand off to an editor for review. Status stays DRAFT until approved.

## Editor Workflow

### 1. Review SEP Draft

Read the SEP and evaluate:
- Is the problem clearly defined?
- Are acceptance criteria specific and testable?
- Are the areas reasonable?

### 2. Give Feedback

Use Claude to discuss the SEP:

```
/sep-discuss 0001
```

Claude will help identify gaps and suggest improvements.

### 3. Approve for Implementation

When satisfied, update status to ACCEPTED:

```bash
vibe sep update 0001 ACCEPTED
```

This signals to pilots that the SEP is ready for implementation.

### 4. Review Completed Work

After implementation, verify:
- Acceptance criteria are met
- Implementation matches intent
- No unintended side effects

## Pilot Workflow

### 1. Check the Pipeline

```bash
vibe sep pipeline
```

Review active SEPs and check for conflicts:

```
SEP Pipeline - Area Conflicts
==================================================

DRAFT:
  SEP-0001: User Authentication ⚠️  CONFLICT with SEP-0002
    areas: internal/auth/*, internal/user/*
  SEP-0002: User Profile Management ⚠️  CONFLICT with SEP-0001
    areas: internal/user/*, api/routes/profile.go

--------------------------------------------------
⚠️  Conflicts detected:
  SEP-0001 ↔ SEP-0002: internal/user/*

→ Implement conflicting SEPs sequentially or coordinate pilots
```

### 2. Sync and Check Pipeline

Pull latest changes and see who's working on what:

```bash
vibe sep sync
```

### 3. Claim a SEP

Claim the SEP (assigns, commits, and pushes automatically):

```bash
vibe sep claim 0001 @yourname
```

This makes your claim visible to all other pilots immediately.

If a SEP conflicts with one already claimed:
- Coordinate with the assigned pilot
- Work sequentially (wait for them to finish)
- Split the work (divide the shared areas)

### 4. Create Implementation Plan

```
/sep-plan 0001
```

Claude will:
1. Read the SEP document
2. Explore the current codebase
3. Create a concrete implementation plan
4. Write the plan to the SEP

### 5. Review the Plan

Check the "## Plan" section in the SEP. Verify:
- Files to modify are correct
- Approach makes sense
- No major concerns

If changes needed, discuss with editor or re-run `/sep-plan`.

### 6. Implement

```
/sep-implement 0001
```

Claude will:
1. Follow the plan step-by-step
2. Check off completed acceptance criteria
3. Add implementation notes
4. Update status to DONE

### 7. Commit with SEP Reference

Include the SEP number in every commit:

```
SEP-0001: Add password reset request endpoint
SEP-0001: Implement email sending for reset link
SEP-0001: Add password reset confirmation page
```

This creates traceability from commits to features.

### 8. Update Status

When complete:

```bash
vibe sep update 0001 DONE
```

If blocked:

```bash
vibe sep update 0001 BLOCKED
```

Then add a note in the SEP explaining why it's blocked.

## Parallel Work

When multiple pilots work simultaneously:

### Check for Conflicts

```bash
vibe sep pipeline
```

The pipeline shows which SEPs touch the same files. Conflicts mean:
- Two SEPs modify the same code
- Merge conflicts are likely if implemented in parallel

### Resolve Conflicts

**Option 1: Sequential Implementation**

Pilot A completes SEP-0001, then Pilot B starts SEP-0002.

**Option 2: Area Assignment**

Split shared areas:
- Pilot A handles `internal/user/auth.go`
- Pilot B handles `internal/user/profile.go`

Update the SEP `areas` field to reflect the split.

**Option 3: Pair Implementation**

Both pilots work together on the shared areas, communicating changes in real-time.

## Best Practices

### For Authors

1. **Be specific** - Vague criteria lead to wrong implementations
2. **Think in outcomes** - What should the user be able to do?
3. **List edge cases** - What happens on error? Empty state?

### For Editors

1. **Push back on vague SEPs** - Better to clarify now than fix later
2. **Define areas** - Even a rough guess helps pilots coordinate
3. **Consider dependencies** - Does this SEP depend on others?

### For Pilots

1. **Check pipeline first** - Know what others are working on
2. **Claim before starting** - Avoid duplicate work
3. **Commit frequently** - Small commits with SEP references
4. **Add notes** - Future pilots will thank you
5. **Update status promptly** - Keep the pipeline accurate

### For Teams

1. **Daily sync** - Quick check on who's working on what
2. **One SEP at a time** - Complete before starting the next
3. **Review together** - Editor + pilot review completed work
4. **Iterate** - SEPs can spawn follow-up SEPs
