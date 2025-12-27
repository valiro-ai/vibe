package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/valiro-ai/vibe/internal/sep"
	"github.com/valiro-ai/vibe/internal/templates"
)

var newCmd = &cobra.Command{
	Use:   "new [title]",
	Short: "Create a new SEP",
	Long:  `Create a new SEP from the template with the next available number.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		title := args[0]

		// Get next number
		nextNum, err := sep.NextNumber(sepDir)
		if err != nil {
			return fmt.Errorf("failed to determine next SEP number: %w", err)
		}

		// Create slug from title
		slug := createSlug(title)
		filename := fmt.Sprintf("%s-%s.md", nextNum, slug)
		filePath := filepath.Join(sepDir, filename)

		// Check if file already exists
		if _, err := os.Stat(filePath); err == nil {
			return fmt.Errorf("file already exists: %s", filePath)
		}

		// Ensure directory exists
		if err := os.MkdirAll(sepDir, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		// Try local template first, fall back to embedded
		var templateContent []byte
		templatePath := filepath.Join(sepDir, "SEP-TEMPLATE.md")
		if content, err := os.ReadFile(templatePath); err == nil {
			templateContent = content
		} else {
			// Use embedded template
			content, err := templates.FS.ReadFile("seps/SEP-TEMPLATE.md")
			if err != nil {
				return fmt.Errorf("failed to read template: %w", err)
			}
			templateContent = content
		}

		// Replace placeholders
		today := time.Now().Format("2006-01-02")
		content := string(templateContent)
		content = strings.ReplaceAll(content, "SEP-XXXX", fmt.Sprintf("SEP-%s", nextNum))
		content = strings.ReplaceAll(content, "XXXX", nextNum)
		content = strings.ReplaceAll(content, "[Title]", title)
		content = strings.ReplaceAll(content, "YYYY-MM-DD", today)

		// Write new file
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to create SEP: %w", err)
		}

		fmt.Printf("Created: %s\n", filePath)
		fmt.Printf("→ SEP-%s: %s\n", nextNum, title)

		return nil
	},
}

func init() {
	sepCmd.AddCommand(newCmd)
}

func createSlug(title string) string {
	// Convert to lowercase
	slug := strings.ToLower(title)

	// Replace spaces and special chars with hyphens
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	slug = reg.ReplaceAllString(slug, "-")

	// Trim leading/trailing hyphens
	slug = strings.Trim(slug, "-")

	// Truncate if too long
	if len(slug) > 40 {
		slug = slug[:40]
		slug = strings.TrimRight(slug, "-")
	}

	return slug
}
