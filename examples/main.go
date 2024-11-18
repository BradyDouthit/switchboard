package main

import (
	"fmt"

	"github.com/BradyDouthit/switchboard"
)

func handleGreet(ctx *switchboard.Context) {
	firstName, hasFirst := ctx.Values["firstName"].(string)

	if !hasFirst {
		fmt.Println("Please provide a name using -N (Name) flag")
		return
	}

	lastName, hasLast := ctx.Values["lastName"].(string)
	if hasLast {
		fmt.Printf("Hello %s %s\n", firstName, lastName)
	} else {
		fmt.Printf("Hello %s\n", firstName)
	}
}

func main() {
	app := switchboard.NewCLI()

	app.Command("greet", func(c *switchboard.Command) {
		c.Flag("N", "name", func(value string) {
			c.Context.Values["firstName"] = value
		})
		c.Flag("L", "lastname", func(value string) {
			c.Context.Values["lastName"] = value
		})
	}, handleGreet)

	app.Run()
}
