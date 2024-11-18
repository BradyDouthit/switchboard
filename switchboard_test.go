package switchboard

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

// Helper function to capture stdout
func captureOutput(fn func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String()
}

func TestBasicCommand(t *testing.T) {
	output := captureOutput(func() {
		app := New()
		app.Command("hello", "Basic hello command", func(c *Command) {
			c.Run(func() {
				fmt.Println("Hello, World!")
			})
		})
		os.Args = []string{"app", "hello"}
		app.Run()
	})

	if strings.TrimSpace(output) != "Hello, World!" {
		t.Errorf("Expected 'Hello, World!', got '%s'", strings.TrimSpace(output))
	}
}

func TestRequiredFlag(t *testing.T) {
	output := captureOutput(func() {
		app := New()
		app.Command("greet", "Greet command", func(c *Command) {
			var name string
			c.Flag("n", "name", "Name to greet", true, func(value string) error {
				name = value
				return nil
			})
			c.Run(func() {
				fmt.Printf("Hello, %s!", name)
			})
		})
		os.Args = []string{"app", "greet", "--name", "John"}
		app.Run()
	})

	if strings.TrimSpace(output) != "Hello, John!" {
		t.Errorf("Expected 'Hello, John!', got '%s'", strings.TrimSpace(output))
	}
}

func TestMissingRequiredFlag(t *testing.T) {
	output := captureOutput(func() {
		app := New()
		app.Command("greet", "Greet command", func(c *Command) {
			c.Flag("n", "name", "Name to greet", true, func(value string) error {
				return nil
			})
			c.Run(func() {
				t.Error("Run function should not be called")
			})
		})
		os.Args = []string{"app", "greet"}
		app.Run()
	})

	if !strings.Contains(output, "Missing required flags") {
		t.Errorf("Expected missing required flag error, got '%s'", output)
	}
}

func TestBooleanFlag(t *testing.T) {
	output := captureOutput(func() {
		app := New()
		app.Command("verbose", "Verbose command", func(c *Command) {
			var verbose bool
			c.BoolFlag("v", "verbose", "Enable verbose mode", func(value bool) error {
				verbose = value
				return nil
			})
			c.Run(func() {
				if verbose {
					fmt.Print("Verbose mode enabled")
				}
			})
		})
		os.Args = []string{"app", "verbose", "--verbose"}
		app.Run()
	})

	if strings.TrimSpace(output) != "Verbose mode enabled" {
		t.Errorf("Expected 'Verbose mode enabled', got '%s'", strings.TrimSpace(output))
	}
}

func TestSubCommand(t *testing.T) {
	output := captureOutput(func() {
		app := New()
		app.Command("server", "Server commands", func(c *Command) {
			c.SubCommand("start", "Start server", func(sc *Command) {
				var port string
				sc.Flag("p", "port", "Port number", true, func(value string) error {
					port = value
					return nil
				})
				sc.Run(func() {
					fmt.Printf("Server starting on port %s", port)
				})
			})
		})
		os.Args = []string{"app", "server", "start", "--port", "8080"}
		app.Run()
	})

	if strings.TrimSpace(output) != "Server starting on port 8080" {
		t.Errorf("Expected 'Server starting on port 8080', got '%s'", strings.TrimSpace(output))
	}
}

func TestPositionalArgs(t *testing.T) {
	output := captureOutput(func() {
		app := New()
		app.Command("echo", "Echo command", func(c *Command) {
			c.Run(func(args []string) {
				fmt.Print(strings.Join(args, " "))
			})
		})
		os.Args = []string{"app", "echo", "hello", "world"}
		app.Run()
	})

	if strings.TrimSpace(output) != "hello world" {
		t.Errorf("Expected 'hello world', got '%s'", strings.TrimSpace(output))
	}
}

func TestMixedFlagsAndArgs(t *testing.T) {
	output := captureOutput(func() {
		app := New()
		app.Command("copy", "Copy command", func(c *Command) {
			var verbose bool
			c.BoolFlag("v", "verbose", "Verbose output", func(value bool) error {
				verbose = value
				return nil
			})
			c.Run(func(args []string) {
				if verbose {
					fmt.Printf("Copying %s to %s", args[0], args[1])
				}
			})
		})
		os.Args = []string{"app", "copy", "source.txt", "dest.txt", "--verbose"}
		app.Run()
	})

	if strings.TrimSpace(output) != "Copying source.txt to dest.txt" {
		t.Errorf("Expected 'Copying source.txt to dest.txt', got '%s'", strings.TrimSpace(output))
	}
}

func TestFlagError(t *testing.T) {
	output := captureOutput(func() {
		app := New()
		app.Command("port", "Port command", func(c *Command) {
			c.Flag("p", "port", "Port number", true, func(value string) error {
				return fmt.Errorf("invalid port")
			})
			c.Run(func() {
				t.Error("Run function should not be called")
			})
		})
		os.Args = []string{"app", "port", "--port", "invalid"}
		app.Run()
	})

	if !strings.Contains(output, "invalid port") {
		t.Errorf("Expected 'invalid port' error, got '%s'", output)
	}
}

func TestShortFlags(t *testing.T) {
	output := captureOutput(func() {
		app := New()
		app.Command("greet", "Greet command", func(c *Command) {
			var name string
			c.Flag("n", "name", "Name to greet", true, func(value string) error {
				name = value
				return nil
			})
			c.Run(func() {
				fmt.Printf("Hello, %s!", name)
			})
		})
		os.Args = []string{"app", "greet", "-n", "John"}
		app.Run()
	})

	if strings.TrimSpace(output) != "Hello, John!" {
		t.Errorf("Expected 'Hello, John!', got '%s'", strings.TrimSpace(output))
	}
}
