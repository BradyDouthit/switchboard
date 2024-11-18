package main

import (
	"fmt"

	"github.com/BradyDouthit/switchboard"
)

func main() {
	app := switchboard.New()

	app.Command("greet", func(c *switchboard.Command) {
		var greeting string
		var fullName string

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
				fmt.Printf("%s %s\n", greeting, fullName)
				return nil
			})
	})

	app.Run()
}
