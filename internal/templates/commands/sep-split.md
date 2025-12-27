---
description: Split a large DRAFT SEP into multiple focused SEPs
argument-hint: XXXX
---

1. Read `docs/seps/$1-*.md`
2. Verify status is DRAFT (can't split DONE SEPs)
3. Analyze the scope and identify logical splits
4. Propose split:
   - How many SEPs (typically 2-3)
   - What each would cover
   - Suggested dependency order
5. Ask for confirmation
6. Create new SEPs:
   - Find next available numbers
   - Create new SEP files
   - Distribute content from original
   - Set up dependencies
7. Update original SEP:
   - Mark as DONE (obsoleted by split)
   - Add note pointing to new SEPs
8. Update README.md if it exists

Output:
```
SEP-$1: [Original Title]

This SEP covers too much. I suggest splitting into:

SEP-YYYY: [Focused Feature 1]
- [Scope description]
- Depends on: None

SEP-ZZZZ: [Focused Feature 2]
- [Scope description]
- Depends on: SEP-YYYY

SEP-AAAA: [Focused Feature 3]
- [Scope description]
- Depends on: SEP-YYYY

Proceed with split? (y/n)
```

**After confirmation:**
```
Split complete:
✓ Created SEP-YYYY: [Feature 1]
✓ Created SEP-ZZZZ: [Feature 2]
✓ Created SEP-AAAA: [Feature 3]
✓ Marked SEP-$1 as superseded

Next: Review and discuss each new SEP with /sep-discuss
```

**Rules:**
- Original SEP status → DONE with note "Split into SEP-X, SEP-Y, SEP-Z"
- Each new SEP is DRAFT and focused
- Dependency order is suggested based on logical sequence
