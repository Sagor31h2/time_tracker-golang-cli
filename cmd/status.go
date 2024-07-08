package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status of all tracked tasks and their durations",
	Run: func(cmd *cobra.Command, args []string) {
		listTasks()
	},
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks are currently being tracked.")
		return
	}

	fmt.Println("Current tasks:")
	for _, task := range tasks {
		duration := time.Since(task.StartTime)
		if task.Name == activeTask {
			fmt.Printf("- %s: %s (active)\n", task.Name, duration.Round(time.Second))
		} else {
			fmt.Printf("- %s: %s\n", task.Name, duration.Round(time.Second))
		}
	}
}
