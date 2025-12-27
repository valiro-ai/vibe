package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vibe",
	Short: "AI-native development workflow tool",
	Long:  `Vibe is a CLI tool for managing Enhancement Proposals (EPs) in an AI-native development workflow.`,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
