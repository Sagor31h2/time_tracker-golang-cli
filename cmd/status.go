/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status of all tracked tasks and their durations",
	Run: func(cmd *cobra.Command, args []string) {
		if len(tasks) <= 0 {
			fmt.Println("No Tasks are currently being tracked.")
		} else {

			fmt.Println("Current tracking status.")
			for task, startTime := range tasks {
				duration := time.Since(startTime)
				if task == activeTask {
					fmt.Printf("- %s: %s (active)\n", task, duration.Round(time.Second))
				} else {
					fmt.Printf("- %s: %s\n", task, duration.Round(time.Second))
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
