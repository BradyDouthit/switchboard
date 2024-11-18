package main

import (
	"fmt"
	"os"

	"github.com/BradyDouthit/switchboard"
)

func main() {
	app := switchboard.New()

	// Variables to store flag values
	var greeting string = "Hello"
	var language string = "en"

	// Register a command with flags
	app.Command("greet", "Greets the user with the provided name", func(args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("please provide a name")
		}

		switch language {
		case "es":
			greeting = "Hola"
		case "fr":
			greeting = "Bonjour"
		}

		fmt.Printf("%s, %s!\n", greeting, args[0])
		return nil
	}).
		Flag("F", "formal", "Use formal greeting", func(value string) error {
			if value == "true" {
				greeting = "Good day"
			}
			return nil
		}).
		Flag("l", "lang", "Greeting language (en, es, fr)", func(value string) error {
			switch value {
			case "en", "es", "fr":
				language = value
				return nil
			default:
				return fmt.Errorf("unsupported language: %s", value)
			}
		})

	// Run the application
	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
