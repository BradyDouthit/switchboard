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

		c.Flag("g", "greeting", "Greeting to use", true,
			func(value string, _ interface{}) switchboard.FlagResult {
				greeting = value
				if greeting == "" {
					greeting = "Hello"
				}
				return switchboard.FlagResult{Value: value}
			})

		c.Flag("n", "name", "First name", true,
			func(value string, _ interface{}) switchboard.FlagResult {
				fullName = value
				return switchboard.FlagResult{Value: value}
			})

		c.Flag("l", "lastname", "Last name", false,
			func(value string, _ interface{}) switchboard.FlagResult {
				fmt.Println("Last Name!", value)
				if value != "" {
					fullName += " " + value
				}
				return switchboard.FlagResult{Value: fullName}
			})

		c.Run(func() {
			fmt.Printf("%s %s\n", greeting, fullName)
		})
	})

	app.Run()
}
