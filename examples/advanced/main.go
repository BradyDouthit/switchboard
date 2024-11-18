package main

import (
	"fmt"

	"github.com/BradyDouthit/switchboard"
)

func main() {
	app := switchboard.New()

	app.Command("server", "Server management commands", func(c *switchboard.Command) {
		// Start subcommand
		c.SubCommand("start", "Start the server", func(sc *switchboard.Command) {
			var port string
			var debug bool
			var configFile string

			sc.Flag("p", "port", "Port to listen on", false, func(value string) error {
				port = value
				if port == "" {
					port = "8080" // default value
				}
				return nil
			})

			sc.BoolFlag("d", "debug", "Enable debug mode", func(value bool) error {
				debug = value
				return nil
			})

			sc.Flag("c", "config", "Path to config file", false, func(value string) error {
				configFile = value
				return nil
			})

			sc.Run(func() {
				fmt.Printf("Starting server on port %s\n", port)
				if debug {
					fmt.Println("Debug mode enabled")
				}
				if configFile != "" {
					fmt.Printf("Using config file: %s\n", configFile)
				}
			})
		})

		// Status subcommand
		c.SubCommand("status", "Check server status", func(sc *switchboard.Command) {
			var format string
			var verbose bool

			sc.Flag("f", "format", "Output format (json|text)", false, func(value string) error {
				format = value
				if format == "" {
					format = "text"
				}
				if format != "json" && format != "text" {
					return fmt.Errorf("invalid format: %s", format)
				}
				return nil
			})

			sc.BoolFlag("v", "verbose", "Show detailed status", func(value bool) error {
				verbose = value
				return nil
			})

			sc.Run(func() {
				if format == "json" {
					fmt.Printf(`{"status": "running", "verbose": %v}`, verbose)
				} else {
					fmt.Printf("Server Status: Running\n")
					if verbose {
						fmt.Printf("Uptime: 2h 30m\n")
						fmt.Printf("Active connections: 42\n")
						fmt.Printf("Memory usage: 128MB\n")
					}
				}
			})
		})

		// Stop subcommand
		c.SubCommand("stop", "Stop the server", func(sc *switchboard.Command) {
			var force bool
			var timeout string

			sc.BoolFlag("f", "force", "Force immediate shutdown", func(value bool) error {
				force = value
				return nil
			})

			sc.Flag("t", "timeout", "Shutdown timeout in seconds", false, func(value string) error {
				timeout = value
				if timeout == "" {
					timeout = "30"
				}
				return nil
			})

			sc.Run(func() {
				if force {
					fmt.Println("Force stopping server...")
				} else {
					fmt.Printf("Gracefully stopping server (timeout: %ss)...\n", timeout)
				}
			})
		})
	})

	app.Run()
}
