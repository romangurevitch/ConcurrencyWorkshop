# Understanding Go's `context` Package

The `context` package in Go provides essential functionality for passing deadlines, cancelation signals, and other
request-scoped values across API boundaries and between processes.

<img src="https://miro.medium.com/v2/resize:fit:946/1*fSq3uLTcwVUvZQWXyTgRTQ.png" alt="drawing" height="400"/>

## Table of Contents

1. [Introduction to Context](#introduction)
2. [Usage of Context](#usage)
3. [Use Cases](#use-cases)
4. [Common Pitfalls](#common-pitfalls)
5. [Best Practices](#best-practices)
6. [Resources](#resources)

## Introduction to Context

`context.Context` is a type that carries deadlines, cancellations, and other common request-scoped values across API
boundaries and between processes.

## Usage of Context

```go
parentCtx := context.Background()  // Get an empty context

ctx, cancel := context.WithCancel(parentCtx)
cancel()  // Cancel the context

ctx, cancel := context.WithDeadline(parentCtx, deadline)
cancel()  // Cancel the context (optional if deadline reaches first)

ctx, cancel := context.WithTimeout(parentCtx, timeout)
cancel()  // Cancel the context (optional if timeout reaches first)

ctx := context.WithValue(parentCtx, key, value)  // Associate a key-value pair with context
value := ctx.Value(key)
if value != nil {
    fmt.Println("Value found:", value)
}

select {
case <-ctx.Done():  // Waiting for the context to be cancelled
    fmt.Println("Context cancelled:", ctx.Err())  // Print the reason for cancellation when it occurs
}
```

## Examples and tests

See [package](.)

## Use Cases

- **Request Scoping**:
    - Passing data that's scoped to a particular request through the call stack.
- **Deadline Propagation**:
    - Ensuring that operations complete within a specified amount of time.
- **Cancellation Propagation**:
    - Propagating cancellation signals to free up resources when operations are no longer
      needed.
- **Tracing Propagation**:
    - Context is commonly used to carry tracing information across API boundaries and between
      processes for monitoring and debugging purposes.  
      This is crucial for microservices architecture where a request might span multiple services and you want to trace
      its execution path.

## Common Pitfalls

- **Misuse of context.WithValue**:
    - According to Go's documentation, programmers should define their own types for keys
      to avoid collisions.  
      The correct way to use `context.WithValue` is to define a new type for the key, and use a value of that type as
      the
      key.
- **Misuse of ctx.Done()**:
    - The `ctx.Done()` method returns a channel that gets closed when the context gets
      cancelled.  
      This channel can be used to listen for the cancellation event in a select statement or in a blocking receive
      operation.  
      However, it inherits the usual challenges associated with channels in Go.
- **Ignoring Cancelation**:
    - Failing to respect cancelation signals can lead to resource leaks.
- **Overloading Context**:
    - Storing too much data in context or using it as a means of passing optional parameters.

## Best Practices

- **Minimalism**:
    - Only store essential data in context.
- **Cancelation Respect**:
    - Always respect cancelation signals to ensure resource cleanup.

## Resources

- [Official Go Blog - Context](https://blog.golang.org/context)
- [Go Documentation on context package](https://pkg.go.dev/context)
