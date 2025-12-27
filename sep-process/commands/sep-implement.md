---
description: Execute an ACCEPTED SEP step-by-step
argument-hint: XXXX
---

1. Read `docs/seps/$1-*.md`
2. Verify status is DRAFT (ready to build)
3. Understand "What & Why" and "Done When"
4. Build the feature using my judgment
5. Check off "Done When" items as completed
6. Run tests, type check, lint
7. Update status: DRAFT → DONE
8. Update `docs/seps/README.md` if it exists
9. Add implementation notes to SEP
10. Commit with `SEP-$1:` prefix

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
