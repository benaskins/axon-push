# axon-push

Generic in-memory pub/sub and Server-Sent Events primitives.

Import: `github.com/benaskins/axon-push`

## What it does

axon-push provides two things:

1. **EventBus[T]**: a thread-safe, generic pub/sub that delivers events to subscriber channels with backpressure handling (drops events for slow subscribers)
2. **SSE helpers**: `SetSSEHeaders` and `SendEvent` for writing Server-Sent Events over HTTP

No external dependencies (stdlib only).

## Usage

```go
import push "github.com/benaskins/axon-push"

bus := push.NewEventBus[MyEvent]()

// Subscribe
ch := bus.Subscribe("client-1")

// Publish
bus.Publish(MyEvent{...})

// SSE endpoint
push.SetSSEHeaders(w)
push.SendEvent(w, flusher, data)
```

## Build & Test

```bash
go test ./...
go vet ./...
```
