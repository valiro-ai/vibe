package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/valiro-ai/vibe/internal/sep"
)

var assignCmd = &cobra.Command{
	Use:   "assign <number> <pilot>",
	Short: "Assign a pilot to a SEP",
	Long: `Assign a pilot to work on a SEP. Use "unassign" or empty string to remove assignment.

Examples:
  vibe sep assign 0001 @alice
  vibe sep assign 0001 ""        # unassign`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		number := args[0]
		pilot := args[1]

		foundSEP, err := sep.FindByNumber(sepDir, number)
		if err != nil {
			return err
		}

		oldAssigned := foundSEP.Assigned
		if oldAssigned == "" {
			oldAssigned = "(unassigned)"
		}

		if err := foundSEP.Assign(pilot); err != nil {
			return fmt.Errorf("failed to assign: %w", err)
		}

		newAssigned := pilot
		if newAssigned == "" {
			newAssigned = "(unassigned)"
		}

		fmt.Printf("Updated %s: %s → %s\n", foundSEP.ID(), oldAssigned, newAssigned)
		return nil
	},
}

func init() {
	sepCmd.AddCommand(assignCmd)
}
