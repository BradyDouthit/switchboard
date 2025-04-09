# Switchboard

A lightweight CLI framework for Go that makes building command-line applications simple and intuitive. 

## Features
- Simple, intuitive API
- Closure-based state management
- Support for short and long flag names
- Boolean flags
- Required flags
- Subcommands
- Error handling
- Command descriptions

## Installation

```bash
go get github.com/BradyDouthit/switchboard
```

## Basic Usage

Here's a simple example showing basic command creation:

```go
package main

import (
    "fmt"
    "github.com/BradyDouthit/switchboard"
)

func main() {
    app := switchboard.New()
    
    // Simplest possible command
    app.Command("hello", "Say hello to the world", func(c *switchboard.Command) {
        c.Run(func() {
            fmt.Println("Hello, World!")
        })
    })

    app.Run()
}
```

### Usage with flags
Flags can be added to commands with short names, long names, whether or not they are required, and descriptions:

```go
app.Command("greet", "Greet someone", func(c *switchboard.Command) {
    var name string
    var greeting string
    c.Flag("n", "name", "Name to greet", true, func(value string) error {
        name = value
        return nil
    })
    c.Flag("g", "greeting", "Custom greeting", false, func(value string) error {
        greeting = value
        if greeting == "" {
            greeting = "Hello"
        }
        return nil
    })
    c.Run(func() {
        fmt.Printf("%s, %s!\n", greeting, name)
    })
})
```

### Subcommands
You can create nested command structures using subcommands:

```go
app.Command("config", "Manage configuration", func(c *switchboard.Command) {
    c.SubCommand("set", "Set a configuration value", func(sc *switchboard.Command) {
        var key, value string
        sc.Flag("k", "key", "Configuration key", true, func(v string) error {
            key = v
            return nil
        })
        sc.Flag("v", "value", "Configuration value", true, func(v string) error {
            value = v
            return nil
        })
        sc.Run(func() {
            fmt.Printf("Setting %s to %s\n", key, value)
        })
    })
    
    c.SubCommand("get", "Get a configuration value", func(sc *switchboard.Command) {
        var key string
        sc.Flag("k", "key", "Configuration key", true, func(v string) error {
            key = v
            return nil
        })
        sc.Run(func() {
            fmt.Printf("Getting value for %s\n", key)
        })
    })
})
```

This creates a command structure that can be used like:
```bash
myapp config set --key theme --value dark
myapp config get --key theme
```

### Advanced Usage
For more advanced usage see `/advanced/main.go`
