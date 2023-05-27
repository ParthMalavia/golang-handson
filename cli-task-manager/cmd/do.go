package cmd

import (
	"fmt"
	"strconv"
	"task-manager/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			return
		}
		for _, arg := range args {
			if id, err := strconv.Atoi(arg); err != nil {
				fmt.Println("Failed to parse %s argument.", arg)
			} else {
				if id < 1 || id > len(tasks) {
					fmt.Println("Invalid task number:", id)
					continue
				}
				task := tasks[id-1]
				err := db.DeleteTask(task.Key)
				if err != nil {
					fmt.Printf("Failed to delete '%d' task. Error: %s \n", id, err)
				} else {
					fmt.Printf("Deleted '%d' task.\n", id)
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
