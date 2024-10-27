# Race detection flag

## Overview

Data races are among the most common and hardest to debug types of bugs in concurrent systems. 

A data race occurs when two goroutines access the same variable concurrently and at least one of the accesses is a write.

## Usage

```go
go test -race mypkg // to test the package
go run -race mysrc.go // to run the source file
go build -race mycmd  // to build the command
go install -race mypkg // to install the package
```

## Runtime Overhead

The cost of race detection varies by program, but for a typical program,
memory usage may increase by **5-10x** and execution time by **2-20x**.

## Additional information

[Race detector docs](https://go.dev/doc/articles/race_detector)
