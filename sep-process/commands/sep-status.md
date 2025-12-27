---
description: Show current state of all SEPs and recommend next action
---

1. List all SEP files in `docs/seps/`
2. Read status from each
3. Group by status: DRAFT, BLOCKED, CANCELLED, DONE
4. Show dependencies if any
5. Recommend next action

Output:
```
SEP Status
==========

DRAFT:
- SEP-0003: Feature X (created 2025-11-18)
- SEP-0002: Feature Y (created 2025-11-15) [depends on: SEP-0001]

BLOCKED:
- SEP-0004: Feature Z (created 2025-11-16) - waiting for API access

DONE:
- SEP-0001: Initial MVP (completed 2025-11-10)
- SEP-0000: SEP Process (completed 2025-11-18)

CANCELLED:
- SEP-0005: Old Idea (cancelled 2025-11-12)

NEXT: Implement SEP-0003 (no dependencies) or discuss blocked items
```
