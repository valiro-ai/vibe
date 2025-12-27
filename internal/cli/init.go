package cli

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/valiro-ai/vibe/internal/templates"
)

var forceInit bool

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize vibe workflow in a repository",
	Long: `Initialize the vibe workflow by creating:
  - docs/seps/          SEP process documentation and template
  - .claude/commands/   Claude Code custom commands for SEP workflow`,
	RunE: func(cmd *cobra.Command, args []string) error {
		created := 0
		skipped := 0

		// Copy SEP templates
		sepsCreated, sepsSkipped, err := copyEmbeddedDir("seps", "docs/seps")
		if err != nil {
			return err
		}
		created += sepsCreated
		skipped += sepsSkipped

		// Copy Claude commands
		cmdCreated, cmdSkipped, err := copyEmbeddedDir("commands", ".claude/commands")
		if err != nil {
			return err
		}
		created += cmdCreated
		skipped += cmdSkipped

		fmt.Printf("\nInitialized vibe workflow: %d files created", created)
		if skipped > 0 {
			fmt.Printf(", %d skipped", skipped)
		}
		fmt.Println()

		fmt.Println("\nNext steps:")
		fmt.Println("  1. Review docs/seps/0000-sep-process.md")
		fmt.Println("  2. Create your first SEP: vibe sep new \"Your Feature\"")
		fmt.Println("  3. Use Claude commands: /sep-status, /sep-new, /sep-implement")

		return nil
	},
}

func copyEmbeddedDir(srcDir, destDir string) (created, skipped int, err error) {
	// Create destination directory
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return 0, 0, fmt.Errorf("failed to create %s: %w", destDir, err)
	}

	// Read embedded directory
	entries, err := fs.ReadDir(templates.FS, srcDir)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to read embedded dir %s: %w", srcDir, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		srcPath := filepath.Join(srcDir, entry.Name())
		destPath := filepath.Join(destDir, entry.Name())

		// Check if file exists
		if _, err := os.Stat(destPath); err == nil && !forceInit {
			fmt.Printf("  Skipped: %s (exists)\n", destPath)
			skipped++
			continue
		}

		// Read and write file
		content, err := templates.FS.ReadFile(srcPath)
		if err != nil {
			return created, skipped, fmt.Errorf("failed to read %s: %w", srcPath, err)
		}

		if err := os.WriteFile(destPath, content, 0644); err != nil {
			return created, skipped, fmt.Errorf("failed to write %s: %w", destPath, err)
		}

		fmt.Printf("  Created: %s\n", destPath)
		created++
	}

	return created, skipped, nil
}

func init() {
	RootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&forceInit, "force", "f", false, "Overwrite existing files")
}
