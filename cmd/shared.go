package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Task struct {
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time"`
}

var (
	tasks      = make(map[string]Task)
	activeTask string
	configFile string
)

func init() {
	var err error
	configFile, err = getConfigFilePath()
	if err != nil {
		fmt.Printf("Error determining config file path: %v\n", err)
		os.Exit(1)
	}

	loadTasks()
}

func getConfigFilePath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	appConfigDir := filepath.Join(configDir, "timetracker")
	err = os.MkdirAll(appConfigDir, 0755)
	if err != nil {
		return "", err
	}
	return filepath.Join(appConfigDir, "tasks.json"), nil
}

func loadTasks() {
	data, err := os.ReadFile(configFile)
	if os.IsNotExist(err) {
		return // It's okay if the file does not exist
	} else if err != nil {
		fmt.Printf("Error reading tasks file: %v\n", err)
		return
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Printf("Error unmarshalling tasks: %v\n", err)
	}
	for name, task := range tasks {
		if !task.StartTime.IsZero() {
			activeTask = name
			break
		}
	}
}

func saveTasks() {
	data, err := json.Marshal(tasks)
	if err != nil {
		fmt.Printf("Error marshalling tasks: %v\n", err)
		return
	}
	err = os.WriteFile(configFile, data, 0644)
	if err != nil {
		fmt.Printf("Error writing tasks file: %v\n", err)
	}
}
