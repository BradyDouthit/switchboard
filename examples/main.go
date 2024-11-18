package main

import (
	"fmt"

	"github.com/BradyDouthit/switchboard"
)

func main() {
	app := switchboard.New()

	// Simple usage
	app.Command("hello", func() {
		fmt.Println("Hello, World!")
	})

	// Advanced usage with flags and context
	app.Command("greet",
		func(c *switchboard.Command) {
			c.Flag("N", "name", true, func(value string) {
				c.Context.Values["firstName"] = value
			})
			c.Flag("L", "lastname", false, func(value string) {
				c.Context.Values["lastName"] = value
			})
		},
		func(ctx *switchboard.Context) {
			firstName := ctx.Values["firstName"].(string)
			lastName, hasLast := ctx.Values["lastName"].(string)

			if hasLast {
				fmt.Printf("Hello %s %s\n", firstName, lastName)
			} else {
				fmt.Printf("Hello %s\n", firstName)
			}
		},
	)

	app.Run()
}
