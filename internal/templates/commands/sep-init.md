---
description: Bootstrap SEP process in your project
argument-hint: "[path-to-prd.md]"
---

Initialize the SEP system, optionally importing an existing PRD as SEP-0001.

**Two modes:**

## Mode 1: Fresh start (no argument)

1. Create `docs/seps/` directory if needed
2. Copy SEP process files:
   - `0000-sep-process.md` (the process doc)
   - `SEP-TEMPLATE.md` (template)
   - `README.md` (index)
3. Create `.claude/commands/` directory if needed
4. Copy SEP commands:
   - `sep-new.md`
   - `sep-discuss.md`
   - `sep-split.md`
   - `sep-implement.md`
   - `sep-status.md`

Output:
```
SEP system initialized ✓

Created:
- docs/seps/0000-sep-process.md
- docs/seps/SEP-TEMPLATE.md
- docs/seps/README.md
- .claude/commands/sep-*.md

Next: /sep-new "Your First Feature"
```

## Mode 2: Import existing PRD (with argument)

1. Do everything from Mode 1
2. Read the PRD file at `$1`
3. Convert to SEP-0001:
   - Extract title
   - Parse into "What & Why" section
   - Extract/identify acceptance criteria as "Done When"
   - Create `docs/seps/0001-[slug].md`
4. Update `docs/seps/README.md` with SEP-0001
5. Set status: DRAFT

Output:
```
SEP system initialized with PRD import ✓

Created:
- docs/seps/0000-sep-process.md
- docs/seps/SEP-TEMPLATE.md
- docs/seps/README.md
- docs/seps/0001-[slug].md (from your PRD)
- .claude/commands/sep-*.md

Imported PRD as SEP-0001: [Title]
Status: DRAFT

Next: /sep-discuss 0001 to refine, then /sep-implement 0001
```

**Examples:**

```bash
# Fresh start
/sep-init

# Import existing PRD
/sep-init docs/mvp-requirements.md
/sep-init README.md
/sep-init docs/product/initial-spec.md
```

**Notes:**
- Safe to run if files already exist (won't overwrite)
- PRD can be any markdown file with requirements
- Intelligently extracts content and structure from PRD
