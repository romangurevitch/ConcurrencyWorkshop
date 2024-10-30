# Real-World Patterns and Applications

## Table of Contents

1. [Why Concurrency Matters](#why-concurrency-matters)
2. [Real-Life Problems and Solutions Summary](#real-life-problems-and-solutions-summary)
3. [Concurrent Design Patterns Examples in Go](#concurrent-design-patterns-examples-in-go)
    - [Future](#future)
    - [Pipeline](#pipeline)
    - [Fan-Out/Fan-In](#fan-outfan-in)
    - [Worker Pool](#worker-pool)
    - [Dynamic Rate-Limited Worker Pool](#dynamic-rate-limited-worker-pool)
    - [Publish-Subscribe (Pub/Sub)](#publish-subscribe-pubsub)
4. [Comparison Table](#comparison-table)

---

## Why Concurrency Matters

In today's fast-paced digital environment, applications must handle increasing loads, provide real-time
responses, and process data efficiently.  
Concurrency allows programs to execute multiple operations
simultaneously, leading to:

- **Improved Performance**:
    - Efficiently utilizes multi-core processors by parallelizing tasks.
- **Responsiveness**:
    - Keeps applications reactive, enhancing user experience during long-running
      operations.
- **Scalability**:
    - Handles growing workloads without a proportional increase in resource consumption.

---

## Real-Life Problems and Solutions Summary

### 1. Reducing Latency in Fetching Data from Multiple APIs

- **Problem**:
    - Your application needs to aggregate data from several internal or external APIs, such as
      fetching user details, recent activities, and recommendations.  
      Doing this sequentially increases latency, leading to a sluggish user experience.
- **Example**:
    - A social media app retrieving user profiles, friend lists, and recent posts when a user
      logs in.
- **Solution**:
    - Implement the **[Future](#future)** pattern to execute API calls concurrently,
      significantly reducing total wait time and improving responsiveness.

### 2. Efficiently Processing Data Through Multiple Stages

- **Problem**:
    - Your application processes data that must go through several transformations, such as
      parsing, validation, and storage.  
      Doing this sequentially can become a bottleneck.
- **Example**:
    - A file processing application that ingests uploaded CSV files, parses them, validates the
      data, and stores it in a database.
- **Solution**:
    - Use a **[Pipeline](#pipeline)** to process data concurrently at each stage, enhancing
      throughput and allowing for efficient data handling within a single application.

### 3. Accelerating Batch Processing of Independent Tasks

- **Problem**:
    - Your application needs to process a large number of independent tasks, like resizing images
      or computing analytics, which can be time-consuming when done one after another.
- **Example**:
    - A photo editing app that applies filters to a batch of images uploaded by the user.
- **Solution**:
    - Apply the **[Fan-Out/Fan-In](#fan-outfan-in)** pattern to distribute tasks across multiple
      goroutines within your application and aggregate results efficiently, drastically reducing processing
      time.

### 4. Managing High Volumes of Unpredictable Tasks Without Overloading Resources

- **Problem**:
    - Your application experiences unpredictable spikes in user requests or tasks, risking system
      overload if all tasks are processed simultaneously.
- **Example**:
    - A chat application handling message sending and receiving during peak usage times.
- **Solution**:
    - Employ a **[Worker Pool](#worker-pool)** within your application to control concurrency
      levels, distribute workloads evenly, and maintain system stability even under high load.

### 5. Complying with External Rate Limits While Maximizing Task Throughput

- **Problem**:
    - Your application interacts with external services that impose rate limits.  
      Exceeding these can lead to errors or service denial.
- **Example**:
    - An app that retrieves stock prices from a third-party API that allows only a certain number
      of requests per minute.
- **Solution**:
    - Utilize a **[Dynamic Rate-Limited Worker Pool](#dynamic-rate-limited-worker-pool)** to
      adjust processing rates dynamically within your application, ensuring compliance with rate limits while
      maximizing throughput.

### 6. Building Scalable, Decoupled Components for Asynchronous Event Handling

- **Problem**:
    - Tight coupling within your application's modules makes it difficult to scale and maintain,
      as changes in one part can cause unintended effects elsewhere.
- **Example**:
    - An online game where player actions need to update various systems like scoring,
      achievements, and notifications without causing performance issues.
- **Solution**:
    - Implement the **[Publish-Subscribe (Pub/Sub)](#publish-subscribe-pubsub)** pattern within
      your application to decouple components, allowing modules to communicate asynchronously and scale
      independently.

---

## Concurrent Design Patterns Examples in Go

Below are several concurrent design patterns in Go, each suited to different real-world scenarios within
a single application context.  
Understanding these patterns helps in designing systems that are both
efficient and maintainable.

### Future

**Scenario**:

- In your application, you need to perform several independent operations before proceeding,
  such as fetching data from multiple APIs or databases.  
  Doing these one after another slows down the
  overall processing time.

**When to Use**:

- When you have independent computations or I/O operations that can happen asynchronously.
- When you need to initiate tasks without blocking the main execution flow and retrieve results later.

**Why Use It**:

- **Reduced Latency**:
    - Executes independent tasks in parallel, minimizing total processing time.
- **Resource Efficiency**:
    - Allows the main thread to continue executing while waiting for results.
- **Simplified Error Handling**:
    - Manages asynchronous operations and their outcomes in a unified way.

**Real-Life Applications**:

- Fetching user data, preferences, and notifications concurrently upon login.
- Performing background computations while the user interacts with the UI.
- Aggregating data from multiple microservices within the same application.

**Code Example**:

[Further details and code example](future/README.md)

---

### Pipeline

**Scenario**:

- Your application processes data that goes through several stages, such as data ingestion,
  transformation, and storage.  
  Processing each item sequentially through all stages can become inefficient.

**When to Use**:

- When processing data involves multiple sequential stages.
- When dealing with continuous data that requires efficient handling within your application.

**Why Use It**:

- **Increased Throughput**:
    - Each stage processes data concurrently, maximizing resource utilization.
- **Modularity**:
    - Separates concerns by dividing processing into distinct stages within your application.
- **Scalability**:
    - Easily scales by adding more workers to stages to handle increased load.

**Real-Life Applications**:

- Processing user-uploaded files through validation, compression, and storage.
- Handling incoming network packets through parsing, filtering, and routing.
- Transforming and analyzing data streams within the application.

[Further details and code example](pipeline/README.md)

---

### Fan-Out/Fan-In

**Scenario**:

- Your application needs to perform resource-intensive computations on a set of data, such as
  generating thumbnails for a list of images or performing calculations on data sets.

**When to Use**:

- When tasks can be executed independently and in parallel.
- When results from these tasks need to be collected and possibly aggregated.

**Why Use It**:

- **Performance Boost**:
    - Reduces total computation time by utilizing multiple cores.
- **Efficient Aggregation**:
    - Collects and combines results seamlessly within your application.
- **Scalability**:
    - Can scale by adjusting the number of worker goroutines.

**Real-Life Applications**:

- Processing a batch of user requests simultaneously.
- Running multiple simulations or calculations concurrently.
- Generating reports based on different data segments at the same time.

[Further details and code example](fanoutin/README.md)

---

### Worker Pool

**Scenario**:

- Your application handles tasks that arrive unpredictably, such as incoming requests to a
  server or jobs added to a queue, and you need to process them without overwhelming the system.

**When to Use**:

- When you have a high volume of tasks that need to be processed asynchronously.
- When it's necessary to limit the number of concurrent operations to manage resources.

**Why Use It**:

- **Resource Management**:
    - Prevents system overload by controlling concurrency within your application.
- **Improved Throughput**:
    - Keeps workers busy, optimizing resource utilization.
- **Reliability**:
    - Helps in graceful degradation under heavy load conditions.

**Real-Life Applications**:

- Handling HTTP requests in a web server.
- Processing tasks from a job queue in a background service.
- Managing database queries to avoid connection pool exhaustion.

[Further details and code example](workerpool/README.md)

---

### Dynamic Rate-Limited Worker Pool

**Scenario**:

- Your application interacts with external APIs that enforce rate limits, and you need to
  ensure you don't exceed these limits while still processing tasks efficiently.

**When to Use**:

- When interacting with external services that enforce strict rate limits.
- When the processing rate needs to adjust dynamically based on external feedback or quotas.

**Why Use It**:

- **Compliance**:
    - Ensures adherence to external rate limits, avoiding errors.
- **Adaptability**:
    - Dynamically adjusts to changing rate limits or quotas.
- **Efficiency**:
    - Maximizes throughput within allowed limits.

**Real-Life Applications**:

- Fetching data from third-party APIs with request limits.
- Sending emails or notifications where providers limit the sending rate.
- Scraping websites that throttle based on request frequency.

[Further details and code example](dynamic/README.md)

---

### Publish-Subscribe (Pub/Sub)

**Scenario**:

- Within your application, you have components that need to communicate events or updates to
  other parts without tight coupling.

**When to Use**:

- When multiple components need to react to events without direct dependencies.
- When building systems that require scalability and maintainability within a single application.

**Why Use It**:

- **Loose Coupling**:
    - Allows independent development and scaling of modules.
- **Asynchronous Communication**:
    - Decouples the timing between event producers and consumers.
- **Scalability**:
    - Easily accommodates more subscribers or publishers within the application.

**Real-Life Applications**:

- An event system where user actions trigger updates in various modules.
- Notification systems within the app that inform different components of state changes.
- Logging and monitoring components reacting to events within the application.

[Further details and code example](pubsub/README.md)

---

## Comparison Table

This table provides a comparative overview of various concurrent design patterns in Go, highlighting
their key attributes, typical use cases, and application examples within a single application context.

| Design Pattern                   | Input Type | Process Duration | Synchronization                   | Latency                                                    | Throughput                                            | Data Flow                        | Use Case and Application Examples                                     |
|----------------------------------|------------|------------------|-----------------------------------|------------------------------------------------------------|-------------------------------------------------------|----------------------------------|-----------------------------------------------------------------------|
| Future                           | Single     | Short            | Blocking until result ready       | <span style="color:red">Potential Increased Latency</span> | Standard Throughput                                   | Request -> Computation -> Result | Async Computations, Async API Calls                                   |
| Pipeline                         | Unbounded  | Long             | Sequential Execution              | <span style="color:red">Sequential Latency</span>          | <span style="color:red">Sequential Throughput</span>  | Stage-wise Processing            | Stream Processing, Data Transformation Pipelines                      |
| Fan-out Fan-in                   | Bounded    | Short to Long    | Goroutine synchronization         | <span style="color:green">Reduced Latency</span>           | <span style="color:green">Increased Throughput</span> | Task -> Worker -> Aggregator     | CPU bound parallel tasks, Data Processing, Image Processing           |
| Worker Pool                      | Unbounded  | Short to Long    | Worker Coordination               | <span style="color:green">Reduced Latency</span>           | <span style="color:green">Increased Throughput</span> | Task -> Worker -> Result         | I/O or CPU Bound Tasks, Task Processing Systems                       |
| Dynamic Rate-Limited Worker Pool | Unbounded  | Long             | Rate Limiter, Worker Coordination | Controlled Latency                                         | Controlled Throughput                                 | Task -> Worker -> Result         | External Rate Limits, Resource Management, API Clients, Microservices |
| Pub-Sub                          | Unbounded  | Long             | Topic-based Subscription          | Event Delivery Latency                                     | Varied Based on Subscribers                           | Event Broadcast                  | Event Broadcasting, Event Notification Systems                        |

---
