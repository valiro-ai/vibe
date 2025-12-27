package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/valiro-ai/vibe/internal/sep"
)

var listStatus string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all SEPs",
	Long:  `List all SEPs, optionally filtered by status.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		seps, err := sep.List(sepDir)
		if err != nil {
			return fmt.Errorf("failed to list SEPs: %w", err)
		}

		if len(seps) == 0 {
			fmt.Println("No SEPs found.")
			return nil
		}

		// Filter by status if specified
		if listStatus != "" {
			var filtered []*sep.SEP
			for _, s := range seps {
				if s.Status == listStatus {
					filtered = append(filtered, s)
				}
			}
			seps = filtered
		}

		groups := sep.GroupByStatus(seps)

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		// Display order: ACCEPTED, DRAFT, BLOCKED, DONE, CANCELLED
		displayOrder := []string{sep.StatusAccepted, sep.StatusDraft, sep.StatusBlocked, sep.StatusDone, sep.StatusCancelled}

		for _, status := range displayOrder {
			statusSeps, ok := groups[status]
			if !ok || len(statusSeps) == 0 {
				continue
			}

			fmt.Fprintf(w, "\n%s:\n", status)

			for _, s := range statusSeps {
				title := truncate(s.Title, 50)
				deps := ""
				if len(s.DependsOn) > 0 {
					deps = fmt.Sprintf(" [depends on: SEP-%s]", s.DependsOn[0])
				}
				fmt.Fprintf(w, "  SEP-%s\t%s\t(created %s)%s\n", s.Number, title, s.Created, deps)
			}
		}

		w.Flush()

		// Print summary
		fmt.Printf("\n---\nTotal: %d SEPs\n", len(seps))

		return nil
	},
}

func init() {
	sepCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&listStatus, "status", "s", "", "Filter by status (DRAFT, ACCEPTED, BLOCKED, DONE, CANCELLED)")
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
