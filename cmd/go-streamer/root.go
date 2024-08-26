package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-streamer",
	Short: "Simple CLI to interact with Kafka to send and receive logs from specified file paths",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Global flags or persistent flags can be added here
}
