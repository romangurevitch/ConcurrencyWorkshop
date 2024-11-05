# Understanding the Pipeline Design Pattern in Go

The Pipeline design pattern is a powerful concurrency model in Go that allows for processing data through a series of stages, where each stage performs a specific operation.  
This pattern enables efficient data processing by leveraging Go's goroutines and channels to create a chain of processing steps, improving the performance and scalability of your applications.

This guide will explain how to implement and use the Pipeline pattern in Go, focusing on practical aspects, common issues, and best practices.  
We'll walk through a step-by-step implementation and demonstrate how to integrate it into your projects.

---

## Table of Contents

1. [Introduction](#introduction)
2. [Implementation Example](#implementation-example)
3. [How to Use the Pipeline Implementation](#how-to-use-the-pipeline-implementation)
4. [Common Issues and Pitfalls](#common-issues-and-pitfalls)
5. [Best Practices](#best-practices)
6. [Resources](#resources)

---

## Introduction

In Go, the Pipeline pattern involves setting up a series of stages connected by channels, where each stage is a goroutine that processes data and passes the result to the next stage.  
This pattern is particularly useful for processing streams of data, transforming data through multiple steps, and handling complex workflows.

![Pipeline Diagram](../../../docs/images/pipeline_graph.png)

This pattern is beneficial when dealing with:

- **Data Transformation**: Applying a series of transformations to data.
- **Concurrent Processing**: Leveraging multiple CPU cores to process data in parallel.
- **Modular Design**: Breaking down complex processing into manageable, reusable components.

**Note:** In this implementation, each stage operates with a single goroutine, processing one item at a time in the order they are received.  
This design keeps the implementation simple and predictable.  
However, you can extend this pattern to allow each stage to process multiple items concurrently by incorporating worker pools or similar concurrency mechanisms within each stage.

---

## Implementation Example

See [package](.)

The implementation uses a generic `Pipe` function to create pipeline stages.  
Each stage reads from an input channel, processes the data using a provided function, and sends the result to an output channel.  
The stages are connected together to form a pipeline.

**Key Components:**

- **`Result[T any]`**: A generic type that holds a value and an error.
- **`ProcessFunc[T any, U any]`**: Defines how to process data from type `T` to type `U`.
- **`Pipe`**: Creates a pipeline stage that processes data from an input channel and sends results to an output channel.

---

## How to Use the Pipeline Implementation

### Step 1: Define the Processing Functions

Create functions that match the `ProcessFunc[T, U]` signature.  
Each function represents a stage in the pipeline.

```go
func processStage1(ctx context.Context, input Result[T]) Result[U] {
    // Implement your processing logic for stage 1.
    // Return the result and any errors.
}

func processStage2(ctx context.Context, input Result[U]) Result[V] {
    // Implement your processing logic for stage 2.
    // Return the result and any errors.
}

// Add more stages as needed.
```

For example, you might have functions that fetch data and then process it:

```go
func fetchData(ctx context.Context, input Result[int]) Result[string] {
    // Fetch data based on input.Value.
    // Return Result[string]{Value: data, Err: err}
}

func processData(ctx context.Context, input Result[string]) Result[bool] {
    // Process the fetched data.
    // Return Result[bool]{Value: true/false, Err: err}
}
```

### Step 2: Create the Initial Input Channel

Initialize a channel to serve as the input to the first pipeline stage and populate it with data.

```go
inputCh := make(chan Result[T])

go func() {
    defer close(inputCh)
    // Send data into the inputCh.
    // for _, data := range dataSet {
    //     inputCh <- Result[T]{Value: data}
    // }
}()
```

### Step 3: Set Up the Pipeline Stages

Use the `Pipe` function to create each stage of the pipeline, connecting the output of one stage to the input of the next.

```go
stage1Ch := Pipe(ctx, inputCh, processStage1)
stage2Ch := Pipe(ctx, stage1Ch, processStage2)
// Continue chaining stages as needed.
```

**Note:** In this implementation, each `Pipe` function starts a single goroutine for the stage, processing one item at a time in the order they are received.  
If you need to process multiple items concurrently within a stage, consider extending the `Pipe` function to include a worker pool or similar concurrency mechanism.

### Step 4: Consume the Output

Read from the output channel of the final stage to retrieve the processed results.

```go
for result := range stage2Ch {
    if result.Err != nil {
        // Handle error.
        continue
    }
    // Use result.Value.
}
```

---

## Common Issues and Pitfalls

### 1. Deadlocks Due to Unclosed Channels

**Issue**: If channels are not properly closed, goroutines may block indefinitely, leading to deadlocks.

**Solution**: Ensure that all input channels are closed once no more data will be sent. Use `defer close(channel)` in the goroutine that writes to the channel.

### 2. Goroutine Leaks

**Issue**: Goroutines may continue running if they are not properly managed, consuming resources unnecessarily.

**Solution**: Ensure that goroutines exit when the context is canceled. Check `ctx.Done()` within your processing functions and return promptly when canceled.

### 3. Context Cancellation Not Respected

**Issue**: If your processing functions do not respect context cancellation, the pipeline may not shut down gracefully.

**Solution**: Pass `context.Context` to all processing functions and check for cancellation within each function.

### 4. Error Handling

**Issue**: Errors may not be properly propagated through the pipeline, leading to incorrect results or silent failures.

**Solution**: Use the `Result[T]` type to encapsulate both the value and any errors.  
Check for errors at each stage and handle them appropriately.

---

## Best Practices

### 1. Use Contexts Wisely

- **Cancellation and Timeouts**: Pass `context.Context` to your processing functions to handle cancellation and timeouts.
- **Respect Context**: Ensure that all goroutines check the context and exit promptly when canceled.

### 2. Proper Channel Management

- **Closing Channels**: Close channels when no more data will be sent to signal downstream stages.
- **Buffered Channels**: Use buffered channels if stages produce data faster than they are consumed to prevent blocking.

### 3. Error Handling

- **Propagate Errors**: Use the `Result[T]` type to pass errors downstream.
- **Handle Errors at Each Stage**: Check for errors in your processing functions and decide whether to continue or abort processing.

### 4. Modular Design

- **Reusable Components**: Design your processing functions to be reusable across different pipelines.
- **Separation of Concerns**: Keep each stage focused on a single task for better maintainability.

### 5. Avoid Shared State

- **Immutable Data**: Prefer passing data by value or using immutable data structures to avoid synchronization issues.
- **Synchronization Primitives**: If shared state is necessary, protect it with synchronization mechanisms like mutexes.

---

## Resources

- [Go Concurrency Patterns: Pipelines and cancellation](https://blog.golang.org/pipelines)
- [Building pipelines in Go](https://medium.com/statuscode/pipeline-patterns-in-go-a37bb3a7e61d)
