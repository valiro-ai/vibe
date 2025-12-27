package main

import (
	"os"

	"github.com/valiro-ai/vibe/internal/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
