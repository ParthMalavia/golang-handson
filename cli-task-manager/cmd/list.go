package cmd

import (
	"fmt"

	"task-manager/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lost down all the remaining tasks from the list.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks. Take a break.")
			return
		}
		fmt.Println("You have following tasks:")
		for i, t := range tasks {
			fmt.Printf("%d: %s \n", i+1, t.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
