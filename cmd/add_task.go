package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func AddTask() {
	// Prompt user for task input
	fmt.Print("Enter task: ")
	reader := bufio.NewReader(os.Stdin)
	task, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Trim whitespace and newline
	task = strings.TrimSpace(task)

	if task == "" {
		fmt.Println("Task cannot be empty")
		return
	}

	file, err := os.OpenFile("tasks.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("Error opening tasks.txt:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(task + "\n"); err != nil {
		fmt.Println("Error writing to tasks.txt:", err)
		return
	}

	fmt.Println("Task added successfully!")
}

var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  "Add a new task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		AddTask()
	},
}
