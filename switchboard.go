package switchboard

import (
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
	callback    func(string)
}

// NewCLI creates a new CLI instance
func NewCLI() *CLI {
	return &CLI{
		commands: make(map[string]*Command),
		context:  NewContext(),
	}
}

// Command adds a new command to the CLI
func (c *CLI) Command(name string, fn func(*Command), callback func(*Context)) {
	cmd := &Command{
		Name:     name,
		Flags:    make(map[string]*Flag),
		callback: callback,
		flagCbs:  make([]func(), 0),
		Context:  c.context,
	}
	fn(cmd)
	c.commands[name] = cmd
}

// Flag adds a new flag to the command
func (c *Command) Flag(name string, description string, callback func(string)) {
	flag := &Flag{
		Name:        name,
		Description: description,
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
		// Parse remaining args for flags
		for i := 1; i < len(args); i++ {
			arg := args[i]
			if strings.HasPrefix(arg, "-") {
				flagName := strings.TrimPrefix(arg, "-")
				if flag, exists := cmd.Flags[flagName]; exists {
					if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
						flag.callback(args[i+1])
						i++
					} else {
						flag.callback("")
					}
				}
			}
		}

		// Execute command callback with context
		cmd.callback(c.context)
	}
}
