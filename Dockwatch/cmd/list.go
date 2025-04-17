/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os/exec"
	"github.com/spf13/cobra"
	"strings"
	"time"
    "text/tabwriter"
    "os"
    "github.com/fatih/color"
)

type Container struct {
	ID string
	Name string
}

type ContainerStats struct {
    Name     string
    CPU      string
    Memory   string
    Uptime   string
    Error    error
}

var containers []Container

var containersStats []ContainerStats

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ListContainers()
		ContainerStatsList()
	},
}

func ListContainers() {
	cmd := exec.Command("docker", "ps", "--format", "{{.ID}} {{.Names}}")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines{
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, " ", 2)
		if len(parts)==2{
			var c Container
			c.ID = parts[0]
			c.Name = parts[1]
			containers = append(containers, c)
		}
	}
}

func GetContainerStats(conatiner Container) ContainerStats {
	stats := ContainerStats{Name: conatiner.Name}

	cmd := exec.Command("docker", "stats", conatiner.ID, "--no-stream", "--format", "{{.CPUPerc}} {{.MemUsage}}")
	output, err := cmd.CombinedOutput()

	if err != nil {
		stats.Error = fmt.Errorf("Failed to get stats: %v", err)
		return stats
	}

	parts := strings.Split(string(output), " ")
    if len(parts) >= 2 {
        stats.CPU = parts[0]
        stats.Memory = parts[1]
    }

	cmd = exec.Command("docker", "inspect", "-f", "{{.State.StartedAt}}", conatiner.ID)
	output, err = cmd.CombinedOutput()

	if err != nil {
		stats.Error = fmt.Errorf("Failed to get stats: %v", err)
	}

	startedAtstr := strings.TrimSpace(string(output))

	startedAt, err := time.Parse(time.RFC3339Nano, startedAtstr)

	if err == nil {
		duration := time.Since(startedAt).Round(time.Minute)
		stats.Uptime = duration.String()
	} else {
		stats.Uptime = "Unknown"
	}
	return stats
}

func PrintStatsTable() {
    w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

    // Color functions
    headerColor := color.New(color.FgCyan).SprintFunc()
    nameColor := color.New(color.FgYellow).SprintFunc()
    cpuColor := color.New(color.FgGreen).SprintFunc()
    memColor := color.New(color.FgMagenta).SprintFunc()
    uptimeColor := color.New(color.FgBlue).SprintFunc()
    errorColor := color.New(color.FgRed, color.Bold).SprintFunc()

    // Print header
    fmt.Fprintln(w, headerColor("NAME")+"\t"+
        headerColor("CPU")+"\t"+
        headerColor("MEMORY")+"\t"+
        headerColor("UPTIME")+"\t")

    for _, stat := range containersStats {
        if stat.Error != nil {
            fmt.Fprintf(w, "%s\t%s\t\t\t\n",
                nameColor(stat.Name),
                errorColor("ERROR: "+stat.Error.Error()),
            )
        } else {
            fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n",
                nameColor(stat.Name),
                cpuColor(stat.CPU),
                memColor(stat.Memory),
                uptimeColor(stat.Uptime),
            )
        }
    }

    w.Flush()
}

func ContainerStatsList() {
	for _, c := range containers {
		stats := GetContainerStats(c)
		containersStats = append(containersStats, stats)
	}
	PrintStatsTable()
}


func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
