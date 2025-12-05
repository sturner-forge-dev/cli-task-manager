// Package cmd implements the command-line interface for managing tasks.
package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli-task-manager",
	Short: "Task Manager CLI Application",
	Long:  "Task Manager is a simple CLI tool for managing your tasks",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(addTaskCmd)
	rootCmd.AddCommand(listTasksCmd)
	rootCmd.AddCommand(tuiCmd)
}
