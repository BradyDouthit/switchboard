package main

import (
	"fmt"

	"github.com/BradyDouthit/switchboard"
)

func main() {
	app := switchboard.New()

	app.Command("greet", func(c *switchboard.Command) {
		c.Flag("name", "First name", true,
			func(value string, _ interface{}) switchboard.FlagResult {
				return switchboard.FlagResult{
					Value: value,
				}
			})

		c.Flag("L", "Last name", false,
			func(value string, prevResult interface{}) switchboard.FlagResult {
				firstName := prevResult.(string)
				fullName := fmt.Sprintf("%s %s", firstName, value)
				fmt.Printf("Hello %s\n", fullName)
				return switchboard.FlagResult{Value: fullName}
			})
	})

	app.Run()
}
