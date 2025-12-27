package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/valiro-ai/vibe/internal/sep"
)

var claimCmd = &cobra.Command{
	Use:   "claim <number> <pilot>",
	Short: "Claim a SEP (assign + commit + push)",
	Long: `Claim a SEP by assigning yourself, committing, and pushing to share with other pilots.

This is equivalent to:
  vibe sep assign <number> <pilot>
  git add <sep-file>
  git commit -m "SEP-<number>: claimed by <pilot>"
  git push

Use "unclaim" or empty pilot to release:
  vibe sep claim 0001 ""`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		number := args[0]
		pilot := args[1]

		// Find the SEP
		foundSEP, err := sep.FindByNumber(sepDir, number)
		if err != nil {
			return err
		}

		// Check if already assigned to someone else
		if foundSEP.Assigned != "" && foundSEP.Assigned != pilot && pilot != "" {
			return fmt.Errorf("SEP-%s is already claimed by %s. Coordinate with them first", number, foundSEP.Assigned)
		}

		// Assign
		oldAssigned := foundSEP.Assigned
		if err := foundSEP.Assign(pilot); err != nil {
			return fmt.Errorf("failed to assign: %w", err)
		}

		// Get relative path for git
		relPath, err := filepath.Rel(".", foundSEP.FilePath)
		if err != nil {
			relPath = foundSEP.FilePath
		}

		// Git add
		gitAdd := exec.Command("git", "add", relPath)
		gitAdd.Stdout = os.Stdout
		gitAdd.Stderr = os.Stderr
		if err := gitAdd.Run(); err != nil {
			return fmt.Errorf("git add failed: %w", err)
		}

		// Git commit
		var commitMsg string
		if pilot == "" {
			commitMsg = fmt.Sprintf("SEP-%s: unclaimed (was %s)", number, oldAssigned)
		} else if oldAssigned == "" {
			commitMsg = fmt.Sprintf("SEP-%s: claimed by %s", number, pilot)
		} else {
			commitMsg = fmt.Sprintf("SEP-%s: reassigned from %s to %s", number, oldAssigned, pilot)
		}

		gitCommit := exec.Command("git", "commit", "-m", commitMsg)
		gitCommit.Stdout = os.Stdout
		gitCommit.Stderr = os.Stderr
		if err := gitCommit.Run(); err != nil {
			return fmt.Errorf("git commit failed: %w", err)
		}

		// Git push
		gitPush := exec.Command("git", "push")
		gitPush.Stdout = os.Stdout
		gitPush.Stderr = os.Stderr
		if err := gitPush.Run(); err != nil {
			return fmt.Errorf("git push failed: %w", err)
		}

		if pilot == "" {
			fmt.Printf("✓ SEP-%s unclaimed and pushed\n", number)
		} else {
			fmt.Printf("✓ SEP-%s claimed by %s and pushed\n", number, pilot)
		}

		return nil
	},
}

func init() {
	sepCmd.AddCommand(claimCmd)
}
