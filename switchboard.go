package switchboard

import (
	"fmt"
	"strings"
)

// Flag represents a command-line flag with its properties
type Flag struct {
	Short       string
	Long        string
	Description string
	Required    bool
	Value       *string
}

// Command represents a CLI command with its callback function
type Command struct {
	Name        string
	Description string
	Callback    func(args []string, flags map[string]*Flag) error
	flags       map[string]*Flag // maps both short and long names to the same flag
}

// App represents the main CLI application
type App struct {
	commands map[string]*Command
}

// New creates a new CLI application
func New() *App {
	return &App{
		commands: make(map[string]*Command),
	}
}

// Command adds a new command to the application
func (a *App) Command(name string, description string, callback func(args []string, flags map[string]*Flag) error) *Command {
	cmd := &Command{
		Name:        name,
		Description: description,
		Callback:    callback,
		flags:       make(map[string]*Flag),
	}
	a.commands[name] = cmd
	return cmd
}

// Flag adds a new flag to the command
func (c *Command) Flag(short, long, description string, required bool) *Command {
	flag := &Flag{
		Short:       short,
		Long:        long,
		Description: description,
		Required:    required,
		Value:       new(string),
	}

	if short != "" {
		c.flags[short] = flag
	}
	if long != "" {
		c.flags[long] = flag
	}

	return c
}

// Run executes the CLI application with the provided arguments
func (a *App) Run(args []string) error {
	if len(args) < 2 {
		return a.showHelp()
	}

	command := args[1]
	cmd, exists := a.commands[command]
	if !exists {
		return a.showHelp()
	}

	// Parse flags and regular arguments
	parsedArgs, err := parseArgs(args[2:], cmd.flags)
	if err != nil {
		return err
	}

	// Check for required flags
	for _, flag := range cmd.flags {
		if flag.Required && flag.Value == nil {
			return fmt.Errorf("required flag missing: --%s", flag.Long)
		}
	}

	return cmd.Callback(parsedArgs, cmd.flags)
}

// parseArgs separates flags from regular arguments
func parseArgs(args []string, flags map[string]*Flag) ([]string, error) {
	var regularArgs []string

	for i := 0; i < len(args); i++ {
		arg := args[i]

		// Check if argument is a flag
		if strings.HasPrefix(arg, "-") {
			flagName := strings.TrimPrefix(arg, "-")
			if strings.HasPrefix(flagName, "-") {
				flagName = strings.TrimPrefix(flagName, "-")
			}

			flag, exists := flags[flagName]
			if !exists {
				return nil, fmt.Errorf("unknown flag: %s", arg)
			}

			// Check if there's a value following the flag
			if i+1 >= len(args) || strings.HasPrefix(args[i+1], "-") {
				*flag.Value = "true" // Flag present without value
			} else {
				*flag.Value = args[i+1]
				i++ // Skip the next argument since it's the flag value
			}
			continue
		}

		regularArgs = append(regularArgs, arg)
	}

	return regularArgs, nil
}

// showHelp displays available commands and their flags
func (a *App) showHelp() error {
	println("Available commands:")
	for _, cmd := range a.commands {
		println(cmd.Name + "\t" + cmd.Description)
		if len(cmd.flags) > 0 {
			println("  Flags:")
			for _, flag := range cmd.flags {
				shortFlag := ""
				if flag.Short != "" {
					shortFlag = "-" + flag.Short + ", "
				}
				required := ""
				if flag.Required {
					required = " (required)"
				}
				println("    " + shortFlag + "--" + flag.Long + "\t" + flag.Description + required)
			}
		}
		println()
	}
	return nil
}
