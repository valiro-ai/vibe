---
description: Create implementation plan for a DRAFT SEP
argument-hint: XXXX
---

1. Read `docs/seps/$1-*.md`
2. Verify status is DRAFT
3. Understand "What & Why" and "Done When" criteria
4. Explore the current codebase:
   - Check files listed in `areas`
   - Understand existing patterns and architecture
   - Identify dependencies and integration points
5. Create a concrete implementation plan:
   - List specific files to create/modify
   - Outline changes for each file
   - Note any new dependencies needed
   - Identify potential risks or blockers
6. Write the plan to the "## Plan" section of the SEP
7. Update `areas` if the plan reveals additional files

Output:
```
SEP-$1: [Title] - Plan Created

Files to modify:
- path/to/file1.go (add X, modify Y)
- path/to/file2.go (new file)

Steps:
1. [Step 1]
2. [Step 2]
3. [Step 3]

Risks:
- [Any concerns or blockers]

→ Review the plan, then run /sep-implement $1
```

**Important:** The plan is a proposal based on current codebase state. Review before implementing.
