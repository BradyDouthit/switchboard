# Switchboard

A lightweight, CLI framework for Go that makes building command-line applications simple and intuitive. I was originally inspired by another tool I was working on that quickly got difficult to reason about due to multiple commands and flags. As such, I took inspiration from Express.js for the nice separation of concerns that their API provides.

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
    "github.com/yourusername/switchboard"
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

### Advanced Usage
For more advanced usage see `/advanced/main.go`