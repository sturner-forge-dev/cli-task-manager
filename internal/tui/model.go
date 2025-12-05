package tui

import (
	"cli-task-manager/internal/tasks"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	tasks  []tasks.Task
	cursor int
	status string
}

// NewModel returns a TUI model populated with tasks and a cursor at the first entry
func NewModel(taskList []tasks.Task) Model {
	return Model{
		append([]tasks.Task(nil), taskList...),
		0,
		"",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return update(msg, m)
}

func (m Model) View() string {
	return view(m)
}
