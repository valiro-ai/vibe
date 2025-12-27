---
description: Review and discuss the feature requirements for a DRAFT SEP
argument-hint: XXXX
---

1. Read `docs/seps/$1-*.md`
2. Show SEP summary:
   - Number and title
   - Status
   - What & Why (current content)
   - Done When criteria (current list)
3. Analyze and identify potential issues:
   - Is scope too large? (suggest split)
   - Is scope unfocused? (suggest split)
   - Missing acceptance criteria?
   - Unclear requirements?
   - Technical risks or dependencies?
4. Present assessment and ask what to discuss

Output:
```
SEP-$1: [Title]
Status: DRAFT
Created: YYYY-MM-DD

What & Why:
[Show current content or "Not yet filled in"]

Done When:
- [ ] Criteria 1
- [ ] Criteria 2
[Or "Not yet specified"]

Assessment:
⚠ Scope seems large - could split into 2-3 focused SEPs
✓ Acceptance criteria clear
? Missing details about [specific area]

What would you like to discuss or refine?
```

**During discussion:**
- Update SEP content based on conversation
- Add/refine "Done When" criteria as discovered
- Suggest `/sep-split $1` if scope is too large
- Save changes to the SEP file
- Iterate until requirements are clear

**Purpose:** Collaboratively refine requirements, discover missing details, and ensure SEP is focused and ready.
