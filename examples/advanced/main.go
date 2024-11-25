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

			portFlag := switchboard.Flag{
				Short:       "p",
				Long:        "port",
				Description: "Port to listen on",
				Required:    false,
			}

			sc.Flag(&portFlag, func(value string) error {
				port = value
				if port == "" {
					port = "8080" // default value
				}
				return nil
			})

			debugFlag := switchboard.Flag{
				Short:       "d",
				Long:        "debug",
				Description: "Enable debug mode",
				Required:    false,
			}

			sc.BoolFlag(&debugFlag, func(value bool) error {
				debug = value
				return nil
			})

			configFlag := switchboard.Flag{
				Short:       "c",
				Long:        "config",
				Description: "Path to config file",
				Required:    false,
			}

			sc.Flag(&configFlag, func(value string) error {
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

			formatFlag := switchboard.Flag{
				Short:       "f",
				Long:        "format",
				Description: "Output format (json|text)",
				Required:    false,
			}

			sc.Flag(&formatFlag, func(value string) error {
				format = value
				if format == "" {
					format = "text"
				}
				if format != "json" && format != "text" {
					return fmt.Errorf("invalid format: %s", format)
				}
				return nil
			})

			verboseFlag := switchboard.Flag{
				Short:       "v",
				Long:        "verbose",
				Description: "Show detailed status",
				Required:    false,
			}

			sc.BoolFlag(&verboseFlag, func(value bool) error {
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

			forceFlag := switchboard.Flag{
				Short:       "f",
				Long:        "force",
				Description: "Force immediate shutdown",
				Required:    false,
			}

			sc.BoolFlag(&forceFlag, func(value bool) error {
				force = value
				return nil
			})

			timeoutFlag := switchboard.Flag{
				Short:       "t",
				Long:        "timeout",
				Description: "Shutdown timeout in seconds",
				Required:    false,
			}

			sc.Flag(&timeoutFlag, func(value string) error {
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
