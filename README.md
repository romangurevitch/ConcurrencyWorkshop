# Concurrency in Go Workshop

Welcome to the immersive full-day workshop on Fundamentals of Concurrent Programming in Go!
This workshop is crafted to deliver a profound understanding of concurrent programming techniques in Go, with a
substantial focus on pragmatic application and implementation.

Throughout this hands-on workshop, delve into the nuances of building production-grade services employing concurrent
design patterns and the sync package.
Under the expert guidance of the instructor, unravel the design philosophies and engineering decisions that underpin
effective concurrent programming in Go.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Getting Started](#Getting-Started)
3. [Concurrency Building Blocks in Go](#Concurrency-Building-Blocks-in-Go)
4. [Working with Goroutines](#Working-with-Goroutines)
5. [Common Concurrent Design Patterns](#common-concurrent-design-patterns)
6. [Challenges](#challenges)

## Prerequisites

Intermediate knowledge of the Go programming language is recommended.
If you are unfamiliar with Go, it's advisable to take the Go Tour and review basic Go concepts.

Basic experience with command-line tools is required.

## Getting Started

1. **Clone the Repository**:
    - Run the following command to clone the repository to your local machine:
      ```bash
      git clone https://github.com/romangurevitch/ConcurrencyWorkshop.git
      ```
    - Navigate to the project directory:
      ```bash
      cd ConcurrencyWorkshop
      ```

2. **Install or update Go to the latest version**:
    - Follow the [installation instructions](https://golang.org/doc/install) provided on the Go website.

3. **Makefile Help**:
    - While in the project directory, run the following command to view the available make targets and their
      descriptions:
      ```bash
      make help
      ```

4. **Verify Installation**:
    - Run the following command to execute the test suite and ensure everything is set up correctly:
      ```bash
      make test
      ```
5. **Open the Project in an IDE**:
    - It's recommended to open the project in an Integrated Development Environment (IDE) for a better programming
      experience. Two popular choices for Go development are:
        - [GoLand](https://www.jetbrains.com/go/): A powerful IDE by JetBrains dedicated to Go development.
        - [Visual Studio Code (VSCode)](https://code.visualstudio.com/): A free, open-source editor with support for Go
          via extensions like
          the [Go extension by Microsoft](https://marketplace.visualstudio.com/items?itemName=golang.Go).
    - Once you have your preferred IDE installed, open the project by navigating to `File -> Open...` and selecting the
      project directory you cloned earlier.

## Concurrency Building Blocks in Go

Dive into the building blocks of concurrency in Go by exploring the following topics:

- [Channels](internal/buildingblocks/channel/README.md)
- [Context](internal/buildingblocks/context/README.md)
- [ErrorGroup](internal/buildingblocks/errgroup/README.md)
- [Runtime (dir)](internal/buildingblocks/runtime)
    - [Gosched](internal/buildingblocks/runtime/gosched/README.md)
    - [Goexit](internal/buildingblocks/runtime/goexit/README.md)
- [Sync (dir)](internal/buildingblocks/sync)
    - [atomic](internal/buildingblocks/sync/atomic/README.md)
    - [locks](internal/buildingblocks/sync/locks/README.md)
    - [map](internal/buildingblocks/sync/map/README.md)
    - [once](internal/buildingblocks/sync/once/README.md)
    - [pool](internal/buildingblocks/sync/pool/README.md)

Navigate to the respective [directories](internal/buildingblocks) to find READMEs and code examples.

## Working with Goroutines

Discover Best Practices for Using Goroutines in Concurrency

- [Working with Goroutines](internal/goroutine/README.md)

Navigate to the respective [directories](internal/goroutine) to find READMEs and code examples.

## Common Concurrent Design Patterns

Explore various concurrent design patterns in Go:

[Comparison table](internal/pattern/README.md)

- [Future](internal/pattern/future/README.md)
- [Pipeline](internal/pattern/pipeline/README.md)
- [Fan-out Fan-in ](internal/pattern/fanoutin/README.md)
- [Worker Pool](internal/pattern/workerpool/README.md)
- [Dynamic Rate-Limited Worker Pool](internal/pattern/dynamic/README.md)
- [Pub-Sub](internal/pattern/pubsub/README.md)

Navigate to the respective [directories](internal/pattern) to find READMEs and code examples.

## Challenges

Take on a variety of challenges to test your understanding of concurrency in Go:

- [Fix Me](internal/challenge/fixme)
- [Implement Me](internal/challenge/implme)

Navigate to the respective [directories](internal/challenge) to find READMEs and code examples.
