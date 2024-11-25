package main

import (
	"fmt"

	"github.com/BradyDouthit/switchboard"
)

func main() {
	app := switchboard.New()

	app.Command("greet", "Greet a person with optional customizations", func(c *switchboard.Command) {
		// 1. Define any state to be shared for the whole commmand
		var greeting string
		var fullName string

		greetingFlag := switchboard.Flag{
			Short:       "g",
			Long:        "greeting",
			Description: "Greeting to use",
			Required:    false,
		}

		// 2. Define functionality for all needed flags
		c.Flag(&greetingFlag,
			func(value string) error {
				greeting = value
				if greeting == "" {
					greeting = "Hello"
				}
				return nil
			})

		nameFlag := switchboard.Flag{
			Short:       "n",
			Long:        "name",
			Description: "First name",
			Required:    false,
		}

		c.Flag(&nameFlag,
			func(value string) error {
				fullName = value
				return nil
			})

		lastnameFlag := switchboard.Flag{
			Short:       "l",
			Long:        "lastname",
			Description: "Last name",
			Required:    false,
		}

		c.Flag(&lastnameFlag,
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
