/*
Copyright © 2025 Abhinav Pratap <youremail@example.com>
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dockwatch",
	Short: "Docker Container Resource Watcher",
	Long:  `dockwatch is a lightweight CLI tool to monitor Docker containers' CPU, memory usage, and uptime in real-time.`,
	Run: func(cmd *cobra.Command, args []string) {
		// If no subcommand is provided, default to the list command
		listCmd.Run(cmd, args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// It is called by main.main().
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Global flags (if any) can be defined here using rootCmd.PersistentFlags()

	// Example local flag (not used in this tool yet)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}