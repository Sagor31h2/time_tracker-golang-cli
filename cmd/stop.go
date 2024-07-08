package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop tracking time for the active task",
	Run: func(cmd *cobra.Command, args []string) {
		if activeTask == "" {
			fmt.Println("No active task to stop")
			return
		}
		stopTask(activeTask)
	},
}

func stopTask(name string) {
	task, exists := tasks[name]
	if !exists {
		fmt.Printf("Task not found: %s\n", name)
		return
	}
	duration := time.Since(task.StartTime)
	delete(tasks, name)
	if activeTask == name {
		activeTask = ""
	}
	saveTasks()
	fmt.Printf("Stopped tracking time for: %s (Duration: %s)\n", name, duration.Round(time.Second))
}
