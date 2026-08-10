# 🚀 GoLang Journey - Day 36

---

# 🏗️ Theme

Today marked another major milestone in my Go backend journey as I learned **Redis Pub/Sub** and how Redis can be used not only as a caching system but also as a lightweight messaging mechanism for event-driven backend architectures.

Until today, I primarily viewed Redis as an in-memory data store used for caching and specialized data structures. Today I learned how applications can use Redis to publish events to channels and allow other components to subscribe to those events without directly communicating with each other.

The biggest realization of the day was understanding the difference between **live message broadcasting** and **durable event processing**.

I learned that Redis Pub/Sub is excellent for real-time and ephemeral communication, while Redis Streams are more appropriate when events need persistence, acknowledgements, consumer groups, replay, and reliable processing.

By the end of the day, I fully understood publishers, subscribers, channels, Redis Pub/Sub in Go, JSON event publishing, long-running subscribers, goroutines, Pub/Sub connections, real-time notifications, Pub/Sub limitations, and the architectural differences between Redis Pub/Sub and Redis Streams.

---

# 🎯 Goal of the Day

Today's goal was to understand how Redis can be used as a messaging system and how independent backend components can communicate through events.

The objective was to master Redis Pub/Sub by learning publishers, subscribers, channels, message publishing, message consumption, JSON event serialization, subscriber lifecycle, Redis connection behavior, real-time notification architecture, and Pub/Sub limitations.

The final objective was to understand when Redis Pub/Sub should be used and when Redis Streams are a better architectural choice.

---

# 📚 Topics Covered

## Redis Pub/Sub

Studied the Publish/Subscribe messaging model.

Covered:

- Redis Pub/Sub
- Publisher
- Subscriber
- Channels
- Message Broadcasting
- Event-Driven Architecture
- Decoupled Services

Learned how publishers send messages to Redis channels while subscribers listen for messages without the publisher needing to know who consumes them.

---

## Publishers

Studied how applications publish events.

Covered:

- Publisher
- `Publish()`
- Redis Channels
- Message Payload
- Event Publishing
- Error Handling

Learned how Go applications can publish messages using the Redis client and how the publisher communicates with a channel rather than directly with another service.

---

## Redis Channels

Studied the role of channels in Pub/Sub.

Covered:

- Channels
- Channel Names
- Event Types
- `order.created`
- Message Routing
- Multiple Subscribers

Learned that channels act as communication paths through which publishers broadcast events to interested subscribers.

---

## Subscribers

Studied how applications consume Pub/Sub messages.

Covered:

- Subscriber
- `Subscribe()`
- `ReceiveMessage()`
- Message Payload
- Long-Running Processes
- Continuous Message Processing

Learned how a subscriber continuously waits for messages and processes incoming events.

---

## Redis Pub/Sub in Go

Implemented the basic Publisher and Subscriber concepts using Go.

Covered:

- `redis.Client`
- `Publish()`
- `Subscribe()`
- `ReceiveMessage()`
- `context.Context`
- Error Handling
- Pub/Sub Lifecycle

Learned how Go communicates with Redis Pub/Sub through the Redis client and how subscribers can continuously listen for events.

---

## JSON Event Architecture

Studied how structured events can be communicated between services.

Covered:

- Event Structures
- JSON Serialization
- `json.Marshal()`
- `json.Unmarshal()`
- Event Payloads
- Event Contracts

Learned how Go structs can be serialized into JSON before publishing and reconstructed by subscribers after receiving the event.

---

## Goroutines and Subscribers

Connected Redis Pub/Sub with Go concurrency.

Covered:

- Goroutines
- Blocking Operations
- Long-Running Subscribers
- Concurrent Processing
- Application Lifecycle

Learned why a subscriber can run inside a goroutine while the main Go application continues handling HTTP requests and other work.

---

## Redis Pub/Sub Connections

Studied how Pub/Sub differs from normal Redis operations.

Covered:

- Redis Client
- Redis Connections
- Connection Pooling
- Pub/Sub Subscription
- Long-Lived Connections
- Request-Response Operations

Learned that normal Redis commands such as `GET` and `SET` follow a request-response model, while Pub/Sub subscriptions remain active and continuously receive messages.

---

## Real-Time Notification Architecture

Built the conceptual architecture for event-driven notifications.

Covered:

- Event Publishing
- Notification Service
- Order Events
- Event Consumers
- Service Decoupling
- Real-Time Notifications

Learned how an order service can publish an `order.created` event while independent notification, analytics, or other services consume the event.

---

## Pub/Sub Limitations

Studied the limitations of Redis Pub/Sub.

Covered:

- Message Loss
- No Message History
- No Replay
- No Durable Acknowledgements
- Subscriber Disconnects
- Slow Consumers
- Delivery Guarantees

Learned that basic Redis Pub/Sub is designed for live communication and does not provide the durable event-processing guarantees required by many critical backend workflows.

---

## Redis Streams

Introduced Redis Streams as a more durable event-processing mechanism.

Covered:

- Redis Streams
- Persistent Events
- Event History
- Consumer Groups
- Acknowledgements
- Replay
- Work Distribution

Learned why Redis Streams are better suited for durable event processing where messages need to survive consumer downtime and be processed reliably.

---

## Pub/Sub vs Redis Streams

Compared the two messaging models.

Covered:

- Broadcast vs Work Distribution
- Ephemeral vs Durable Events
- Message Persistence
- Consumer Groups
- Acknowledgements
- Replay
- Event Processing
- Architectural Trade-offs

Learned how to decide between Redis Pub/Sub and Redis Streams based on application requirements and delivery guarantees.

---

# 💻 Concepts Learned

- Redis Pub/Sub
- Publisher
- Subscriber
- Redis Channels
- `PUBLISH`
- `SUBSCRIBE`
- `Publish()`
- `Subscribe()`
- `ReceiveMessage()`
- Message Payload
- JSON Events
- `json.Marshal()`
- `json.Unmarshal()`
- Goroutines
- Context
- Redis Connections
- Connection Pooling
- Event-Driven Architecture
- Service Decoupling
- Real-Time Notifications
- Message Broadcasting
- Message Persistence
- Message Replay
- Acknowledgements
- Consumer Groups
- Redis Streams
- Delivery Guarantees
- Distributed Systems

---

# 🧠 Important Concepts Learned

- Redis Pub/Sub allows publishers to broadcast messages through channels.
- Publishers do not need to know which services consume their events.
- Subscribers listen to specific Redis channels.
- Multiple subscribers can receive the same Pub/Sub message.
- Pub/Sub is primarily designed for live and ephemeral communication.
- A basic Pub/Sub subscriber does not receive messages that were published while it was offline.
- Pub/Sub does not provide durable message history or replay.
- Pub/Sub does not provide the durable acknowledgement workflow needed for reliable event processing.
- Go goroutines allow long-running subscribers to operate concurrently with HTTP servers and other application components.
- JSON can be used to transport structured events between services.
- Redis Streams provide persistent event entries and support consumer groups and acknowledgements.
- Pub/Sub is better suited for live broadcasting while Streams are better suited for durable event processing.
- Messaging architecture should be selected based on delivery guarantees and business requirements.

---

# ⚠️ Common Mistakes I Learned

- Assuming Redis Pub/Sub stores messages permanently.
- Assuming offline subscribers receive missed messages after reconnecting.
- Treating Pub/Sub as a durable job queue.
- Creating subscriptions inside individual HTTP request handlers.
- Forgetting that subscribers are long-running processes.
- Ignoring errors returned by `Publish()`.
- Publishing unstructured messages when consumers require meaningful event data.
- Forgetting to serialize structured events before publishing.
- Forgetting to deserialize JSON payloads inside subscribers.
- Assuming Pub/Sub provides acknowledgements and retries.
- Using Pub/Sub for critical events that cannot safely be lost.
- Confusing Pub/Sub broadcasting with queue-based work distribution.
- Assuming Redis Streams are simply a better version of Pub/Sub instead of recognizing that they solve different problems.

---

# 🎯 Interview Notes

## What Is Redis Pub/Sub?

Redis Pub/Sub is a messaging mechanism where publishers send messages to channels and subscribers receive messages from those channels.

---

## What Is A Publisher?

A publisher is the application or component that sends a message to a Redis channel.

---

## What Is A Subscriber?

A subscriber is an application or component that listens to a Redis channel and receives messages published to it.

---

## What Is A Redis Channel?

A Redis channel is a named communication path through which publishers broadcast messages to subscribers.

---

## What Happens If A Subscriber Is Offline?

With basic Redis Pub/Sub, the subscriber misses messages published while it is disconnected. Those messages are not stored for later replay.

---

## Why Use Pub/Sub?

Pub/Sub is useful when applications need real-time communication and the loss of an individual event is acceptable.

Examples include:

- Live notifications
- Typing indicators
- Live dashboards
- Presence updates
- Real-time updates

---

## Why Isn't Pub/Sub Suitable For Every Backend Event?

Basic Pub/Sub does not provide durable message storage, replay, or the acknowledgement and recovery mechanisms needed for reliable processing of critical events.

---

## What Is Redis Streams?

Redis Streams provide a persistent event-log style data structure that supports event history, consumer groups, acknowledgements, and more reliable event processing.

---

## What Is The Difference Between Pub/Sub And Streams?

Pub/Sub is primarily designed for live broadcast communication, while Streams are designed for durable event processing and workload distribution.

---

## Why Are Consumer Groups Useful?

Consumer groups allow multiple consumers to share the processing of stream entries rather than having every consumer process every event.

---

## Why Are Acknowledgements Important?

Acknowledgements allow a consumer to indicate that an event has been successfully processed, which is important for recovery and reliable event-processing workflows.

---

# 🏛️ Architecture Reinforced Today

```text
                    Go Backend
                        │
                        ▼
                     Service
                        │
                        ▼
                  Event Publisher
                        │
                        │ PUBLISH
                        ▼
                ┌───────────────┐
                │     Redis     │
                │               │
                │ order.created │
                └───────┬───────┘
                        │
                        ▼
                  Subscriber
                        │
                        ▼
                Notification Logic
```
