# Binance Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/openxapi/binance-go.svg)](https://pkg.go.dev/github.com/openxapi/binance-go)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Binance REST and WebSocket SDKs for Go, generated from the official OpenAPI and AsyncAPI specs via [OpenXAPI](https://github.com/openxapi/openxapi). The code under `rest/` and `ws/` is regenerated frequently—update specs or templates upstream and rerun OpenXAPI instead of editing the generated clients by hand.

## Modules at a Glance

| Module | Description | Key Packages |
|--------|-------------|--------------|
| `rest/` | OpenAPI-driven HTTP clients covering Spot, Margin, Futures, Options, and Portfolio Margin endpoints. | `rest/spot`, `rest/umfutures`, `rest/cmfutures`, `rest/options`, `rest/pmargin` |
| `ws/` | AsyncAPI-driven WebSocket clients for request/response channels and market/user data streams. | `ws/spot`, `ws/spot-streams`, `ws/umfutures`, `ws/umfutures-streams`, `ws/cmfutures`, `ws/cmfutures-streams`, `ws/options-streams`, `ws/pmargin-streams`, `ws/pmarginpro-streams` |

Each module ships with its own `go.mod`, README, and product-specific models; import only what you need (`go get github.com/openxapi/binance-go/ws/spot`, etc.).

## Installation

```bash
# REST root or specific products
go get github.com/openxapi/binance-go/rest
go get github.com/openxapi/binance-go/rest/spot
go get github.com/openxapi/binance-go/rest/umfutures

# WebSocket root or streams
go get github.com/openxapi/binance-go/ws
go get github.com/openxapi/binance-go/ws/spot
go get github.com/openxapi/binance-go/ws/spot-streams
go get github.com/openxapi/binance-go/ws/umfutures-streams
```

## Key Features

- **REST (`rest/`)**: Complete Binance API coverage, typed request builders, multi-auth (HMAC, RSA, Ed25519), rate-limit aware HTTP client, and reusable configuration helpers (`NewConfiguration`, `ContextServerIndex`).
- **WebSocket (`ws/`)**: AsyncAPI channel handlers, strongly typed event payloads, listen-key management, combined-stream builders, proxy-friendly dialers, pluggable worker pools, and unified auth helpers shared across packages.

## Quick Start

### REST Example

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/openxapi/binance-go/rest/spot"
)

func main() {
    cfg := spot.NewConfiguration()
    client := spot.NewAPIClient(cfg)
    ctx := context.Background()

    info, _, err := client.SpotAPI.GetExchangeInfoV3(ctx).Symbol("BTCUSDT").Execute()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("symbols: %d\n", len(info.Symbols))

    auth := spot.NewAuth(os.Getenv("BINANCE_API_KEY"))
    auth.SetSecretKey(os.Getenv("BINANCE_SECRET_KEY"))
    ctx, err = auth.ContextWithValue(ctx)
    if err != nil {
        log.Fatal(err)
    }

    account, _, err := client.SpotAPI.GetAccountV3(ctx).Timestamp(time.Now().UnixMilli()).Execute()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("maker commission:", account.MakerCommission)
}
```

### WebSocket Example

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
    client.SetActiveServer("testnet1") // or mainnet1/mainnet2

    if apiKey := os.Getenv("BINANCE_API_KEY"); apiKey != "" {
        auth := spot.NewAuth(apiKey)
        auth.SetSecretKey(os.Getenv("BINANCE_SECRET_KEY"))
        client.SetAuth(auth)
    }

    ch := spot.NewSpotChannel(client)
    ch.HandleOrderUpdateEvent(func(ctx context.Context, ev *spotmodels.OrderUpdateEvent) error {
        log.Printf("order %d -> %s", ev.Event.OrderID, ev.Event.CurrentOrderStatus)
        return nil
    })

    ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
    defer cancel()

    if err := ch.Connect(ctx); err != nil {
        log.Fatal(err)
    }
    defer ch.Disconnect(ctx)

    req := &spotmodels.TimeRequest{Id: spotmodels.NewMessageIDInt64(1)}
    handler := func(ctx context.Context, res *spotmodels.TimeResponse, wsErr error) error {
        if wsErr != nil {
            return wsErr
        }
        log.Printf("server time: %d", res.Result.ServerTime)
        return nil
    }
    if err := ch.Time(ctx, req, &handler); err != nil {
        log.Fatal(err)
    }
}
```

## Build, Test, and Development

- `cd rest && go build ./...` / `cd ws && go build ./...` – compile each module.
- `go test ./rest/...` and `go test ./ws/...` – run unit suites; add `-run Live` only for opt-in integration tests.
- `go vet ./...` – catch spec/template drift quickly.
- Regenerate code by updating OpenXAPI specs/templates and rerunning the generator; do not hand-edit the SDK.

## Authentication & Configuration

Use `spot.NewAuth` (or equivalent per package) for API keys. `SetSecretKey` accepts HMAC secrets; providing a PEM path enables RSA/Ed25519 signing automatically. Set `ContextServerIndex`, `SetActiveServer`, or override `Configuration.Servers` to switch between production, testnet, or custom endpoints. Proxies can be configured via `HTTP_PROXY` or by supplying a custom `http.Client`.

## Rate Limits & Reliability

- REST clients honor Binance weight limits, but long-running scripts should still back off on `429/418`.
- WebSocket clients must respect Binance’s limit of 5 outgoing messages per second per connection; throttle manual sends and reconnect with jitter.
- Listen keys expire—renew them periodically via the provided request helpers or event handler callbacks.

## Contributing

Report bugs or request endpoints here, but submit code generation updates to [OpenXAPI](https://github.com/openxapi/openxapi) (specs, templates, or generator logic). Include the generator command, spec revision, and `go test` output in PRs so reviewers can reproduce the generated artifacts. Do not reformat generated files manually.

## Further Reading

- REST details: `rest/README.md`
- WebSocket & stream details: `ws/README.md`
- Contributor playbook: `AGENTS.md`
