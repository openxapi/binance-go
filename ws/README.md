# Go WebSocket SDK for Binance

[![Go Reference](https://pkg.go.dev/badge/github.com/openxapi/binance-go/ws.svg)](https://pkg.go.dev/github.com/openxapi/binance-go/ws)

Generated Binance WebSocket clients backed by the official AsyncAPI specification. The SDK is produced by [OpenXAPI](https://github.com/openxapi/openxapi); regenerate from that repository whenever the spec or Go templates change.

> **Do not edit the generated Go code manually.** Update the AsyncAPI spec or templates and rerun OpenXAPI instead.

## Key Capabilities

- AsyncAPI-driven channels for every documented Binance WebSocket API and stream.
- Typed request/reply helpers that surface server-side errors as `*models.ErrorMessage`.
- Event handler registration per channel plus stream-name builders for every pattern (`Build<Stream>`, `<Stream>StreamParams`).
- Shared `ServerManager`, proxy-friendly dialing, and configurable worker pools via `ClientOptions`.
- Multi-key authentication (HMAC, RSA, Ed25519) with helpers for API keys, PEM files, or custom readers.

## Packages

| Product | Type | Import Path | Notes |
|---------|------|-------------|-------|
| Spot | WebSocket API | `github.com/openxapi/binance-go/ws/spot` | Request/response API plus listen-key user data events. |
| Spot | Market Streams | `github.com/openxapi/binance-go/ws/spot-streams` | Single or combined streams for trades, klines, depth, tickers. |
| USDS-M Futures | WebSocket API | `github.com/openxapi/binance-go/ws/umfutures` | Trading & account RPCs for USDS-M futures. |
| USDS-M Futures | Market Streams | `github.com/openxapi/binance-go/ws/umfutures-streams` | Trades, mark price, depth, liquidation feeds. |
| COIN-M Futures | WebSocket API | `github.com/openxapi/binance-go/ws/cmfutures` | Trading & account RPCs for COIN-M futures. |
| COIN-M Futures | Market Streams | `github.com/openxapi/binance-go/ws/cmfutures-streams` | COIN-M depth, liquidation, mark price streams. |
| Options | Market Streams | `github.com/openxapi/binance-go/ws/options-streams` | Options trades, Greeks, user data streams. |
| Portfolio Margin (Classic) | User Data Streams | `github.com/openxapi/binance-go/ws/pmargin-streams` | Listen-key feeds for Portfolio Margin accounts. |
| Portfolio Margin Pro | User Data Streams | `github.com/openxapi/binance-go/ws/pmarginpro-streams` | Listen-key feeds for Portfolio Margin Pro accounts. |

All packages live inside the root module (`github.com/openxapi/binance-go/ws`) and can be imported independently.

## Installation

Install the root module or only the package you need:

```bash
go get github.com/openxapi/binance-go/ws

# Common packages
go get github.com/openxapi/binance-go/ws/spot
go get github.com/openxapi/binance-go/ws/spot-streams
go get github.com/openxapi/binance-go/ws/umfutures
go get github.com/openxapi/binance-go/ws/umfutures-streams
go get github.com/openxapi/binance-go/ws/cmfutures
go get github.com/openxapi/binance-go/ws/cmfutures-streams
go get github.com/openxapi/binance-go/ws/options-streams
go get github.com/openxapi/binance-go/ws/pmargin-streams
go get github.com/openxapi/binance-go/ws/pmarginpro-streams
```

## Usage

### Spot WebSocket API (request/response + user data)

```go
package main

import (
    "context"
    "log"
    "os"
    "time"

    spot "github.com/openxapi/binance-go/ws/spot"
    spotmodels "github.com/openxapi/binance-go/ws/spot/models"
)

func main() {
    client := spot.NewClient()
    _ = client.SetActiveServer("testnet1") // or mainnet1/mainnet2

    if apiKey, secret := os.Getenv("BINANCE_API_KEY"), os.Getenv("BINANCE_SECRET_KEY"); apiKey != "" && secret != "" {
        auth := spot.NewAuth(apiKey)
        auth.SetSecretKey(secret)
        client.SetAuth(auth)
    }

    ch := spot.NewSpotChannel(client)
    ch.HandleOrderUpdateEvent(func(ctx context.Context, ev *spotmodels.OrderUpdateEvent) error {
        log.Printf("order %d status %s", ev.Event.OrderID, ev.Event.CurrentOrderStatus)
        return nil
    })
    ch.HandleBalanceUpdateEvent(func(ctx context.Context, ev *spotmodels.BalanceUpdateEvent) error {
        log.Printf("balance %s delta %s", ev.Event.Asset, ev.Event.BalanceDelta)
        return nil
    })

    ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
    defer cancel()

    if err := ch.Connect(ctx); err != nil {
        log.Fatal(err)
    }
    defer ch.Disconnect(ctx)

    timeReq := &spotmodels.TimeRequest{Id: spotmodels.NewMessageIDInt64(1)}
    onTime := func(ctx context.Context, res *spotmodels.TimeResponse, wsErr error) error {
        if wsErr != nil {
            return wsErr
        }
        log.Printf("server time: %d", res.Result.ServerTime)
        return nil
    }
    if err := ch.Time(ctx, timeReq, &onTime); err != nil {
        log.Fatal(err)
    }

    accountReq := &spotmodels.AccountStatusRequest{Id: spotmodels.NewMessageIDInt64(2)}
    onAccount := func(ctx context.Context, res *spotmodels.AccountStatusResponse, wsErr error) error {
        if wsErr != nil {
            if apiErr, ok := wsErr.(*spotmodels.ErrorMessage); ok {
                log.Printf("request failed: %s", apiErr.Error())
            }
            return wsErr
        }
        log.Printf("account status: %+v", res.Result)
        return nil
    }
    _ = ch.AccountStatus(ctx, accountReq, &onAccount)

    _ = client.Wait(ctx) // blocks until the read loop exits or ctx cancels
}
```

### Spot Market Streams (single + combined)

```go
package main

import (
    "context"
    "log"
    "time"

    spotstreams "github.com/openxapi/binance-go/ws/spot-streams"
    spotstreammodels "github.com/openxapi/binance-go/ws/spot-streams/models"
)

func main() {
    client := spotstreams.NewClient()
    _ = client.SetActiveServer("testnet")

    combined := spotstreams.NewCombinedMarketStreamChannel(client)
    combined.HandleAggregateTradeEvent(func(ctx context.Context, ev *spotstreammodels.AggregateTradeEvent) error {
        log.Printf("agg trade %s price %s", ev.Symbol, ev.Price)
        return nil
    })
    combined.HandleMiniTickerEvent(func(ctx context.Context, ev *spotstreammodels.MiniTickerEvent) error {
        log.Printf("mini ticker %s close %s", ev.Symbol, ev.ClosePrice)
        return nil
    })

    ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
    defer cancel()

    if err := combined.Connect(ctx, "btcusdt@aggTrade/btcusdt@miniTicker"); err != nil {
        log.Fatal(err)
    }
    defer combined.Disconnect(ctx)

    streams := []string{}
    if s, err := spotstreams.BuildAggregateTradeEventStream(0, map[string]string{"symbol": "btcusdt"}); err == nil {
        streams = append(streams, s)
    }
    if s, err := spotstreams.BuildMiniTickerEventStream(0, map[string]string{"symbol": "btcusdt"}); err == nil {
        streams = append(streams, s)
    }

    subReq := &spotstreammodels.SubscribeRequest{
        Id:     spotstreammodels.NewMessageIDInt64(1),
        Params: streams,
    }
    if err := combined.Subscribe(ctx, subReq, nil); err != nil {
        log.Fatal(err)
    }

    time.Sleep(15 * time.Second)

    unsubReq := &spotstreammodels.UnsubscribeRequest{
        Id:     spotstreammodels.NewMessageIDInt64(2),
        Params: streams,
    }
    _ = combined.Unsubscribe(ctx, unsubReq, nil)
}
```

Portfolio Margin (`pmargin-streams`, `pmarginpro-streams`) packages expose `NewUserDataStreamChannel` – connect with a listen key (`/ws/{listenKey}`) and register handlers such as `HandleRiskLevelChangeEvent`. Futures stream packages mirror the Spot stream example above, while `cmfutures`/`umfutures` request APIs mirror the Spot API sample.

## Stream Builders & Events

For every event with defined patterns, the generator creates:

- `<Event>StreamPatterns`, `<Event>StreamExamples`, and optional `<Event>UpdateSpeeds` constants.
- `Build<Event>Stream(patternIndex, values)` and `Build<Event>Streams(values)` helpers.
- Optional typed parameter structs (for patterns with named placeholders).

Use these helpers whenever you need to build subscription names programmatically or validate supported combinations.

## Authentication & Configuration

```go
auth := spot.NewAuth(os.Getenv("BINANCE_API_KEY"))
auth.SetSecretKey(os.Getenv("BINANCE_SECRET_KEY")) // HMAC
// For RSA / Ed25519 keys
auth.SetPrivateKeyPath("/path/to/private_key.pem")
auth.SetPassphrase("optional-password")
client.SetAuth(auth)
```

- `ClientOptions{HandlerWorkers: N}` controls how many goroutines process handlers; default is `runtime.NumCPU()`.
- Every client ships with predefined servers (spot: `mainnet1`, `mainnet2`, `testnet1`, `testnet2`; streams: `mainnet`, `mainnetAlt`, `testnet`, etc.). Call `client.ListServers()` to inspect names or `AddOrUpdateServer` to point at custom gateways.
- `SetActiveServer` switches between mainnet and testnet endpoints without rebuilding the client.
- Use `HTTP_PROXY` / `HTTPS_PROXY` to route traffic through a proxy before dialing.
- Call `client.Wait(ctx)` to block until the background read loop stops, and `channel.Disconnect(ctx)` to tear down handlers cleanly.

## Testing

All packages include fast unit tests (mostly signature helpers). Run them before sending a PR:

```bash
go test ./...
# or target a package
go test ./spot-streams
```

Integration flows require real Binance keys and are not checked into this repository. When you add new AsyncAPI surface area, document any manual verification steps in your PR description.

## Contributing

1. Propose changes to the [AsyncAPI specs or Go templates](https://github.com/openxapi/openxapi).
2. Regenerate this module via OpenXAPI (see the template README for instructions).
3. Format with `gofmt` / `goimports`, run `go test ./...`, and open a PR describing the generated deltas (include commands used to regenerate).

## Support & References

- [Binance Spot WebSocket Docs](https://developers.binance.com/docs/binance-spot-api-docs/websocket-api)
- [USDS-M Futures WebSocket Docs](https://developers.binance.com/docs/derivatives/usds-margined-futures)
- [COIN-M Futures WebSocket Docs](https://developers.binance.com/docs/derivatives/coin-margined-futures)
- [Options Streams Docs](https://developers.binance.com/docs/derivatives/option/websocket-streams)

## License & Disclaimer

This project inherits the parent repository's license. It is an unofficial SDK; always confirm behavior against the official Binance documentation before executing trades in production.
