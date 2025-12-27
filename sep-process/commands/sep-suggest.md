---
description: Suggest new SEP ideas based on current project state
---

1. Read all SEP files from `docs/seps/`
2. Analyze current project state:
   - What SEPs are DONE/COMPLETED
   - What SEPs are DRAFT/active
   - What features exist in the codebase
   - What gaps or natural next steps exist
3. Review CLAUDE.md and codebase architecture
4. Generate 3-5 SEP suggestions that would:
   - Build on completed SEPs
   - Fill gaps in the feature set
   - Enhance user experience
   - Follow natural product evolution
   - Be appropriately scoped (not too large)
5. Write suggestions to docs/seps/SUGGESTIONS.md
   - Suggested title
   - Brief description (2-3 sentences)
   - Why it makes sense now
   - Estimated complexity (Small/Medium/Large)
   - Possible dependencies
6. Present suggestions with:
   - Suggested title
   - Brief description (2-3 sentences)
   - Why it makes sense now
   - Estimated complexity (Small/Medium/Large)
   - Possible dependencies

Output:
```
SEP Suggestions
===============

Based on current project state:
- ✓ Completed: SEP-0001, SEP-0002, SEP-0003, SEP-0004
- 🚧 Active: SEP-0005, SEP-0006

Suggestions:

1. [Title] (Small/Medium/Large)
   Description: [2-3 sentences]
   Why now: [Rationale]
   Depends on: [SEP-XXXX or None]

2. [Title] (Small/Medium/Large)
   Description: [2-3 sentences]
   Why now: [Rationale]
   Depends on: [SEP-XXXX or None]

3. [Title] (Small/Medium/Large)
   Description: [2-3 sentences]
   Why now: [Rationale]
   Depends on: [SEP-XXXX or None]

Which would you like to explore? (Reply with number 1-5, or suggest your own)
```

**After user selects:**
1. Create new SEP using next available number
2. Fill in "What & Why" section based on suggestion
3. Generate initial "Done When" criteria
4. Show created SEP summary

**Purpose:** Proactively discover valuable features to build next, based on project context and natural product evolution.
