package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/valiro-ai/vibe/internal/sep"
)

var pipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Show SEP pipeline with area conflicts",
	Long:  `Display all active SEPs with their areas and highlight potential conflicts for pilot coordination.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		seps, err := sep.List(sepDir)
		if err != nil {
			return fmt.Errorf("failed to list SEPs: %w", err)
		}

		if len(seps) == 0 {
			fmt.Println("No SEPs found. Run 'vibe init' to get started.")
			return nil
		}

		// Find conflicts
		conflicts := sep.FindConflicts(seps)
		conflictMap := make(map[string][]string) // SEP number -> list of conflicting SEP numbers

		for _, c := range conflicts {
			conflictMap[c.SEP1.Number] = append(conflictMap[c.SEP1.Number], c.SEP2.Number)
			conflictMap[c.SEP2.Number] = append(conflictMap[c.SEP2.Number], c.SEP1.Number)
		}

		groups := sep.GroupByStatus(seps)

		fmt.Println("SEP Pipeline - Area Conflicts")
		fmt.Println(strings.Repeat("=", 50))

		// Show active SEPs (ACCEPTED, DRAFT, BLOCKED)
		activeStatuses := []string{sep.StatusAccepted, sep.StatusDraft, sep.StatusBlocked}
		for _, status := range activeStatuses {
			statusSeps, ok := groups[status]
			if !ok || len(statusSeps) == 0 {
				continue
			}

			fmt.Printf("\n%s:\n", status)

			for _, s := range statusSeps {
				// Check for conflicts
				conflictsWith, hasConflict := conflictMap[s.Number]
				conflictMarker := ""
				if hasConflict {
					conflictMarker = fmt.Sprintf(" ⚠️  CONFLICT with SEP-%s", strings.Join(conflictsWith, ", SEP-"))
				}

				// Show assignment
				assignedMarker := ""
				if s.Assigned != "" {
					assignedMarker = fmt.Sprintf(" [%s]", s.Assigned)
				}

				fmt.Printf("  SEP-%s: %s%s%s\n", s.Number, s.Title, assignedMarker, conflictMarker)

				// Show areas
				if len(s.Areas) > 0 {
					fmt.Printf("    areas: %s\n", strings.Join(s.Areas, ", "))
				} else {
					fmt.Printf("    areas: (not specified)\n")
				}

				// Show dependencies
				if len(s.DependsOn) > 0 {
					fmt.Printf("    depends_on: SEP-%s\n", strings.Join(s.DependsOn, ", SEP-"))
				}
			}
		}

		// Show DONE count
		if len(groups[sep.StatusDone]) > 0 {
			fmt.Printf("\nDONE: %d SEPs completed\n", len(groups[sep.StatusDone]))
		}

		// Show conflict details
		if len(conflicts) > 0 {
			fmt.Println()
			fmt.Println(strings.Repeat("-", 50))
			fmt.Println("⚠️  Conflicts detected:")
			for _, c := range conflicts {
				// Show who's assigned if anyone
				assignInfo := ""
				if c.SEP1.Assigned != "" && c.SEP2.Assigned != "" {
					assignInfo = fmt.Sprintf(" (%s vs %s)", c.SEP1.Assigned, c.SEP2.Assigned)
				} else if c.SEP1.Assigned != "" {
					assignInfo = fmt.Sprintf(" (%s assigned)", c.SEP1.Assigned)
				} else if c.SEP2.Assigned != "" {
					assignInfo = fmt.Sprintf(" (%s assigned)", c.SEP2.Assigned)
				}
				fmt.Printf("  SEP-%s ↔ SEP-%s: %s%s\n",
					c.SEP1.Number,
					c.SEP2.Number,
					strings.Join(c.OverlapAreas, ", "),
					assignInfo)
			}
			fmt.Println("\n→ Coordinate with assigned pilots or implement sequentially")
		}

		return nil
	},
}

func init() {
	sepCmd.AddCommand(pipelineCmd)
}
