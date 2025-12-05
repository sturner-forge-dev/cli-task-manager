package cmd

import (
	"cli-task-manager/internal/tasks"
	"cli-task-manager/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Run the tui",
	RunE: func(cmd *cobra.Command, args []string) error {
		tasksList, err := tasks.LoadTasks("")
		if err != nil {
			return err
		}
		model := tui.NewModel(tasksList)

		_, err = tea.NewProgram(model).Run()
		return err
	},
}
