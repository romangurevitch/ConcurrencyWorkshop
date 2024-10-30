# Understanding the Publish/Subscribe (Pub/Sub) Pattern in Go

The Publish/Subscribe (Pub/Sub) pattern is a messaging paradigm where senders (publishers) distribute messages without needing to know the recipients (subscribers), and subscribers receive messages without knowing the senders.  
This decoupling of publishers and subscribers allows for a flexible, scalable, and maintainable architecture, especially in event-driven systems.

This guide will explain how to implement and use the Pub/Sub pattern in Go, focusing on practical aspects, common issues, and best practices.  
We'll walk through a step-by-step implementation and demonstrate how to integrate it into your projects.

---

## Table of Contents

1. [Introduction](#introduction)
2. [Implementation Example](#implementation-example)
3. [How to Use the Pub/Sub Implementation](#how-to-use-the-pubsub-implementation)
4. [Common Issues and Pitfalls](#common-issues-and-pitfalls)
5. [Best Practices](#best-practices)

---

## Introduction

In Go, the Pub/Sub pattern can be implemented using channels and goroutines, leveraging Go's concurrency primitives to build an efficient messaging system.  
Publishers send messages to a topic, and all subscribers to that topic receive the messages asynchronously.

![Pub/Sub Diagram](../../../docs/images/pubsub_graph.png)

This pattern is particularly beneficial when dealing with:

- **Event-Driven Architectures**: Decoupling components that produce events from those that consume them.
- **Real-Time Systems**: Distributing real-time updates to multiple consumers.
- **Scalable Systems**: Allowing easy scaling of publishers and subscribers independently.

---

## Implementation Example

See [main.go](main.go)

The implementation uses a `PubSub` struct to manage topics and their subscribers.  
Subscribers can subscribe to topics, and publishers can publish messages to topics without knowledge of who the subscribers are.

**Key Components:**

- **`PubSub[T any]`**: Manages topics and subscribers, providing methods to subscribe, unsubscribe, and publish messages.
- **`Subscribe`**: Allows subscribers to subscribe to a specific topic.
- **`Unsubscribe`**: Allows subscribers to unsubscribe from a specific topic.
- **`Publish`**: Sends messages to all subscribers of a specific topic.
- **`Result[T any]`**: Encapsulates a value and an error, used for passing messages and errors.

---

## How to Use the Pub/Sub Implementation

### Step 1: Initialize the Pub/Sub System

Create an instance of the `PubSub` struct.

```go
pubSub := NewPubSub[T]()
```

Replace `T` with the type of data you want to publish (e.g., `string`, `int`, or a custom struct).

### Step 2: Subscribers Subscribe to Topics

Subscribers create channels to receive messages and subscribe to a topic.

```go
topicName := "your_topic_name"
subscriberCh := make(chan Result[T], bufferSize) // bufferSize determines the channel capacity

pubSub.Subscribe(topicName, subscriberCh)
```

### Step 3: Publishers Publish Messages to Topics

Publishers send messages to a topic without needing to know who the subscribers are.

```go
message := /* your message of type T */
pubSub.Publish(topicName, message)
```

### Step 4: Subscribers Receive Messages

Subscribers read messages from their channels.

```go
go func() {
    for result := range subscriberCh {
        if result.Err != nil {
            // Handle error
            continue
        }
        // Process result.Value
    }
}()
```

### Step 5: Unsubscribe When Done

Subscribers should unsubscribe when they no longer need to receive messages to prevent memory leaks.

```go
pubSub.Unsubscribe(topicName, subscriberCh)
close(subscriberCh) // Close the channel if no longer needed
```

---

## Common Issues and Pitfalls

### 1. Message Loss

**Issue**: Messages may be lost if there are no subscribers at the time of publishing or if subscriber channels are full.

**Solution**:

- **Ensure Subscribers Are Active**: Start subscribers before publishing messages.
- **Use Buffered Channels**: Use buffered channels for subscribers to prevent blocking publishers if subscribers are slow.
- **Check Channel Capacity**: Monitor and adjust channel buffer sizes based on expected message throughput.

### 2. Slow Subscribers Blocking Publishers

**Issue**: If a subscriber is slow and its channel is full, it can block the publisher if the publish operation waits indefinitely.

**Solution**:

- **Non-Blocking Sends**: In the implementation, the `Publish` method uses a non-blocking send with `select` and `default` to avoid blocking.
- **Handle Slow Subscribers**: Consider mechanisms to handle slow subscribers, such as dropping messages or implementing backpressure.

### 3. Concurrent Access to Subscribers Map

**Issue**: Concurrent access to the subscribers map without proper synchronization can lead to race conditions.

**Solution**:

- **Synchronization**: Use synchronization primitives like `sync.Map` to handle concurrent access safely.

### 4. Memory Leaks Due to Unsubscribed Channels

**Issue**: If subscribers do not unsubscribe when done, the Pub/Sub system may retain references to their channels, leading to memory leaks.

**Solution**:

- **Unsubscribe When Done**: Always call `Unsubscribe` when a subscriber no longer needs to receive messages.

---

## Best Practices

### 1. Use Contexts for Cancellation

- **Graceful Shutdown**: Use `context.Context` to manage the lifecycle of subscribers and publishers, allowing for graceful shutdown.

### 2. Proper Error Handling

- **Capture and Handle Errors**: Use the `Result[T]` type to encapsulate messages and errors.  
Ensure that errors are checked and handled appropriately by subscribers.

### 3. Buffer Subscriber Channels Appropriately

- **Adjust Buffer Sizes**: Set channel buffer sizes based on expected message rates and subscriber processing speeds to prevent message loss or blocking.

### 4. Avoid Blocking Operations in Publishers

- **Non-Blocking Publish**: Ensure that the `Publish` method does not block indefinitely due to slow subscribers.  
Use non-blocking sends or implement timeouts.

### 5. Clean Up Resources

- **Unsubscribe and Close Channels**: Subscribers should unsubscribe and close their channels when no longer needed to free up resources.

### 6. Topic Management

- **Topic Naming Conventions**: Use consistent and clear naming conventions for topics to avoid confusion.
- **Topic Existence Checks**: Implement checks or methods to manage topics, such as listing existing topics or deleting unused ones.

### 7. Scalability Considerations

- **Sharding and Partitioning**: For high-throughput systems, consider sharding topics or partitioning subscribers to distribute the load.
- **Concurrency Control**: Ensure that the Pub/Sub system can handle concurrent publish and subscribe operations efficiently.

---
