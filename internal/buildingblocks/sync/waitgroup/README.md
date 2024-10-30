# Understanding Golang WaitGroup

The `sync.WaitGroup` is a crucial synchronization primitive in Go, which allows a program to wait for a collection of
goroutines to finish executing.

<img src="../../../../docs/images/gopher_wait.png" alt="drawing" height="300"/>

## Table of Contents

1. [Introduction to Golang WaitGroup](#introduction)
2. [Using WaitGroup](#using-waitgroup)
3. [WaitGroup Functions](#waitgroup-functions)
4. [Common Pitfalls and Issues](#common-pitfalls-and-issues)
5. [Best Practices](#best-practices)
6. [Resources](#resources)

## Introduction

A `WaitGroup` waits for a collection of goroutines to finish executing, making it a simple and effective way to wait for
goroutines without needing to set up more complex channel structures.

## Using WaitGroup

Here's a simplified example of how to use a `WaitGroup` to wait for three goroutines to finish:

```go
package main

import (
	"fmt"
	"sync"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter.
		go func(i int) {
			defer wg.Done() // Decrement the counter when the goroutine completes.
			worker(i)
		}(i)
	}

	wg.Wait() // Wait for all the workers to finish.
	fmt.Println("All workers done")
}
```

## WaitGroup Functions

- **`Add(delta int)`**: Increments the `WaitGroup` counter by `delta`.
- **`Done()`**: Decrements the `WaitGroup` counter by one.
- **`Wait()`**: Blocks until the `WaitGroup` counter is zero.


## Common Pitfalls and Issues

- **Forget to call `wg.Add()` or `wg.Done()`**:
    - If you forget to increment or decrement the counter, `wg.Wait()` will not wait or block your goroutines
      indefinitely.


- **Calling `Add` Inside the Goroutine**:
    - Calling the `Add` method inside the goroutine instead of outside can cause a race condition where `wg.Wait()` is
      called before `wg.Add()`.

    ```go
    // Incorrect
    go func() {
        wg.Add(1)  // Risk of race condition
        defer wg.Done()
        worker(i)
    }()
    ```

- **Negative Counter**:
    - Calling `wg.Done()` more times than `wg.Add()` will cause a panic due to a negative WaitGroup counter.
    ```go
    wg.Add(1)
    wg.Done()
    wg.Done()  // Panic: negative WaitGroup counter
    ```

- **Reuse of a WaitGroup before all goroutines call Done**:
    - Reusing a WaitGroup for new tasks before all the previous tasks have called Done may lead to undefined behavior.
    ```go
    wg.Add(1)
    go worker(&wg, 1)
    wg.Wait()
    wg.Add(1)  // Undefined behavior if worker hasn’t called Done yet
    ```

## Best Practices

- **Deferred Call to `wg.Done()`**:
    - It's a good practice to defer the call to `wg.Done()` to ensure it gets called whenever the function returns.
    ```go
    func worker(wg *sync.WaitGroup, id int) {
        defer wg.Done()  // Good practice
    }
    ```

- **Separation of Concerns**:
    - Keep the concurrency control (`Add` and `Done`) close together and separate from the business logic of the
      goroutine.  
      Do not pass the `WaitGroup` forward to other functions or goroutines.  
      This makes the code easier to understand and maintain.

    ```go
    // Correct
    wg.Add(1)
    go func(i int) {
        defer wg.Done()
        worker(i)  // Business logic is separate
    }(i)
    ```

- **If Necessary, Pass WaitGroup as a Pointer**:
    - Always pass `WaitGroup` to functions as a pointer, as `WaitGroup` contains a counter that should be shared among
      all goroutines.
    ```go
    func worker(wg *sync.WaitGroup, id int) { /* ... */ }
    ```

## Resources

- [Go by Example: WaitGroups](https://gobyexample.com/waitgroups)
- [sync.WaitGroup documentation](https://pkg.go.dev/sync#WaitGroup)