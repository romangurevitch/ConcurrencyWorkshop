# Understanding the Future Design Pattern in Go

The Future design pattern is a concurrency pattern that represents a result that will become available at some point in the future.  
It's particularly useful for handling asynchronous computations and can significantly improve the
performance and responsiveness of your Go applications.

This guide will explain how to implement and use the Future pattern in Go, focusing on practical aspects, common issues,
and best practices.  
We'll walk through a concrete implementation and demonstrate how to integrate it into your projects.

---

## Table of Contents

1. [Introduction](#introduction)
2. [Implementation Example](#implementation-example)
3. [How to Use the Future Implementation](#how-to-use-the-future-implementation)
4. [Common Issues and Pitfalls](#common-issues-and-pitfalls)
5. [Best Practices](#best-practices)
6. [Common Implementation](#Common-Implementations)

---

## Introduction

In Go, the Future pattern can be implemented using goroutines and channels.  
It allows you to start a computation
asynchronously and retrieve the result later, enabling non-blocking execution and better resource utilization.

<img src="../../../docs/images/future_graph.png" alt="drawing" height="300"/>

This pattern is especially beneficial when dealing with:

- **I/O-bound operations**: Such as network requests or file I/O.
- **CPU-bound computations**: Intensive calculations that can run concurrently without affecting the main execution
  flow.

---

## Implementation Example

See [main.go](main.go)

To implement the Future pattern, you can define a generic `Future` type that encapsulates an asynchronous computation.
The `Future` uses a channel to deliver the result once it's available and handles context cancellation.

**Key Components:**

- **`Result[T any]`**: A generic type that holds the value or an error from the computation.
- **`Future[T any]`**: Encapsulates the future result using a channel.
- **`ProcessFunc[T any]`**: A function type that performs the computation.
- **`NewFuture`**: Initializes the `Future` and starts the asynchronous operation, handling context cancellation.
- **`Result()`**: Retrieves the computation result, blocking if it's not yet available.

---

## How to Use the Future Implementation

### Step 1: Define the Process Function

Create a function that matches the `ProcessFunc[T]` signature.  
This function performs the asynchronous task.

```go
func fetchData(ctx context.Context) (DataType, error) {
    // Perform computation or I/O-bound operation.
}
```

### Step 2: Create a Future Instance

Initialize a `Future` by passing the context and your process function to `NewFuture`.

```go
future := NewFuture(ctx, fetchData)
```

### Step 3: Perform Other Tasks (Optional)

While the `Future` is computing, you can perform other operations without blocking.

```go
// Do other work...
```

### Step 4: Retrieve the Result

Call the `Result()` method to get the computation result.  
This call will block until the result is available.

```go
result := future.Result()
if result.Err != nil {
    // Handle error.
}
// Use result.Value.
```

---

## Common Issues and Pitfalls

### 1. Blocking Forever

**Issue**: If the computation never completes or the result channel is not closed, `Result()` will block indefinitely.

**Solution**: Ensure that the Future implementation handles context cancellation properly.  
The `Future` should check the context before and during execution, returning promptly if the context is canceled to
prevent indefinite blocking.

### 2. Ignoring Context Cancellation When Implementing Future

**Issue**: If the Future implementation does not respect context cancellation, it may not return promptly when the
context is canceled.  
This can lead to the `Result()` method blocking indefinitely or unnecessary computations continuing
in the background.

**Solution**: Ensure that the Future implementation checks for context cancellation before starting the computation and
returns immediately if the context is done.  
In this example implementation, the `NewFuture` function selects on `ctx.Done()` before executing the process function.

```go
go func () {
  defer close(f.result)
  select {
      case <-ctx.Done():
          f.result <- Result[T]{Err: ctx.Err()}
      default:
          value, err := processFunc(ctx)
          f.result <- Result[T]{Value: value, Err: err}
      }
}()
```

By doing this, the Future respects context cancellation, ensuring that it doesn't block indefinitely and doesn't perform
unnecessary work if the context is canceled.

### 3. Resource Leaks

**Issue**: Failing to release resources like open files or network connections can lead to resource exhaustion.

**Solution**: Use `defer` statements to ensure resources are released, and handle all error cases where resources might
not be automatically released.

### 4. Unhandled Errors

**Issue**: Errors returned by the process function may be ignored if not properly checked.

**Solution**: Always check `result.Err` after retrieving the result and handle errors appropriately.

---

## Best Practices

### 1. Use Contexts Wisely

- Pass contexts to your process functions to manage cancellation and timeouts.
- Ensure that your Future implementation respects context cancellations to avoid unnecessary work.

### 2. Handle Errors Properly

- Always check for errors when retrieving the result.
- Propagate errors back to the caller or handle them within the process function.

### 3. Limit Concurrency

- Be cautious when creating many futures in a loop; uncontrolled concurrency can overwhelm system resources.
- Consider using a worker pool or limiting the number of concurrent futures.

---

## Common Implementations

- **ErrGroup**: The [`errgroup`](https://pkg.go.dev/golang.org/x/sync/errgroup) package in Go provides a way to
  synchronize and collect errors from a group of goroutines, effectively implementing a form of the Future pattern.
  With `errgroup`, you can start multiple goroutines, wait for all of them to complete, and collect any errors that
  occur.
