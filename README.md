# Switchboard

A lightweight, CLI framework for Go that makes building command-line applications simple and intuitive. I was originally inspired by another tool I was working on that quickly got difficult to reason about due to multiple commands and flags (skill issue, but I wanted something more readable). As such, I took inspiration from Express.js and the Go http implementation for the nice separation of concerns that their API provides.

[!NOTE]
I am fairly new to Go and its ecosystem. Since starting this project I have discovered [Cobra](https://github.com/spf13/cobra) (a real project). If you are building a real CLI my recommendation is that you should use Cobra because it is well supported and feature rich. I will probably keep developing this because it is a fun learning experience.

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
