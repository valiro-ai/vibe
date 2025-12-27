package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/valiro-ai/vibe/internal/sep"
)

var updateCmd = &cobra.Command{
	Use:   "update [number] [status]",
	Short: "Update a SEP's status",
	Long: fmt.Sprintf(`Update the status of a SEP.

Valid statuses: %s`, strings.Join(sep.ValidStatuses, ", ")),
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		number := args[0]
		newStatus := strings.ToUpper(args[1])

		// Validate status
		validStatus := false
		for _, s := range sep.ValidStatuses {
			if s == newStatus {
				validStatus = true
				break
			}
		}
		if !validStatus {
			return fmt.Errorf("invalid status: %s\nValid statuses: %s", newStatus, strings.Join(sep.ValidStatuses, ", "))
		}

		// Find the SEP
		foundSEP, err := sep.FindByNumber(sepDir, number)
		if err != nil {
			return err
		}

		oldStatus := foundSEP.Status

		// Update status using proper YAML parsing
		if err := foundSEP.UpdateStatus(newStatus); err != nil {
			return fmt.Errorf("failed to update SEP: %w", err)
		}

		fmt.Printf("Updated SEP-%s: %s → %s\n", foundSEP.Number, oldStatus, newStatus)

		return nil
	},
}

func init() {
	sepCmd.AddCommand(updateCmd)
}
