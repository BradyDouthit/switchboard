package main

import (
	"fmt"

	"github.com/BradyDouthit/switchboard"
)

func main() {
	app := switchboard.New()

	app.Command("greet", func(c *switchboard.Command) {
		// 1. Define any state to be shared for the whole commmand
		var greeting string
		var fullName string

		// 2. Define functionality for all needed flags
		c.Flag("g", "greeting", "Greeting to use", false,
			func(value string) error {
				greeting = value
				if greeting == "" {
					greeting = "Hello"
				}
				return nil
			})

		c.Flag("n", "name", "First name", true,
			func(value string) error {
				fullName = value
				return nil
			})

		c.Flag("l", "lastname", "Last name", false,
			func(value string) error {
				if value != "" {
					fullName += " " + value
				}
				return nil
			})

		// 3. Run the full command with the state modified by each flag
		c.Run(func() {
			fmt.Printf("%s %s\n", greeting, fullName)
		})
	})

	app.Run()
}
