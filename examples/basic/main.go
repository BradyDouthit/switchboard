package main

import (
	"fmt"

	"github.com/BradyDouthit/switchboard"
)

func main() {
	app := switchboard.New()

	// Simplest possible command with no flags
	app.Command("hello", "Say hello to the world", func(c *switchboard.Command) {
		c.Run(func() {
			fmt.Println("Hello, World!")
		})
	})

	// Simple command with one required flag
	app.Command("echo", "Echo a message", func(c *switchboard.Command) {
		var message string

		flag := switchboard.Flag{
			Short:       "m",
			Long:        "message",
			Description: "Message to echo",
			Required:    true,
		}

		c.Flag(&flag, func(value string) error {
			message = value
			return nil
		})

		c.Run(func() {
			fmt.Println(message)
		})
	})

	app.Run()
}
