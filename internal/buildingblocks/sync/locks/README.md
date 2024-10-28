# Understanding Go's `sync` Package: Locks

The `sync` package in Go is essential for writing concurrent programs, providing synchronization primitives such as
locks to manage access to shared resources.

<img src="https://miro.medium.com/v2/1*bZHBo75FSyKre5pk2-HPmw.png" alt="drawing" height="400"/>

## Table of Contents

1. [Introduction to Locks](#introduction)
2. [Mutex](#mutex)
3. [RWMutex](#rwmutex)
4. [Use Cases](#use-cases)
5. [Common Pitfalls](#common-pitfalls)
6. [Best Practices](#best-practices)
7. [Resources](#resources)

## Introduction to Locks

Locks are fundamental for managing concurrent access to shared resources. They ensure that only one goroutine can access
a resource at a time, preventing data races and ensuring data consistency. Two primary types of locks are provided in
Go's `sync` package: `Mutex` and `RWMutex`.

## Mutex

A Mutex (mutual exclusion lock) provides exclusive access to a resource.

```go
var mu sync.Mutex

func accessResource() {
    mu.Lock()  // Lock the mutex before accessing the resource
    // Access the shared resource
    mu.Unlock()  // Unlock the mutex after accessing the resource
}
```

### Key Points

- `Lock()` acquires the lock, blocking if necessary until available.
- `Unlock()` releases the lock.
- Always unlock the mutex to prevent deadlocks.

## RWMutex

RWMutex is a reader/writer mutual exclusion lock allowing concurrent read access but exclusive write access.

```go
var rwmu sync.RWMutex

func readResource() {
    rwmu.RLock()  // Lock the mutex for reading
    // Read the shared resource
    rwmu.RUnlock()  // Unlock the mutex after reading
}

func writeResource() {
    rwmu.Lock()  // Lock the mutex for writing
    // Write to the shared resource
    rwmu.Unlock()  // Unlock the mutex after writing
}
```

### Key Points

- Multiple goroutines can hold a read lock at the same time, but only one can hold a write lock.
- `RLock()` and `RUnlock()` are used for reading, `Lock()` and `Unlock()` are used for writing.

## Use Cases

Locks are crucial when multiple goroutines need to access shared resources but not simultaneously to prevent data races
and ensure data consistency.

- Protecting shared data structures.
- Managing access to shared resources like file handles or network connections.

## Common Pitfalls

- Deadlocks: Forgetting to unlock a mutex or locking it multiple times without unlocking can lead to deadlocks.
- Runtime Panics: Unlocking a mutex that is not locked or in the wrong goroutine can cause runtime panics.
- Starvation: Excessive use of write locks can starve readers in the case of `RWMutex`.

## Best Practices

- Always unlock a mutex in the same goroutine that locked it.
- Use `defer` to ensure a mutex is unlocked even if a function exits early.
- Prefer `RWMutex` if the resource is read-mostly to allow concurrent read access.

## Resources

- [Official Go Documentation on sync package](https://pkg.go.dev/sync#Mutex)
- [Golang Mutex - A Complete Guide](https://www.kelche.co/blog/go/mutex)

