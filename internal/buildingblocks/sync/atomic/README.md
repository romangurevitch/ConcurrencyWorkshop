# Understanding Go's `sync/atomic` Package

The `sync/atomic` package in Go provides low-level atomic memory primitives useful for implementing synchronisation
algorithms.

<img src="https://golangforall.com/assets/kanat.svg" alt="drawing" height="300"/>

## Table of Contents

1. [Introduction to Atomic Operations](#introduction)
2. [Atomic Variables](#atomic-variables)
3. [Atomic Functions](#atomic-functions)
4. [Use Cases](#use-cases)
5. [Common Pitfalls](#common-pitfalls)
6. [Best Practices](#best-practices)
7. [Resources](#resources)

## Introduction

Atomic operations provide a way to read or modify values in memory in a way that ensures that no other operations can
occur concurrently that would see a partially completed state or interfere with the operation.

## Atomic Variables

```go
import "sync/atomic"

counter := atomic.Int64{}
```

## Atomic Functions

### Add

```go
// Increment counter atomically
value := atomic.AddInt64(&counter, 1)
```

### Load and Store

```go
// Atomically load the value of counter
value := atomic.LoadInt64(&counter)

// Atomically store a value in counter
atomic.StoreInt64(&counter, 42)
```

### Compare and Swap

```go
// Atomically compare counter to old and if they are equal set counter to new
swapped := atomic.CompareAndSwapInt64(&counter, old, new)
```

## Examples and tests

See [package](.)

## Use Cases

Atomic operations are used when you need to ensure that operations on memory are completed without interruption, such as
when implementing counters, flags, and other synchronisation primitives.

## Common Pitfalls

- Mixing atomic operations with regular, non-atomic operations on the same variable can lead to data races and
  unpredictable behavior.

```go
var counter int64 // shared variable

// Incorrect Usage: Mixing atomic and non-atomic operations
func incrementNonAtomic() {
    counter++
}

func incrementAtomic() {
    atomic.AddInt64(&counter, 1)
}
```

## Best Practices

- Avoid mixing atomic and non-atomic operations on the same variable.
- Utilize the specialized types in the sync/atomic package over general integer types for atomic operations to ensure
  type safety, platform independence, and clearer, less error-prone code.

```go
var counterInt int
var counterInt64 atomic.Int64

// Avoid increment using int variable
atomic.AddInt64((*int64)(&counterInt), 1)

// Instead increment using atomic.Int64 variable
counterInt64.Add(1)
```

## Resources

- [Official Go Documentation on sync/atomic](https://pkg.go.dev/sync/atomic)
- [Mastering Atomic Counters in Go](https://towardsdev.com/mastering-atomic-counters-in-go-a-guide-to-efficient-state-management-cbcd2a4e5b0)

