package switchboard

import (
	"fmt"
	"os"
	"strings"
)

type FlagResult struct {
	Value interface{}
	Error error
}

type Flag struct {
	Short       string
	Long        string
	Description string
	Required    bool
	processor   func(string, interface{}) FlagResult
}

type Command struct {
	Name  string
	Flags map[string]*Flag // Key is the long name
	order []string
	runFn func()
}

type CLI struct {
	commands map[string]*Command
}

func New() *CLI {
	return &CLI{
		commands: make(map[string]*Command),
	}
}

func (c *CLI) Command(name string, fn func(*Command)) {
	cmd := &Command{
		Name:  name,
		Flags: make(map[string]*Flag),
		order: make([]string, 0),
	}
	fn(cmd)
	c.commands[name] = cmd
}

func (c *Command) Flag(short, long, description string, required bool, processor func(string, interface{}) FlagResult) {
	flag := &Flag{
		Short:       short,
		Long:        long,
		Description: description,
		Required:    required,
		processor:   processor,
	}
	c.Flags[long] = flag
	c.order = append(c.order, long)
}

func (c *Command) Run(fn func()) {
	c.runFn = fn
}

func (c *CLI) Run() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	if cmd, ok := c.commands[args[0]]; ok {
		flagValues := make(map[string]string)
		requiredFlags := make(map[string]bool)

		// Create short name lookup
		shortToLong := make(map[string]string)
		for longName, flag := range cmd.Flags {
			if flag.Short != "" {
				shortToLong[flag.Short] = longName
			}
			if flag.Required {
				requiredFlags[longName] = false
			}
		}

		// Parse args
		for i := 1; i < len(args); i++ {
			arg := args[i]
			if strings.HasPrefix(arg, "-") {
				var flagName string
				if strings.HasPrefix(arg, "--") {
					flagName = strings.TrimPrefix(arg, "--")
				} else {
					shortName := strings.TrimPrefix(arg, "-")
					if longName, exists := shortToLong[shortName]; exists {
						flagName = longName
					}
				}

				if flag, exists := cmd.Flags[flagName]; exists {
					if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
						flagValues[flagName] = args[i+1]
						if flag.Required {
							requiredFlags[flagName] = true
						}
						i++
					} else {
						flagValues[flagName] = ""
						if flag.Required {
							requiredFlags[flagName] = true
						}
					}
				}
			}
		}

		// Check required flags
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

		// Process flags in order
		var lastResult interface{}
		for _, flagName := range cmd.order {
			flag := cmd.Flags[flagName]
			if value, exists := flagValues[flagName]; exists {
				result := flag.processor(value, lastResult)
				if result.Error != nil {
					fmt.Printf("Error processing flag %s: %v\n", flagName, result.Error)
					return
				}
				lastResult = result.Value
			}
		}

		if cmd.runFn != nil {
			cmd.runFn()
		}
	}
}
