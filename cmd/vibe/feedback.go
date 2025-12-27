package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	feedbackSEP  string
	feedbackFile string
)

var feedbackCmd = &cobra.Command{
	Use:   "feedback [message]",
	Short: "Submit feedback about vibe or a SEP",
	Long: `Submit feedback, suggestions, or issues. Feedback is appended to a log file.

Examples:
  vibe feedback "The claim command is really useful"
  vibe feedback --sep 0001 "Acceptance criteria could be clearer"
  vibe feedback  # Opens prompt for multi-line feedback`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var message string

		if len(args) > 0 {
			message = strings.Join(args, " ")
		} else {
			// Interactive mode
			fmt.Println("Enter feedback (empty line to finish):")
			scanner := bufio.NewScanner(os.Stdin)
			var lines []string
			for scanner.Scan() {
				line := scanner.Text()
				if line == "" {
					break
				}
				lines = append(lines, line)
			}
			if err := scanner.Err(); err != nil {
				return err
			}
			message = strings.Join(lines, "\n")
		}

		if strings.TrimSpace(message) == "" {
			return fmt.Errorf("feedback message cannot be empty")
		}

		// Ensure directory exists
		dir := filepath.Dir(feedbackFile)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		// Open file for appending
		f, err := os.OpenFile(feedbackFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("failed to open feedback file: %w", err)
		}
		defer f.Close()

		// Format entry
		timestamp := time.Now().Format("2006-01-02 15:04")
		var entry string
		if feedbackSEP != "" {
			entry = fmt.Sprintf("[%s] SEP-%s: %s\n", timestamp, feedbackSEP, message)
		} else {
			entry = fmt.Sprintf("[%s] %s\n", timestamp, message)
		}

		// Write
		if _, err := f.WriteString(entry); err != nil {
			return fmt.Errorf("failed to write feedback: %w", err)
		}

		fmt.Println("✓ Feedback recorded")
		return nil
	},
}

var feedbackListCmd = &cobra.Command{
	Use:   "list",
	Short: "View recorded feedback",
	RunE: func(cmd *cobra.Command, args []string) error {
		content, err := os.ReadFile(feedbackFile)
		if os.IsNotExist(err) {
			fmt.Println("No feedback recorded yet.")
			return nil
		}
		if err != nil {
			return fmt.Errorf("failed to read feedback: %w", err)
		}

		if len(content) == 0 {
			fmt.Println("No feedback recorded yet.")
			return nil
		}

		fmt.Println("Feedback Log")
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println(string(content))
		return nil
	},
}

var feedbackClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all feedback (after reviewing)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := os.Remove(feedbackFile); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to clear feedback: %w", err)
		}
		fmt.Println("✓ Feedback cleared")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(feedbackCmd)
	feedbackCmd.AddCommand(feedbackListCmd)
	feedbackCmd.AddCommand(feedbackClearCmd)

	feedbackCmd.Flags().StringVar(&feedbackSEP, "sep", "", "Link feedback to a specific SEP number")
	feedbackCmd.PersistentFlags().StringVar(&feedbackFile, "file", "docs/feedback.log", "Feedback log file")
}
