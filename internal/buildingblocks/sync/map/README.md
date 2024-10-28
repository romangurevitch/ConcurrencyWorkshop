# Understanding Go's `sync` Package: Map

The `sync` package in Go provides the `Map` type, a concurrent-safe map implementation.

<img src="../../../../docs/images/sync_map.png" alt="drawing" height="400"/>

## Table of Contents

1. [Introduction to Map](#introduction)
2. [Usage of Map](#usage)
3. [Use Cases](#use-cases)
4. [Common Pitfalls](#common-pitfalls)
5. [Best Practices](#best-practices)
6. [Resources](#resources)

## Introduction to Map

`sync.Map` is a concurrent map type with amortized-constant-time loads, stores, and deletes.

## Usage of Map

```go
var m sync.Map

func main() {
    m.Store("hello", "world")
    if value, ok := m.Load("hello"); ok {
        fmt.Println(value)  // Output: world
    }
}
```

## Use Cases
The Map type is optimized for two common use cases

- When the entry for a given key is only ever written once but read many times, as in caches that only grow.
- when multiple goroutines read, write, and overwrite entries for disjoint sets of keys. In these two cases, use of a Map may significantly reduce lock contention compared to a Go map paired with a separate Mutex or RWMutex.

## Common Pitfalls

- Not a general replacement for Go's built-in map type due to its specific performance characteristics.

## Best Practices

- Prefer Go's built-in map type with proper locking for most use cases.

## Resources

- [Official Go Documentation on sync.Map](https://pkg.go.dev/sync#Map)
- [Effective Go](https://golang.org/doc/effective_go.html#concurrency)
