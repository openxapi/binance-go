# Binance WebSocket Streams SDK (Go)

Generated Go SDK for Binance WebSocket modules (spot, market streams, user data, etc.). It supports channel helpers from the AsyncAPI spec, typed events, and request/response helpers.

Module: github.com/openxapi/binance-go/ws/cmfutures
Version: 0.1.0

## Features

- Channel helpers generated directly from the AsyncAPI specification
- Subscribe/Unsubscribe/ListSubscriptions over WebSocket (when defined)
- Typed handler registration per event
- Request/reply helpers surface server errors through the handler's `error` parameter
- Combined-stream wrapper routing (when combined streams are available)
- Stream-name builders from x-stream-pattern(s), examples, and speeds
- Server management (multiple endpoints, active server switching)

## Install

go get github.com/openxapi/binance-go/ws/cmfutures

## Quick Start

### 1) Create client (server preloaded)

```go
import (
  "context"
  "log"
  ws "github.com/openxapi/binance-go/ws/cmfutures"
  wsmodels "github.com/openxapi/binance-go/ws/cmfutures/models"
)

func main() {
  client := ws.NewClient()
  // Servers from the AsyncAPI spec are added automatically.
  // The first server is active by default — no setup required.
  // Optional: override or add another server
  // _ = client.AddOrUpdateServer("alt", "wss://ws-dapi.binance.com/ws-dapi/v1", "Alt Server", "Optional override")
  // _ = client.SetActiveServer("alt")

  ctx := context.Background()
  // ...
}
```

### 2) Connect

Connect using generated channels:

- Channel cmfutures: /

```go
// Channel cmfutures channel
ch := ws.NewCmfuturesChannel(client)
// Register handler(s) before connect (recommended)
// ch.Handle<EventName>(func(ctx context.Context, ev *wsmodels.EventName) error { return nil })
if err := ch.Connect(ctx); err != nil {
  log.Fatalf("connect failed: %v", err)
}
```

### 3) Send messages

This specification does not define stream-name patterns. Use the generated request helpers on `Cmfutures` (and other channels) to send operations directly. See the request/response example below for the typical flow.


### 4) One-shot request/response (example)

```go
req := &wsmodels.AccountBalanceRequest{
  // Populate request fields here
}
onReply := func(ctx context.Context, res *wsmodels.AccountBalanceResponse, wsErr error) error {
  if wsErr != nil {
    if apiErr, ok := wsErr.(*wsmodels.ErrorMessage); ok {
      log.Printf("request failed: %s", apiErr.Error())
    }
    return wsErr
  }
  log.Printf("reply: %+v", res)
  return nil
}
if err := ch.AccountBalance(ctx, req, &onReply); err != nil {
  log.Fatalf("request failed: %v", err)
}
```

### 5) Wait and shutdown

```go
// Block until read loop ends or context cancels
_ = client.Wait(ctx)

// Close and remove handlers for a channel
_ = ch.Disconnect(ctx)
```

## Stream Builders and Metadata

For each event with patterns, the SDK generates:

- `<Event>StreamPatterns` (`[]string`)
- `<Event>StreamExamples` (`[]string`)
- `<Event>UpdateSpeeds` (`[]string`, when present)
- Builders:
  - `Build<Event>Stream(patternIndex int, values map[string]string) (string, error)`
  - `Build<Event>Streams(values map[string]string) ([]string, error)`
  - If the spec provides `x-stream-params`, typed param helpers are also generated.




## Error Responses

Request/reply error payloads are decoded into `*wsmodels.ErrorMessage`, which implements `error`. Always check the handler's error argument before using the reply value.


## Performance & Dispatch

- Incoming frames are read on a dedicated loop and dispatched asynchronously to a worker pool.
- Slow user handlers do not block `ReadMessage()`; an unbounded in-memory queue buffers messages between the reader and workers.
- Defaults: workers = `runtime.NumCPU()`. Messages are never dropped.
- Customize workers via `NewClientWithOptions(&ws.ClientOptions{ HandlerWorkers: N })`.

```go
client := ws.NewClientWithOptions(&ws.ClientOptions{
  HandlerWorkers: 8,
})
```

## Server Management

The client carries a `ServerManager` to manage endpoints:

- `AddServer(name, url, title, description)`
- `AddOrUpdateServer(...)`, `UpdateServer(...)`, `RemoveServer(name)`
- `SetActiveServer(name)`, `GetActiveServer()`, `GetActiveServerURL()`

## Notes

- For combined connections (when defined), connect without `streams` and use `SUBSCRIBE`/`UNSUBSCRIBE` to manage streams.
- User Data Streams use a dedicated connection (`/ws/{listenKey}`) when present and are not mixed into combined market streams.
- The SDK normalizes empty query parameters for connect paths (avoids "/stream?streams=").
- Servers from the spec are preloaded and first is active; override only if needed.
- Handlers should be registered before connect to avoid missing early messages.

## License

MIT

