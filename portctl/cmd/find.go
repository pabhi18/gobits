/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find a Process using a Port Number",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking port:", port)

		command := exec.Command("lsof", "-i", fmt.Sprintf(":%d", port))
		output, err := command.CombinedOutput()

		if err != nil {
			fmt.Println("Error running lsof: ", err)
		}
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(findCmd)

	findCmd.Flags().IntVarP(&port, "port", "p", 0, "Port number to check")
    findCmd.MarkFlagRequired("port")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
