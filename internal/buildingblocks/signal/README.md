# Understanding Signals in Golang

In computing, a signal is a software interrupt delivered to a process.  
The operating system uses signals to report exceptional situations to an executing program.  
Some signals report errors such as references to non-existent memory.  
Others report asynchronous events, such as disconnection of a terminal line.

In Go, the `os/signal` package provides a mechanism to receive signals.  
This can be used to catch and respond to conditions like a user wanting to interrupt a running program (using `Ctrl+C`,
for instance).

<img src="https://miro.medium.com/v2/resize:fit:1400/format:webp/1*WRUe5p-pzLyCpU6hLvNVUg.png" alt="drawing" height="300"/>

## Table of Contents

1. [Introduction to Signals](#introduction)
2. [Handling Signals in Go](#handling-signals-in-go)
3. [Common Pitfalls and Issues](#common-pitfalls-and-issues)
4. [Best Practices](#best-practices)
5. [Resources](#resources)

## Introduction

Signals are used in systems programming to notify running processes of certain events.  
For instance, the `SIGINT` signal is sent when the user presses `Ctrl+C` in the terminal, and `SIGTERM` is sent to gracefully shut down a process.

## Handling Signals in Go

Here's a simplified example of how to handle signals in Go:

```go
package main

import (
    "fmt"
    "os"
    "os/signal"
)

func main() {
    // Create a channel to receive OS signals
    sigs := make(chan os.Signal, 1)
    
    // Register the channel to receive specified signals
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    
    // Wait for a signal
    sig := <-sigs
    fmt.Println()
    fmt.Println(sig)
}
```

In this example, we first create a channel `sigs` to receive signals.  
We then call `signal.Notify`, registering `sigs` to receive `SIGINT` and `SIGTERM` signals.  
Finally, we wait for a signal by reading from `sigs`.

You can send a signal to a running process from the terminal.  
For instance, to send `SIGINT` to a process, you can use `Ctrl+C`.  
To send a `SIGTERM` signal, you can use the `kill` command followed by the process ID:

```bash
    kill -TERM <pid>
```

## Common Pitfalls and Issues

- **Ignoring Signals**:
    - Ignoring signals can cause your program to behave unpredictably or miss important system events.

- **Blocking on Signal Channels**:
    - If the signal channel is full or not being read from, signals can be missed.

## Best Practices

- **Handle Signals Gracefully**:
    - Use signal handling to clean up resources and shut down your program gracefully.

- **Avoid Long-Running Handlers**:
    - Keep signal handlers short to prevent blocking other processing.

- **Use Buffered Channels**:
    - Use a buffered channel when registering to receive signals to ensure that no signals are missed.

## Resources

- [Go by Example: Signals](https://gobyexample.com/signals)
- [os/signal package documentation](https://pkg.go.dev/os/signal)
