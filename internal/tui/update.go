package tui

import (
	"fmt"

	"cli-task-manager/internal/tasks"

	tea "github.com/charmbracelet/bubbletea"
)

func update(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k", "h":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j", "l":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}
		case " ":
			m.tasks[m.cursor].Toggle()
			err := tasks.SaveTasks("", m.tasks)
			if err != nil {
				m.status = fmt.Sprintf("Error saving tasks: %v", err)
			}
		}
	}
	return m, nil
}
