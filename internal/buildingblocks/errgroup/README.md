# Understanding Go's `errgroup` Package: Group

The `errgroup` package in Go provides a straightforward way to manage the lifecycle of a group of goroutines, and their
associated error handling.

<img src="https://cdn.dribbble.com/users/1778913/screenshots/6562748/dribbble-machucado1.jpg?resize=400x300&vertical=center" alt="drawing" height="400"/>

## Table of Contents

1. [Introduction to Group](#introduction)
2. [Usage of Group](#usage)
3. [Use Cases](#use-cases)
4. [Common Pitfalls](#common-pitfalls)
5. [Best Practices](#best-practices)
6. [Resources](#resources)

## Introduction to Group

The `errgroup.Group` type provides synchronization, error propagation, and Context cancelation for groups of goroutines
working on sub-tasks of a common task.

## Usage of Group

```go
g, ctx := errgroup.WithContext(context.Background())

g.Go(func() error {
    // Your code here
    return nil  // return an error if something goes wrong
})

// Wait for all goroutines to finish and collect any errors
err := g.Wait()
```

## Use Cases

- **Concurrent Error Handling**: Executing multiple goroutines and aggregating their errors.
- **Context Propagation**: Propagating context and cancellation signals across a group of goroutines.

## Common Pitfalls

- **Error Ignorance**: Ignoring errors returned by `Group.Go`.
- **Misusing Context**: Misusing the context returned by `errgroup.WithContext` can lead to unintended behavior. For
  instance, storing values in the context that are supposed to be accessed by goroutines may lead to race conditions if
  not handled properly.

## Best Practices

- **Error Handling**: Always handle errors returned by `Group.Go`.
- **Context Usage**: Use the context returned by `errgroup.WithContext` to propagate cancelation.

## Resources

- [Official Documentation on errgroup.Group](https://pkg.go.dev/golang.org/x/sync/errgroup)
- [LEARNING GO: CONCURRENCY PATTERNS USING ERRGROUP PACKAGE](https://mariocarrion.com/2021/09/03/learning-golang-concurrency-patterns-errgroup-package.html)
