package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [task name]",
	Short: "Start tracking time for a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		startTask(args[0])
	},
}

func startTask(name string) {
	if activeTask != "" {
		stopTask(activeTask)
	}
	tasks[name] = Task{Name: name, StartTime: time.Now()}
	activeTask = name
	saveTasks()
	fmt.Printf("Started tracking time for: %s\n", name)
}
