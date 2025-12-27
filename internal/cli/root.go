package cli

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "vibe",
	Short: "AI-native development workflow tool",
	Long:  `Vibe is a CLI tool for managing Enhancement Proposals (EPs) in an AI-native development workflow.`,
}
