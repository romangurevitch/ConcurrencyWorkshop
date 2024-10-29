# Understanding the Fan-Out/Fan-In Design Pattern in Go

The Fan-Out/Fan-In design pattern is a fundamental concept in concurrent programming, where multiple tasks are distributed among several goroutines (**fan-out**), and their results are collected and combined into a single channel or data structure (**fan-in**).  
This pattern leverages Go's powerful concurrency primitives—goroutines and channels—to perform tasks in parallel, improving the performance and scalability of your applications.

This guide will explain how to implement and use the Fan-Out/Fan-In pattern in Go, focusing on practical aspects, common issues, and best practices.  
We'll walk through a concrete implementation and demonstrate how to integrate it into your projects.

---

## Table of Contents

1. [Introduction](#introduction)
2. [Implementation Example](#implementation-example)
3. [How to Use the Fan-Out/Fan-In Implementation](#how-to-use-the-fan-outfan-in-implementation)
4. [Common Issues and Pitfalls](#common-issues-and-pitfalls)
5. [Best Practices](#best-practices)
6. [Resources](#resources)

---

## Introduction

In Go, the Fan-Out/Fan-In pattern leverages goroutines and channels to perform concurrent operations efficiently.  
The **fan-out** phase involves spawning multiple goroutines to perform tasks in parallel, while the **fan-in** phase involves collecting the results from these goroutines back into a single channel or aggregated result.

![Fan-Out/Fan-In Diagram](../../../docs/images/fanout_in_graph.png)

This pattern is particularly beneficial when dealing with:

- **I/O-bound operations**: Such as making multiple network requests simultaneously.
- **CPU-bound computations**: Tasks that can be parallelized to utilize multiple CPU cores.
- **Data processing pipelines**: Where data needs to be processed in parallel steps.

---

## Implementation Example

See [main.go](main.go)

To implement the Fan-Out/Fan-In pattern, you can create a function that spawns a goroutine for each job and collects the results via a channel.  
The provided implementation uses generics to allow flexibility with different data types.

**Key Components:**

- **`Job[T any]`**: A generic type that holds the job ID and the value to process.
- **`Result[T any, U any]`**: A generic type that holds the job, the result value, and any error that occurred during processing.
- **`ProcessFunc[T any, U any]`**: A function type that defines how to process a job's value.
- **`FanOut`**: The function that fans out the jobs to goroutines and fans in the results.

---

## How to Use the Fan-Out/Fan-In Implementation

### Step 1: Define the Process Function

Create a function that matches the `ProcessFunc[T, U]` signature.  
This function performs the processing task.

```go
func processData(ctx context.Context, value T) (U, error) {
    // Perform processing on value.
    // Return the result and any errors.
}
```

For example, you might have a function that squares a non-negative integer:

```go
func squareNonNegative(ctx context.Context, value int) (int, error) {
    if value < 0 {
        return 0, errors.New("negative value")
    }
    return value * value, nil
}
```

### Step 2: Prepare the Jobs

Create a slice of `Job[T]` that you want to process.

```go
var jobs []Job[int]
for i := 1; i <= numOfJobs; i++ {
    jobs = append(jobs, Job[int]{ID: i, Value: i})
}
```

### Step 3: Fan-Out the Jobs

Call the `FanOut` function with the context, jobs, and process function.

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

results := FanOut(ctx, jobs, processData)
```

### Step 4: Fan-In the Results

Iterate over the results channel to collect and handle the results.

```go
for result := range results {
    if result.Err != nil {
        // Handle error.
        continue
    }
    // Use result.Value.
}
```

For example:

```go
for result := range results {
    if result.Err != nil {
        log.Printf("Error processing job %d: %v", result.Job.ID, result.Err)
        continue
    }
    log.Printf("Result for job %d: %v", result.Job.ID, result.Value)
}
```

---
## Common Issues and Pitfalls

### 1. Context Cancellation Not Respected

**Issue**: Goroutines may continue running even after cancellation, wasting resources and potentially causing inconsistent state.

**Solution**: Ensure that your `ProcessFunc` and any loops within it check `ctx.Done()` and return promptly when the context is canceled.

### 2. Resource Leaks

**Issue**: Failing to release resources like open files or network connections can lead to resource exhaustion.

**Solution**: Use `defer` statements to release resources and handle all error cases where resources might not be automatically released.

### 3. Ignoring Errors

**Issue**: Not handling errors returned by the processing function can lead to incorrect program behavior.

**Solution**: Always check `result.Err` and handle errors appropriately in the fan-in phase.

### 4. Race Conditions

**Issue**: Concurrent access to shared variables without proper synchronization can cause race conditions.

**Solution**: Avoid shared mutable state.  
If necessary, use synchronization primitives like mutexes to protect shared data.

---

## Best Practices

### 1. Proper Synchronization

- **WaitGroups**: Use `sync.WaitGroup` to ensure all goroutines finish executing before closing channels or exiting the program.

- **Channel Management**: Ensure channels are properly closed to prevent deadlocks.

### 2. Handle Context Cancellation

- **Pass Contexts**: Pass `context.Context` to your `ProcessFunc` to handle cancellation and timeouts.

- **Respect Context**: Ensure that your goroutines check the context and exit promptly when canceled.

### 3. Error Handling

- **Check Errors**: Always check for errors in the results and handle them appropriately.

- **Graceful Degradation**: Decide whether to continue processing or cancel based on the error.

### 4. Avoid Shared State

- **Immutable Data**: Prefer passing data by value or using immutable data structures to avoid the need for synchronization.

- **Synchronization Primitives**: If shared state is necessary, protect it with synchronization primitives like mutexes.

### 5. Choose Appropriate Patterns for Concurrency Control

- **Fixed Number of Jobs**: Since this pattern is designed for a predefined number of jobs, it's acceptable to start a goroutine for each job.

- **Limiting Concurrency**: If you need to limit concurrency (e.g., when dealing with a large or unknown number of jobs), consider using a different pattern like a **Worker Pool** or **Dynamic Rate-Limited Worker Pool**.

---

## Resources

- [Go Concurrency Patterns: Pipelines and cancellation](https://blog.golang.org/pipelines)