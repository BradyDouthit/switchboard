package main

import (
	"fmt"
	"os"

	"github.com/BradyDouthit/switchboard"
)

func main() {
	app := switchboard.New()

	// Register a command with flags
	app.Command("greet", "Greets the user with the provided name", func(args []string, flags map[string]*switchboard.Flag) error {
		fmt.Println()
		return nil
	}).Flag("H", "hello", "say hello", false)

	// Run the application
	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
