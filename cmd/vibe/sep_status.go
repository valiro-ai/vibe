package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/valiro-ai/vibe/internal/sep"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show SEP status and recommend next action",
	Long:  `Display current state of all SEPs and recommend what to work on next.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		seps, err := sep.List(sepDir)
		if err != nil {
			return fmt.Errorf("failed to list SEPs: %w", err)
		}

		if len(seps) == 0 {
			fmt.Println("No SEPs found. Run 'vibe init' to get started.")
			return nil
		}

		groups := sep.GroupByStatus(seps)

		fmt.Println("SEP Status")
		fmt.Println(strings.Repeat("=", 40))

		// Show ACCEPTED SEPs (ready for implementation)
		if len(groups[sep.StatusAccepted]) > 0 {
			fmt.Println("\nACCEPTED (ready for implementation):")
			for _, s := range groups[sep.StatusAccepted] {
				deps := ""
				if len(s.DependsOn) > 0 {
					deps = fmt.Sprintf(" [depends on: SEP-%s]", strings.Join(s.DependsOn, ", SEP-"))
				}
				fmt.Printf("  - SEP-%s: %s (created %s)%s\n", s.Number, s.Title, s.Created, deps)
			}
		}

		// Show DRAFT SEPs (awaiting editor review)
		if len(groups[sep.StatusDraft]) > 0 {
			fmt.Println("\nDRAFT (awaiting review):")
			for _, s := range groups[sep.StatusDraft] {
				deps := ""
				if len(s.DependsOn) > 0 {
					deps = fmt.Sprintf(" [depends on: SEP-%s]", strings.Join(s.DependsOn, ", SEP-"))
				}
				fmt.Printf("  - SEP-%s: %s (created %s)%s\n", s.Number, s.Title, s.Created, deps)
			}
		}

		// Show BLOCKED SEPs
		if len(groups[sep.StatusBlocked]) > 0 {
			fmt.Println("\nBLOCKED:")
			for _, s := range groups[sep.StatusBlocked] {
				fmt.Printf("  - SEP-%s: %s (created %s)\n", s.Number, s.Title, s.Created)
			}
		}

		// Show DONE SEPs
		if len(groups[sep.StatusDone]) > 0 {
			fmt.Println("\nDONE:")
			for _, s := range groups[sep.StatusDone] {
				fmt.Printf("  - SEP-%s: %s\n", s.Number, s.Title)
			}
		}

		// Show CANCELLED SEPs
		if len(groups[sep.StatusCancelled]) > 0 {
			fmt.Println("\nCANCELLED:")
			for _, s := range groups[sep.StatusCancelled] {
				fmt.Printf("  - SEP-%s: %s\n", s.Number, s.Title)
			}
		}

		// Recommend next action
		fmt.Println()

		// Priority 1: ACCEPTED SEPs ready for implementation
		if len(groups[sep.StatusAccepted]) > 0 {
			for _, s := range groups[sep.StatusAccepted] {
				canImplement := true
				for _, dep := range s.DependsOn {
					depDone := false
					for _, done := range groups[sep.StatusDone] {
						if done.Number == dep {
							depDone = true
							break
						}
					}
					if !depDone {
						canImplement = false
						break
					}
				}
				if canImplement {
					if len(s.DependsOn) == 0 {
						fmt.Printf("NEXT: Implement SEP-%s with /sep-plan then /sep-implement\n", s.Number)
					} else {
						fmt.Printf("NEXT: Implement SEP-%s (dependencies met)\n", s.Number)
					}
					return nil
				}
			}
		}

		// Priority 2: DRAFT SEPs need editor review
		if len(groups[sep.StatusDraft]) > 0 {
			fmt.Printf("NEXT: Review SEP-%s (editor approval needed before implementation)\n", groups[sep.StatusDraft][0].Number)
			return nil
		}

		// Priority 3: Blocked SEPs
		if len(groups[sep.StatusBlocked]) > 0 {
			fmt.Println("NEXT: Resolve blocked SEPs to continue")
			return nil
		}

		fmt.Println("NEXT: Create a new SEP with 'vibe sep new \"Feature Name\"'")

		return nil
	},
}

func init() {
	sepCmd.AddCommand(statusCmd)
}
