package tasks

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const defaultFile = "tasks.txt"

func LoadTasks(path string) ([]Task, error) {
	if path == "" {
		path = defaultFile
	}

	f, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return []Task{}, nil
	} else if err != nil {
		return nil, err
	}
	defer f.Close()

	tasks := make([]Task, 0)
	scanner := bufio.NewScanner(f)
	id := 1
	for scanner.Scan() {
		desc := strings.TrimSpace(scanner.Text())
		if desc == "" {
			continue
		}
		task := NewTask(id, desc)
		tasks = append(tasks, task)
		id++
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func SaveTasks(path string, entries []Task) error {
	if path == "" {
		path = defaultFile
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	for _, task := range entries {
		if _, err := fmt.Fprintln(writer, strings.TrimSpace(task.Description)); err != nil {
			return err
		}
	}
	// TODO: persist completed status

	return writer.Flush()
}
