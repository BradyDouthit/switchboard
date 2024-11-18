package main

import (
	"fmt"

	"github.com/BradyDouthit/switchboard"
)

func main() {
	app := switchboard.New()

	app.Command("copy", "Copy files", func(c *switchboard.Command) {
		var verbose bool

		c.BoolFlag("v", "verbose", "Show verbose output", func(value bool) error {
			verbose = value
			return nil
		})

		c.Run(func(args []string) {
			if len(args) < 2 {
				fmt.Println("Error: copy requires source and destination arguments")
				return
			}
			source := args[0]
			dest := args[1]

			if verbose {
				fmt.Printf("Copying %s to %s\n", source, dest)
			} else {
				fmt.Printf("Copying files...\n")
			}
		})
	})

	app.Run()
}
