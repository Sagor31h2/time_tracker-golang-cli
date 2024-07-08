package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "timetracker",
	Short: "A simple time tracker app",
	Long:  `A simple CLI time tracker app built with Go and Cobra.`,
}

func init() {
	RootCmd.AddCommand(startCmd)
	RootCmd.AddCommand(stopCmd)
	RootCmd.AddCommand(statusCmd)
}
