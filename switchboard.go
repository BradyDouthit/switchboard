package switchboard

import (
	"fmt"
	"os"
	"strings"
)

type Flag struct {
	Short       string
	Long        string
	Description string
	Required    bool
	IsBoolean   bool
	processor   func(string) error
}

type Command struct {
	Name        string
	Description string
	Flags       map[string]*Flag
	order       []string
	subCommands map[string]*Command
	runFn       func([]string)
	simpleFn    func()
}

type CLI struct {
	commands map[string]*Command
}

func New() *CLI {
	return &CLI{
		commands: make(map[string]*Command),
	}
}

func (c *CLI) Command(name string, description string, fn func(*Command)) {
	cmd := &Command{
		Name:        name,
		Description: description,
		Flags:       make(map[string]*Flag),
		order:       make([]string, 0),
		subCommands: make(map[string]*Command),
	}
	fn(cmd)
	c.commands[name] = cmd
}

func (c *Command) Flag(short, long, description string, required bool, processor func(string) error) {
	flag := &Flag{
		Short:       short,
		Long:        long,
		Description: description,
		Required:    required,
		IsBoolean:   false,
		processor:   processor,
	}
	c.Flags[long] = flag
	c.order = append(c.order, long)
}

func (c *Command) BoolFlag(short, long, description string, processor func(bool) error) {
	flag := &Flag{
		Short:       short,
		Long:        long,
		Description: description,
		Required:    false,
		IsBoolean:   true,
		processor: func(value string) error {
			return processor(value != "")
		},
	}
	c.Flags[long] = flag
	c.order = append(c.order, long)
}

func (c *Command) SubCommand(name string, description string, fn func(*Command)) {
	subcmd := &Command{
		Name:        name,
		Description: description,
		Flags:       make(map[string]*Flag),
		order:       make([]string, 0),
		subCommands: make(map[string]*Command),
	}
	fn(subcmd)
	c.subCommands[name] = subcmd
}

func (c *CLI) Run() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	if cmd, ok := c.commands[args[0]]; ok {
		remaining := args[1:]

		if len(remaining) > 0 {
			if subcmd, ok := cmd.subCommands[remaining[0]]; ok {
				processCommand(subcmd, remaining[1:])
				return
			}
		}

		processCommand(cmd, remaining)
	}
}

func processCommand(cmd *Command, args []string) {
	flagValues := make(map[string]string)
	requiredFlags := make(map[string]bool)
	positionalArgs := []string{}

	shortToLong := make(map[string]string)
	for longName, flag := range cmd.Flags {
		if flag.Short != "" {
			shortToLong[flag.Short] = longName
		}
		if flag.Required {
			requiredFlags[longName] = false
		}
	}

	i := 0
	for i < len(args) {
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
				if flag.IsBoolean {
					flagValues[flagName] = "true"
					if flag.Required {
						requiredFlags[flagName] = true
					}
				} else if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
					flagValues[flagName] = args[i+1]
					if flag.Required {
						requiredFlags[flagName] = true
					}
					i++
				}
			}
			i++
		} else {
			positionalArgs = append(positionalArgs, arg)
			i++
		}
	}

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

	for _, flagName := range cmd.order {
		flag := cmd.Flags[flagName]
		if value, exists := flagValues[flagName]; exists {
			if err := flag.processor(value); err != nil {
				fmt.Printf("Error processing flag %s: %v\n", flagName, err)
				return
			}
		}
	}

	if cmd.runFn != nil {
		cmd.runFn(positionalArgs)
	} else if cmd.simpleFn != nil {
		cmd.simpleFn()
	}
}

func (c *Command) Run(fn interface{}) {
	switch f := fn.(type) {
	case func():
		c.simpleFn = f
	case func([]string):
		c.runFn = f
	default:
		panic("Invalid run function signature")
	}
}
