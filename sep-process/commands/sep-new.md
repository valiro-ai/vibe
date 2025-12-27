---
description: Create a new SEP from template
argument-hint: "Feature Name"
---

1. Find next SEP number (check existing files, use highest + 1)
2. Slugify feature name (lowercase, hyphens)
3. Create `docs/seps/XXXX-slug.md` from `SEP-TEMPLATE.md`
4. Fill in:
   - Number: XXXX (zero-padded)
   - Title: $1
   - Status: DRAFT
   - Created: today (YYYY-MM-DD)
5. Update `docs/seps/README.md` if it exists

Output:
```
Created SEP-XXXX: $1
File: docs/seps/XXXX-slug.md

Next: Fill in "What & Why" and "Done When" sections
```
