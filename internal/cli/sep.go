package cli

import (
	"github.com/spf13/cobra"
)

var sepDir = "docs/seps"

var sepCmd = &cobra.Command{
	Use:   "sep",
	Short: "Manage Software Enhancement Proposals",
	Long:  `Commands for creating, listing, and managing Software Enhancement Proposals (SEPs).`,
}

func init() {
	RootCmd.AddCommand(sepCmd)
	sepCmd.PersistentFlags().StringVarP(&sepDir, "dir", "d", "docs/seps", "Directory containing SEP files")
}
