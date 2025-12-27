package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Pull latest and show pipeline",
	Long: `Sync with remote repository and display the current SEP pipeline.

This is equivalent to:
  git pull
  vibe sep pipeline`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Git pull
		fmt.Println("Pulling latest changes...")
		gitPull := exec.Command("git", "pull")
		gitPull.Stdout = os.Stdout
		gitPull.Stderr = os.Stderr
		if err := gitPull.Run(); err != nil {
			return fmt.Errorf("git pull failed: %w", err)
		}

		fmt.Println()

		// Run pipeline command
		return pipelineCmd.RunE(cmd, args)
	},
}

func init() {
	sepCmd.AddCommand(syncCmd)
}
