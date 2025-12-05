// Package tui provides a terminal user interface for the task manager application.
package tui

import (
	"strings"
)

func view(m Model) string {
	if len(m.tasks) == 0 {
		return "No tasks yet\n"
	}

	var b strings.Builder
	for i, task := range m.tasks {
		cursor := " "
		if i == m.cursor {
			cursor = "> "
		}
		status := "[ ]"
		if task.Completed {
			status = "[x]"
		}
		b.WriteString(cursor)
		b.WriteString(status)
		b.WriteString(" ")
		b.WriteString(task.Description)
		b.WriteString("\n")
	}
	return b.String()
}
