/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"os/exec"
	"github.com/spf13/cobra"
	"strings"
	"strconv"
)

// killCmd represents the kill command
var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kill a process running on a port using the port number or a process ID",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage for this command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the necessary files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		killProcess()
	},
}

func killProcess() {
	if port == 0 && processId == 0 {
		fmt.Println("Please provide either a port (--port) or a process ID (--pid).")
		return
	}
	if port != 0 {
		cmd := exec.Command("sh", "-c", fmt.Sprintf("lsof -i :%s | awk 'NR>1 { print $2; exit }'", strconv.Itoa(port)))
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error finding process:", err)
			return
		}

		pid := strings.TrimSpace(string(output))

		if pid == "" {
			fmt.Printf("No process is running on port: %d\n", port)
			return
		}

		fmt.Printf("🔍 Process found: PID: %s\n", pid)

		killCmd := exec.Command("kill", "-9", pid)
		killOutput, err := killCmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error killing process:", err)
			return
		}
		fmt.Println("✅ Process killed! ", string(killOutput))
	}
	if processId != 0 {
		killCmd := exec.Command("kill", "-9", fmt.Sprintf("%d", processId))
		output, err := killCmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error killing process by PID: ", err)
			return
		}
		fmt.Println("✅ Process killed! ", string(output))
	}
}

func init() {
	rootCmd.AddCommand(killCmd)

	killCmd.Flags().IntVarP(&port, "port", "p", 0, "Port to kill the process running on")
	killCmd.Flags().IntVarP(&processId, "pid", "i", 0, "Process ID to kill directly")	
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// killCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// killCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}