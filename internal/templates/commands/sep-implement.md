---
description: Execute an ACCEPTED SEP step-by-step (user)
argument-hint: XXXX
---

1. Read `docs/seps/$1-*.md`
2. Verify status is ACCEPTED (editor-approved):
   - If DRAFT: Stop and inform that editor approval is needed first
   - If ACCEPTED: Proceed with implementation
3. Check if "## Plan" section has content:
   - If empty: Suggest running `/sep-plan $1` first, or ask if user wants to proceed without a plan
   - If present: Use the plan as implementation guide
4. Follow the plan step-by-step:
   - Create/modify files as specified
   - Check each step against "Done When" criteria
5. For each acceptance criterion completed:
   - Check off the item: `- [ ]` → `- [x]`
6. Run tests, type check, lint
7. Add implementation notes to "## Implementation Notes"
8. Update status: ACCEPTED → DONE
9. Commit with `SEP-$1:` prefix

Output:
```
SEP-$1: [Title] - DONE ✓

Done When:
✓ [Criteria 1]
✓ [Criteria 2]
✓ [Criteria 3]

Commits: [list]
```

**If blocked:** Stop and ask for clarification instead of guessing.
**If plan is outdated:** Note discrepancies in Implementation Notes and adapt as needed.
