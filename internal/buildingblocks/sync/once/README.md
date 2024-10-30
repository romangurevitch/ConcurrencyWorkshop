# Understanding Go's `sync` Package: Once

The `sync` package in Go provides synchronization primitives to ensure safe concurrent access to shared resources.  
One of these primitives is the `Once` type, which ensures that a piece of code is executed only once.

<img src="https://securego.io/img/gosec.svg" alt="drawing" height="400"/>

## Table of Contents

1. [Introduction to Once](#introduction)
2. [Usage of Once](#usage)
3. [Use Cases](#use-cases)
4. [Common Pitfalls](#common-pitfalls)
5. [Best Practices](#best-practices)
6. [Resources](#resources)

## Introduction to Once

The `sync.Once` type is a type of synchronization primitive used to ensure that a particular piece of code is executed
only once, regardless of how many goroutines attempt to execute it.  
This is useful for initializing resources that are shared across multiple goroutines.

## Usage of Once

Here's an example of how to use the `sync.Once` type:

```go
package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func main() {
	once.Do(initialize)
	once.Do(initialize) // initialize will not be called again
}

func initialize() {
	fmt.Println("Initializing...")
}
```

In this code:

- A `sync.Once` variable named `once` is declared.
- The `Do` method of `once` is called with a function `initialize` as an argument.  
  The `initialize` function is executed the first time `Do` is called, but not the second time.

## Use Cases

- **Lazy Initialization**:
  - `sync.Once` is useful for lazy initialization where a resource is initialized only when it is
    needed.
- **Singleton Pattern**:
  - Ensuring a single instance of a struct is created in a concurrent environment.
- **One-time Setup**:
  - For setup that should only occur once but may be attempted from multiple goroutines.

## Common Pitfalls

- **Dependency Cycles**:
  - Be cautious of dependency cycles which could result in deadlocks when using `sync.Once`.
- **Error Handling**:
  - `sync.Once` does not provide built-in error handling, so if the initialization function can fail,
    you'll need to handle errors manually.

## Best Practices

- **Error Handling**:
  - Establish a robust error handling mechanism when using `sync.Once` for critical initialization.
- **Idempotency**:
  - Ensure the initialization function is idempotent if it may be called multiple times outside of
    a `sync.Once` context.

## Resources

- [Official Go Documentation on sync package](https://pkg.go.dev/sync#Once)

