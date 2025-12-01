// Package tasks provides a simple task management system.
package tasks

import (
	"fmt"
	"strings"
)

type Task struct {
	ID          int
	Description string
	Completed   bool
}

func NewTask(id int, description string) Task {
	if description == "" {
		fmt.Println("Error: description cannot be empty")
	}
	// trim whitespace from description
	description = strings.TrimSpace(description)

	return Task{
		ID:          id,
		Description: description,
		Completed:   false,
	}
}

func (t *Task) Toggle() {
	t.Completed = !t.Completed
}

func (t *Task) MarkCompleted() {
	t.Completed = true
}

func (t *Task) MarkIncomplete() {
	t.Completed = false
}

func (t Task) String() string {
	if t.Completed {
		return fmt.Sprintf("[x] %d: %s", t.ID, t.Description)
	}

	return fmt.Sprintf("[ ] %d: %s", t.ID, t.Description)
}
