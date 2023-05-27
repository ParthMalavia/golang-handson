package main

import (
	"fmt"
	"os"
	"path/filepath"

	"task-manager/cmd"
	"task-manager/db"
)

func main() {

	// Get current working directory to save DB.
	cwd, _ := os.Getwd()
	dbPath := filepath.Join(cwd, "test.db")

	ErrorHandler(db.Init(dbPath))
	ErrorHandler(cmd.RootCmd.Execute())
}

func ErrorHandler(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
