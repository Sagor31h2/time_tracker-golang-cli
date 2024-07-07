/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"time"
)

var (
	tasks      = make(map[string]time.Time)
	activeTask string
)

func stopTask(taskName string) {
	startTime, exists := tasks[taskName]

	if !exists {
		return
	}
	delete(tasks, taskName)

	activeTask = ""
	_ = time.Since(startTime)
}
