package switchboard

import (
	"fmt"
	"os"
	"strings"
)

// Context holds shared state between commands and flags
type Context struct {
	Values map[string]interface{}
}

// NewContext creates a new context with initialized values
func NewContext() *Context {
	return &Context{
		Values: make(map[string]interface{}),
	}
}

// CLI represents the command line interface
type CLI struct {
	commands map[string]*Command
	context  *Context
}

// Command represents a CLI command
type Command struct {
	Name     string
	Flags    map[string]*Flag
	Context  *Context
	callback func(*Context)
	flagCbs  []func()
}

// Flag represents a command flag
type Flag struct {
	Name        string
	Description string
	Required    bool
	callback    func(string)
}

// New creates a new CLI instance
func New() *CLI {
	return &CLI{
		commands: make(map[string]*Command),
		context:  NewContext(),
	}
}

// Command adds a new command to the CLI
func (c *CLI) Command(name string, handlers ...interface{}) {
	cmd := &Command{
		Name:    name,
		Flags:   make(map[string]*Flag),
		Context: c.context,
		flagCbs: make([]func(), 0),
	}

	switch len(handlers) {
	case 1:
		// Simple case: just a function with no configuration
		if handler, ok := handlers[0].(func()); ok {
			cmd.callback = func(*Context) {
				handler()
			}
		} else if configFn, ok := handlers[0].(func(*Command)); ok {
			configFn(cmd)
			// If no callback was set during configuration, use empty callback
			if cmd.callback == nil {
				cmd.callback = func(*Context) {}
			}
		} else {
			panic("Invalid command handler type")
		}
	case 2:
		// Advanced case: configuration function and callback
		configFn, okConfig := handlers[0].(func(*Command))
		callbackFn, okCallback := handlers[1].(func(*Context))

		if !okConfig || !okCallback {
			panic("Invalid command handler types")
		}

		configFn(cmd)
		cmd.callback = callbackFn
	default:
		panic("Invalid number of handlers")
	}

	c.commands[name] = cmd
}

// Flag adds a new flag to the command
func (c *Command) Flag(name string, description string, required bool, callback func(string)) {
	flag := &Flag{
		Name:        name,
		Description: description,
		Required:    required,
		callback:    callback,
	}
	c.Flags[name] = flag
}

// Run executes the CLI
func (c *CLI) Run() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	if cmd, ok := c.commands[args[0]]; ok {
		// Track which required flags have been set
		requiredFlags := make(map[string]bool)
		for flagName, flag := range cmd.Flags {
			if flag.Required {
				requiredFlags[flagName] = false
			}
		}

		// Parse remaining args for flags
		for i := 1; i < len(args); i++ {
			arg := args[i]
			if strings.HasPrefix(arg, "-") {
				flagName := strings.TrimPrefix(arg, "-")
				if flag, exists := cmd.Flags[flagName]; exists {
					if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
						flag.callback(args[i+1])
						if flag.Required {
							requiredFlags[flagName] = true
						}
						i++
					} else {
						flag.callback("")
						if flag.Required {
							requiredFlags[flagName] = true
						}
					}
				}
			}
		}

		// Check if all required flags were provided
		missingFlags := []string{}
		for flagName, provided := range requiredFlags {
			if !provided {
				missingFlags = append(missingFlags, flagName)
			}
		}

		if len(missingFlags) > 0 {
			fmt.Printf("Error: Missing required flags: %v\n", missingFlags)
			return
		}

		// Execute command callback with context
		cmd.callback(c.context)
	}
}
