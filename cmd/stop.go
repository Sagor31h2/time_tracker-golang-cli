/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop tracking time",
	Run: func(cmd *cobra.Command, args []string) {
		if activeTask == "" {
			fmt.Println("No active task to stop")

		} else {
			stopActiveTask()
		}

	},
}

func stopActiveTask() {
	taskName := activeTask
	startTime, exists := tasks[taskName]
	if !exists {
		fmt.Printf("Task not found: %s\n", taskName)
		return
	}
	duration := time.Since(startTime)
	stopTask(taskName)
	fmt.Printf("Stopped tracking time for: %s (Duration: %s)\n", taskName, duration.Round(time.Second))
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
