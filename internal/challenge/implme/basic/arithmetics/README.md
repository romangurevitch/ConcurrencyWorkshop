# Arithmetic Challenge: Parallel Sum Implementation

## Overview

The Arithmetic Challenge aims to enhance arithmetic computation performance by implementing a parallel sum
method, `ParallelSum`. This method is essential for executing summations concurrently, significantly reducing processing
time for large input sizes.

## Challenge Description

You are provided with an arithmetic package that includes a `SequentialSum` function. This function sequentially
performs summation of squares, which is not efficient for large input sizes. The challenge is to transform this
sequential computation into a parallel one using Go's concurrency features.

Your objective is to develop the `ParallelSum` function to carry out the summation in parallel, thereby reducing
execution time for substantial inputs.

## Provided Materials

- **Sequential Sum implementation**: A reference sequential summation function.
- **Sample process function**: Simulates a delay to mimic computation.
- **Starter code for `ParallelSum`**: A scaffold to guide your parallel implementation.
- **[Understanding Concurrent Design Patterns in Golang](../../../../pattern/README.md)**

## Goals

- Create the `ParallelSum` function in the `arithmetics` package to calculate the sum of squares in parallel.
- Outperform the sequential implementation for large input sizes in terms of efficiency.
- Ensure accurate results under parallel computation.

## Getting Started

1. Review the existing `SequentialSum` function to understand the current sequential computation.
2. Design and implement the `ParallelSum` function, use one of
   the provided concurrent patterns
   in [Understanding Concurrent Design Patterns in Golang](../../../../pattern/README.md).
3. Ensure your implementation is correct and that performance is enhanced for large input sizes.

## Running Tests and Benchmarks

To confirm that your `ParallelSum` function performs correctly and to evaluate its performance, you should run the
provided tests and benchmarks:

Run the unit tests and benchmarks to verify the correctness of your implementation:

```shell
    make implme-basic
```
