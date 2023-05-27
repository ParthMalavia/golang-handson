package cmd

import (
	"fmt"
	"strings"

	"task-manager/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := db.CreateTask(task)
		if err != nil {
			panic(err)
		}
		fmt.Println("Added task with key:", id)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
