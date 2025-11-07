# Binance WebSocket Streams SDK (Go)

Generated Go SDK for Binance WebSocket modules (spot, market streams, user data, etc.). It supports channel helpers from the AsyncAPI spec, typed events, and request/response helpers.

Module: github.com/openxapi/binance-go/ws/options-streams
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

go get github.com/openxapi/binance-go/ws/options-streams

## Quick Start

### 1) Create client (server preloaded)

```go
import (
  "context"
  "log"
  ws "github.com/openxapi/binance-go/ws/options-streams"
  wsmodels "github.com/openxapi/binance-go/ws/options-streams/models"
)

func main() {
  client := ws.NewClient()
  // Servers from the AsyncAPI spec are added automatically.
  // The first server is active by default — no setup required.
  // Optional: override or add another server
  // _ = client.AddOrUpdateServer("alt", "wss://nbstream.binance.com/eoptions", "Alt Server", "Optional override")
  // _ = client.SetActiveServer("alt")

  ctx := context.Background()
  // ...
}
```

### 2) Connect

Connect using generated channels:

- Market Streams (single): /ws/{streamName}
- Combined Market Streams: /stream?streams={streams}
- User Data Streams: /ws/{listenKey}

```go
// Single Market Streams Connection channel
ch := ws.NewMarketStreamChannel(client)
// Register handler(s) before connect (recommended)
ch.HandleNewSymbolInfoEvent(func(ctx context.Context, ev *wsmodels.NewSymbolInfoEvent) error {
  log.Printf("event: %+v", ev)
  return nil
})
if err := ch.Connect(ctx, "ETH@openInterest@221125"); err != nil {
  log.Fatalf("connect failed: %v", err)
}

// Combined Market Streams Connection channel
comb := ws.NewCombinedMarketStreamChannel(client)
if err := comb.Connect(ctx, "ETH@openInterest@221125/ETH@markPrice"); err != nil {
  log.Fatalf("connect combined failed: %v", err)
}

// User Data Streams Connection channel (requires listenKey)
uds := ws.NewUserDataStreamChannel(client)
if err := uds.Connect(ctx, "<listenKey>"); err != nil {
  log.Fatalf("connect user data failed: %v", err)
}
```

### 3) Build stream names and subscribe

Use the generated stream builders derived from the AsyncAPI patterns:

```go
streams := []string{}

// Example using NewSymbolInfoEvent
if s, err := ws.BuildNewSymbolInfoEventStream(0, map[string]string{
  // No placeholder params required for this stream
}); err == nil {
  streams = append(streams, s)
}

// Generate every permutation from supplied values
if many, err := ws.BuildNewSymbolInfoEventStreams(map[string]string{
  // No placeholder params required for this stream
}); err == nil {
  streams = append(streams, many...)
}

subReq := &wsmodels.SubscribeRequest{
  Id:     wsmodels.NewMessageIDInt64(1),
  Params: streams,
}
if err := comb.CombinedMarketStreamSubscribe(ctx, subReq, nil); err != nil {
  log.Fatalf("subscribe failed: %v", err)
}
```


### 4) One-shot request/response (example)

```go
req := &wsmodels.ListSubscriptionsRequest{
  // Populate request fields here
}
onReply := func(ctx context.Context, res *wsmodels.ListSubscriptionsResponse, wsErr error) error {
  if wsErr != nil {
    if apiErr, ok := wsErr.(*wsmodels.ErrorMessage); ok {
      log.Printf("request failed: %s", apiErr.Error())
    }
    return wsErr
  }
  log.Printf("reply: %+v", res)
  return nil
}
if err := comb.CombinedMarketStreamListSubscriptions(ctx, req, &onReply); err != nil {
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


## Selected Event Metadata

- NewSymbolInfoEvent
  Patterns:
  - option_pair
  Examples:
  - (none)
  Update Speeds:
  - 50ms

- OpenInterestEvent
  Patterns:
  - {underlyingAsset}@openInterest@{expirationDate}
  Examples:
  - ETH@openInterest@221125
  Update Speeds:
  - 60s

- MarkPriceEvent
  Patterns:
  - {underlyingAsset}@markPrice
  Examples:
  - ETH@markPrice
  Update Speeds:
  - 1000ms

- KlineEvent
  Patterns:
  - {symbol}@kline_{interval}
  Examples:
  - BTC-200630-9000-P@kline_1m
  Update Speeds:
  - 1000ms

- TickerByUnderlyingEvent
  Patterns:
  - {underlyingAsset}@ticker@{expirationDate}
  Examples:
  - ETH@ticker@220930
  Update Speeds:
  - 1000ms

- IndexPriceEvent
  Patterns:
  - {symbol}@index
  Examples:
  - ETHUSDT@index
  Update Speeds:
  - 1000ms

- TickerEvent
  Patterns:
  - {symbol}@ticker
  Examples:
  - BTC-210630-9000-P@ticker
  Update Speeds:
  - 1000ms

- TradeEvent
  Patterns:
  - {symbol}@trade
  - {underlyingAsset}@trade
  Examples:
  - BTC-210630-9000-P@trade
  - ETH@trade
  Update Speeds:
  - 50ms



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

