# FileFinder Challenge

## Disclaimer

**Advanced Challenge**: This task requires a strong understanding of Go's concurrency features and is considered
advanced.

It is recommended that participants tackle the simpler challenges provided in this workshop before attempting
this Challenge.

This task is intended to test and improve your concurrency skills in Go and is best approached
with prior experience in concurrent programming.

## Overview

Welcome to the FileFinder Challenge! This workshop presents you with the opportunity to transform an existing sequential
file finder algorithm into a concurrent one using Go's robust concurrency features.


<img src="../../../../../docs/images/gopher_find.png" alt="drawing" height="400"/>

## Challenge Description

The provided baseline is a sequential implementation that searches for a file from a specified root directory.

Although it is functional, it does not utilize concurrency, leading to suboptimal performance on large filesystems.

Your mission, should you choose to accept it, is to implement the high-performance, concurrent file finder
without using `filepath.WalkDir`, and instead making use of the `findInDir` auxiliary function.

### Sequential Implementation

We have supplied you with the sequential version of the FileFinder, complete with tests and benchmarks. Reviewing these
will not only help you understand the current logic but also provide a performance benchmark for your concurrent
implementation.

### Concurrent Version

Your [concurrent solution](concurrent.go) should:

- Expedite the file search process via parallel directory searches.
- Employ effective goroutine synchronisation.
- Properly manage errors and context cancellations.
