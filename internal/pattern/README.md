# Real-World Patterns and Applications

## Table of Contents

1. [Why Use Concurrency and Parallelism?](#why-use-concurrency-and-parallelism)
    - [Improved Throughput and Resource Utilisation](#Improved-Throughput-and-Resource-Utilisation)
    - [Low-Latency, Real-Time Responsiveness](#low-latency-real-time-responsiveness)
    - [Handling Variable Workloads and Scaling Flexibility](#handling-variable-workloads-and-scaling-flexibility)
    - [CPU-Intensive Workload Parallelisation](#cpu-intensive-workload-parallelisation)
2. [Key Considerations for Selecting Concurrency Patterns](#key-considerations-for-selecting-concurrency-patterns)
3. [Concurrent Design Patterns Examples in Go](#concurrent-design-patterns-examples-in-go)
     - [Future](#future)
    - [Pipeline](#pipeline)
    - [Fan-Out/Fan-In](#fan-outfan-in)
    - [Worker Pool](#worker-pool)
    - [Dynamic Rate-Limited Worker Pool](#dynamic-rate-limited-worker-pool)
    - [Publish-Subscribe (Pub/Sub)](#publish-subscribe-pubsub)
4. [Comparison Table](#comparison-table)

---

## **Why Use Concurrency and Parallelism?**

### Improved Throughput and Resource Utilisation

- Concurrency helps maximise throughput by allowing multiple tasks to be processed simultaneously, ensuring better
  resource utilisation.
- Examples:
    - **Data Processing Pipelines**: Handling streams of data for ETL (Extract, Transform, Load) processes.
    - **Web Scraping**: Collecting data from multiple sources concurrently.

### Low-Latency, Real-Time Responsiveness

- Concurrency is vital for applications requiring quick responses to multiple incoming events or requests, such as APIs
  and sensor systems.
- Examples:
    - **Real-Time APIs**: Handling multiple requests concurrently.
    - **Sensor Data Processing**: Reacting instantly to incoming data from sensors.

### Handling Variable Workloads and Scaling Flexibility

- Concurrency allows for more efficient scaling, especially when workloads are unpredictable or subject to sudden
  spikes.
- Examples:
    - **Microservices Architecture**: Managing requests across a distributed system.
    - **Background Jobs and Task Queues**: Processing asynchronous tasks concurrently.

### CPU-Intensive Workload Parallelisation

- Parallelism is useful for dividing CPU-bound tasks across multiple cores to increase processing speed.
- Examples:
    - **Image Processing**: Applying transformations to images in parallel.
    - **Machine Learning Training**: Distributing model training to multiple cores.

## Key Considerations for Selecting Concurrency Patterns

To choose the appropriate design pattern for concurrency or parallelism, consider the following constraints:

- **Workload Size**:
    - Is it bounded or unbounded? Knowing whether your workload has clear boundaries will determine the approach.
- **Latency vs Throughput**:
    - Does the use case prioritise low latency (quick response) or high throughput (large volumes)? Different patterns
      will have different impacts on these factors.
- **Scalability**:
    - How much will the workload grow? This determines if the solution should be auto-scalable or if it needs to handle
      variable loads efficiently.
- **Resource Constraints**:
    - Are there limits on memory, CPU, or bandwidth? Understanding resource constraints helps to decide the number of
      concurrent operations.

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
    - Each stage processes data concurrently, maximizing resource utilisation.
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
    - Keeps workers busy, optimizing resource utilisation.
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

| Design Pattern                   | Input Type | Process Duration | Synchronisation                   | Latency                                                    | Throughput                                            | Data Flow                        | Use Case and Application Examples                                     |
|----------------------------------|------------|------------------|-----------------------------------|------------------------------------------------------------|-------------------------------------------------------|----------------------------------|-----------------------------------------------------------------------|
| Future                           | Single     | Short            | Blocking until result ready       | <span style="color:red">Potential Increased Latency</span> | Standard Throughput                                   | Request -> Computation -> Result | Async Computations, Async API Calls                                   |
| Pipeline                         | Unbounded  | Long             | Sequential Execution              | <span style="color:red">Sequential Latency</span>          | <span style="color:red">Sequential Throughput</span>  | Stage-wise Processing            | Stream Processing, Data Transformation Pipelines                      |
| Fan-out Fan-in                   | Bounded    | Short to Long    | Goroutine synchronisation         | <span style="color:green">Reduced Latency</span>           | <span style="color:green">Increased Throughput</span> | Task -> Worker -> Aggregator     | CPU bound parallel tasks, Data Processing, Image Processing           |
| Worker Pool                      | Unbounded  | Short to Long    | Worker Coordination               | <span style="color:green">Reduced Latency</span>           | <span style="color:green">Increased Throughput</span> | Task -> Worker -> Result         | I/O or CPU Bound Tasks, Task Processing Systems                       |
| Dynamic Rate-Limited Worker Pool | Unbounded  | Long             | Rate Limiter, Worker Coordination | Controlled Latency                                         | Controlled Throughput                                 | Task -> Worker -> Result         | External Rate Limits, Resource Management, API Clients, Microservices |
| Pub-Sub                          | Unbounded  | Long             | Topic-based Subscription          | Event Delivery Latency                                     | Varied Based on Subscribers                           | Event Broadcast                  | Event Broadcasting, Event Notification Systems                        |

---
