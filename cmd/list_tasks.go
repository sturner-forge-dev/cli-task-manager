package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func ListTasks() {
	file, err := os.Open("tasks.txt")
	if err != nil {
		fmt.Println("Error opening tasks.txt:", err)
		return
	}
	defer file.Close()

	fmt.Println("Your Tasks:")
	fmt.Println("-----------")
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		task := fileScanner.Text()
		fmt.Println("-", task)
	}
	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error reading tasks.txt:", err)
	}
}

var listTasksCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "List all tasks in your task list",
	Run: func(cmd *cobra.Command, args []string) {
		ListTasks()
	},
}
