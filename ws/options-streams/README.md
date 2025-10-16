# Binance WebSocket Streams SDK (Go)

Generated Go SDK for Binance WebSocket Streams modules (spot, futures, options, etc.). It supports connecting to single and combined streams, dynamic subscription management, typed event handlers, and helpers to build stream names from the spec-defined patterns.

Module: github.com/openxapi/binance-go/ws/options-streams
Version: 0.1.0

## Features

- Single and combined connections (/ws/{streamName}, /stream)
- Subscribe/Unsubscribe/ListSubscriptions over WebSocket
- Typed handler registration per event
- Combined-stream wrapper routing
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

### 2) Connect (single, combined, user data)

Connect using generated channels (per AsyncAPI channel):

- Market Streams (single): /ws/{streamName}
- Combined Market Streams: /stream?streams={streams}
- User Data Streams: /ws/{listenKey}

```go
// Single stream channel
ch := ws.NewMarketStreamsChannel(client)

// Register handler(s) before connect (recommended)
ch.HandleNewSymbolInfoEvent(func(ctx context.Context, ev *wsmodels.NewSymbolInfoEvent) error {
  log.Printf("event: %+v", ev)
  return nil
})

// Connect (arguments depend on channel template)
if err := ch.Connect(ctx, "ETH@openInterest@221125"); err != nil {
  log.Fatalf("connect failed: %v", err)
}

// Combined connection channel (connect first, then SUBSCRIBE)
comb := ws.NewCombinedMarketStreamsChannel(client)
if err := comb.Connect(ctx, "ETH@openInterest@221125/ETH@markPrice"); err != nil {
  log.Fatalf("connect combined failed: %v", err)
}

+// User Data Streams channel (requires listenKey)
uds := ws.NewUserDataStreamsChannel(client)
// Example handler
uds.HandleAccountUpdateEvent(func(ctx context.Context, ev *wsmodels.AccountUpdateEvent) error {
  log.Printf("account update: %+v", ev)
  return nil
})
// Connect with a valid listenKey (from REST userDataStream.start)
if err := uds.Connect(ctx, "<listenKey>"); err != nil {
  log.Fatalf("connect user data failed: %v", err)
}
```

### 3) Build stream names and subscribe

Use generated builders from stream patterns and speeds. Patterns, examples, and speeds are arrays when provided in the spec.

```go
streams := []string{}

// Example: PartialDepthEvent has patterns like "{symbol}@depth{levels}" and "{symbol}@depth{levels}@{speed}"
if s, err := ws.BuildPartialDepthEventStream(1, map[string]string{
  "symbol": "BTC-210630-9000-P",
  "levels": "10",
  "speed":  "100ms",
}); err == nil {
  streams = append(streams, s)
}

// Or get all satisfiable variants for a given set of values
if many, err := ws.BuildTradeEventStreams(map[string]string{"symbol": "BTC-210630-9000-P"}); err == nil {
  streams = append(streams, many...)
}

// Subscribe on the active connection
if err := client.Subscribe(ctx, streams); err != nil {
  log.Fatalf("subscribe failed: %v", err)
}
```

### 4) One-shot request/response (example)

Most control actions use one-shot request/response via an id field. The SDK registers a one-time reply handler by id.

```go
// Build request (method const set automatically by the SDK)
req := &wsmodels.ListSubscriptionsRequest{ Id: 1 }

// Define reply handler
onReply := func(ctx context.Context, res *wsmodels.ListSubscriptionsResponse) error {
  log.Printf("active subscriptions: %+v", res.Result)
  return nil
}

// Send on combined channel (similar methods exist for single channel)
if err := comb.ListSubscriptionsFromCombinedMarketStreams(ctx, req, &onReply); err != nil {
  log.Fatalf("list subscriptions failed: %v", err)
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
  - If the spec provides `x-stream-params`, typed param helpers are also generated:
    - `type <Event>StreamParams struct { ... }`
    - `func (p <Event>StreamParams) Values() map[string]string` to convert to placeholder values

Example (typed params and speeds where available):

```go
// Build with speed as a regular param
s, err := ws.BuildPartialDepthEventStream(
  1,
  map[string]string{"symbol": "BTC-210630-9000-P", "levels": "10", "speed": "100ms"},
)
// Or use typed params, then convert to values
ps := ws.PartialDepthEventStreamParams{
  Symbol: "BTC-210630-9000-P",
  Levels: wsmodels.DepthLevels10,
  Speed:  wsmodels.DepthSpeed100ms, // speed is just another param when defined in x-stream-params
}
s2, err := ws.BuildPartialDepthEventStream(1, ps.Values())

// More typed params examples (types are generated from x-stream-params):
// - KlineEvent: Symbol (string), Interval (models.Interval)
ks := ws.KlineEventStreamParams{
  Symbol:   "BTC-200630-9000-P",
  Interval: wsmodels.Interval("1m"),
}
ksVals := ks.Values()
_ = ksVals

// - MarkPriceEvent: UnderlyingAsset (models.UnderlyingAsset)
mps := ws.MarkPriceEventStreamParams{UnderlyingAsset: wsmodels.UnderlyingAsset("ETH")}
_ = mps
```

## Handler Registration

- Register per-channel handlers using `Handle<Event>(func(ctx, *models.Event) error)`.
- Combined stream wrappers are routed via alias keys; unwrapped events are dispatched by event type.
- RegisterHandlers replaces the handler map for a channel key (subsequent calls overwrite the previous map for that channel).

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

## Selected Event Metadata

Below are selected patterns/examples/speeds from this spec. Use them with the generated builders:

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

## Notes

- For combined connections, connect without `streams` and use `SUBSCRIBE`/`UNSUBSCRIBE` to manage streams.
- User Data Streams use a dedicated connection (`/ws/{listenKey}`) and are not mixed into combined market streams.
- The SDK normalizes empty query parameters for connect paths (avoids "/stream?streams=").
- Servers from the spec are preloaded and first is active; override only if needed.
- Handlers should be registered before connect to avoid missing early messages.

## License

MIT

