package templates

import "embed"

//go:embed seps/* commands/*
var FS embed.FS

// Embedded directory structure:
//
// seps/
//   0000-sep-process.md  - The SEP process documentation
//   SEP-TEMPLATE.md      - Template for new SEPs
//
// commands/
//   sep-init.md          - Initialize SEP process
//   sep-new.md           - Create new SEP
//   sep-discuss.md       - Discuss/refine SEP
//   sep-implement.md     - Implement SEP
//   sep-split.md         - Split large SEP
//   sep-suggest.md       - Suggest new SEPs
//   sep-status.md        - Show SEP status
