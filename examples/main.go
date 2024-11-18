package main

import (
	"fmt"

	"github.com/BradyDouthit/switchboard"
)

func main() {
	app := switchboard.NewCLI()

	app.Command("greet", func(c *switchboard.Command) {
		c.Flag("name", "First name", func(value string) {
			c.context.Values["firstName"] = value
		})
		c.Flag("L", "Last name", func(value string) {
			c.context.Values["lastName"] = value
		})
	}, func(ctx *switchboard.Context) {
		firstName, hasFirst := ctx.Values["firstName"].(string)
		if !hasFirst {
			fmt.Println("Please provide a name using -name flag")
			return
		}

		lastName, hasLast := ctx.Values["lastName"].(string)
		if hasLast {
			fmt.Printf("Hello %s %s\n", firstName, lastName)
		} else {
			fmt.Printf("Hello %s\n", firstName)
		}
	})

	app.Run()
}
