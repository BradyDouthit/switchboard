# Switchboard

A lightweight, CLI framework for Go that makes building command-line applications simple and intuitive. I was originally inspired by another tool I was working on that quickly got difficult to reason about due to multiple commands and flags. As such, I took inspiration from Express.js for the nice separation of concerns that their API provides.

## Installation

```bash
go get github.com/BradyDouthit/switchboard
```

## TODO: Quick Example

```go
```

## Usage

```bash
$ app hello
Hello, World!

$ app greet Alice
Hello, Alice!

$ app greet Alice --uppercase
Hello, ALICE!

$ app
Available commands:
  hello    Says hello
  greet    Greets someone
    -u, --uppercase    Convert to uppercase
```

## Features

- Simple, chainable API
- Short (-u) and long (--uppercase) flags
- Automatic help text generation
- Built-in error handling

## License

MIT