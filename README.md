# Binance Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/openxapi/binance-go.svg)](https://pkg.go.dev/github.com/openxapi/binance-go)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Comprehensive Go SDK for Binance cryptocurrency exchange APIs, supporting both REST and WebSocket APIs across all Binance products.

## 🚀 Features

- **Complete API Coverage**: REST and WebSocket APIs for all Binance products
- **Type Safety**: Fully typed Go structs for all API responses
- **Multiple Authentication**: HMAC-SHA256, RSA, and Ed25519 support
- **Production Ready**: Comprehensive error handling and rate limiting
- **Well Tested**: Extensive integration tests with real API endpoints
- **Auto-Generated**: Generated from official OpenAPI/AsyncAPI specifications

## 📦 Installation

```bash
go get github.com/openxapi/binance-go
```

## 🏗️ Project Structure

This SDK is organized into two main components:

### REST APIs (`rest/`)
- **OpenAPI-generated** clients for HTTP REST endpoints
- Synchronous request-response operations
- Complete coverage of all Binance REST APIs

### WebSocket APIs (`ws/`)
- **AsyncAPI-generated** clients for WebSocket connections
- Real-time data streams and async operations
- Event-driven architecture with handlers

## 🛠️ Supported Products

| Product | REST API | WebSocket API | WebSocket Streams |
|---------|----------|---------------|-------------------|
| **Spot** | ✅ [`rest/spot`](rest/) | ✅ [`ws/spot`](ws/) | ✅ [`ws/spot-streams`](ws/) |
| **USDS-M Futures** | ✅ [`rest/umfutures`](rest/) | ✅ [`ws/umfutures`](ws/) | ✅ [`ws/umfutures-streams`](ws/) |
| **COIN-M Futures** | ✅ [`rest/cmfutures`](rest/) | ✅ [`ws/cmfutures`](ws/) | ✅ [`ws/cmfutures-streams`](ws/) |
| **Options** | ✅ [`rest/options`](rest/) | - | - |
| **Portfolio Margin** | ✅ [`rest/pmargin`](rest/) | - | - |

## 🚀 Quick Start

### REST API Example

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
    // Create client
    config := spot.NewConfiguration()
    client := spot.NewAPIClient(config)
    ctx := context.Background()

    // Public endpoint - no authentication required
    info, _, err := client.SpotTradingAPI.GetExchangeInfoV3(ctx).Execute()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Exchange info: %+v\n", info)

    // Private endpoint - requires authentication
    apiKey := os.Getenv("BINANCE_API_KEY")
    secretKey := os.Getenv("BINANCE_SECRET_KEY")

    if apiKey != "" && secretKey != "" {
        auth := spot.NewAuth(apiKey)
        auth.SetSecretKey(secretKey)
        ctx, err = auth.ContextWithValue(ctx)
        if err != nil {
            log.Fatal(err)
        }

        account, _, err := client.SpotTradingAPI.GetAccountV3(ctx).
            Timestamp(time.Now().UnixMilli()).Execute()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Account: %+v\n", account)
    }
}
```

### WebSocket API Example

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    spotws "github.com/openxapi/binance-go/ws/spot"
    "github.com/openxapi/binance-go/ws/spot/models"
)

func main() {
    // Create WebSocket client
    client := spotws.NewClient()

    // Authentication
    apiKey := os.Getenv("BINANCE_API_KEY")
    secretKey := os.Getenv("BINANCE_SECRET_KEY")

    if apiKey != "" && secretKey != "" {
        auth := &spotws.Auth{APIKey: apiKey}
        auth.SetSecretKey(secretKey)
        client.SetAuth(auth)
    }

    // Register event handlers
    client.HandleExecutionReport(func(event *models.ExecutionReport) error {
        fmt.Printf("Order update: %+v\n", event)
        return nil
    })

    // Connect and use
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err := client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect()

    // Get server time
    timeResp, err := client.Time(ctx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Server time: %d\n", timeResp.ServerTime)
}
```

### WebSocket Streams Example

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/openxapi/binance-go/ws/spot-streams"
    "github.com/openxapi/binance-go/ws/spot-streams/models"
)

func main() {
    // Create streams client
    client := spotstreams.NewClient()

    // Register event handlers
    client.HandleTradeEvent(func(event *models.TradeEvent) error {
        fmt.Printf("Trade: %s @ %s (%s)\n", 
            event.Symbol, event.Price, event.Quantity)
        return nil
    })

    client.HandleTickerEvent(func(event *models.TickerEvent) error {
        fmt.Printf("Ticker %s: %s (%s%%)\n", 
            event.Symbol, event.CurrentClosePrice, event.PriceChangePercent)
        return nil
    })

    // Connect
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err := client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect()

    // Subscribe to streams
    streams := []string{"btcusdt@trade", "btcusdt@ticker"}
    _, err = client.Subscribe(ctx, streams)
    if err != nil {
        log.Fatal(err)
    }

    // Keep alive
    time.Sleep(30 * time.Second)
}
```

## 🔐 Authentication

All authentication methods are supported across both REST and WebSocket APIs:

### HMAC-SHA256 (Most Common)

```go
// REST API
auth := spot.NewAuth(apiKey)
auth.SetSecretKey(secretKey)
ctx, err := auth.ContextWithValue(ctx)

// WebSocket API
auth := &spotws.Auth{APIKey: apiKey}
auth.SetSecretKey(secretKey)
client.SetAuth(auth)
```

### RSA

```go
// REST API
auth := spot.NewAuth(apiKey)
auth.PrivateKeyPath = "/path/to/private_key.pem"
ctx, err := auth.ContextWithValue(ctx)

// WebSocket API
auth := &spotws.Auth{
    APIKey: apiKey,
    PrivateKeyPath: "/path/to/private_key.pem",
}
client.SetAuth(auth)
```

### Ed25519

```go
// WebSocket API (Ed25519 sessions)
auth := &spotws.Auth{
    APIKey: apiKey,
    PrivateKeyPath: "/path/to/ed25519_private_key.pem",
}
client.SetAuth(auth)
```

## 🌐 Environments

### Production
- **REST**: `https://api.binance.com`
- **WebSocket API**: `wss://ws-api.binance.com/ws-api/v3`
- **WebSocket Streams**: `wss://stream.binance.com/ws`

### Testnet (for testing)
- **REST**: `https://testnet.binance.vision`
- **WebSocket API**: `wss://ws-api.testnet.binance.vision/ws-api/v3`
- **WebSocket Streams**: `wss://stream.testnet.binance.vision/ws`

```go
// Switch to testnet
client.SetActiveServer("testnet1")
```

## 📖 Documentation

- **[REST API Documentation](rest/README.md)** - Complete REST API guide
- **[WebSocket API Documentation](ws/README.md)** - WebSocket API and Streams guide
- **[Official Binance API Docs](https://developers.binance.com/)** - Binance developer documentation

## 🧪 Testing

The SDK includes comprehensive integration tests that demonstrate real-world usage:

```bash
# Set up environment variables
export BINANCE_API_KEY="your_api_key"
export BINANCE_SECRET_KEY="your_secret_key"

# Test REST APIs
cd rest/spot && go test -v
cd rest/umfutures && go test -v

# Test WebSocket APIs  
cd ws/spot && go test -v
cd ws/spot-streams && go test -v
cd ws/umfutures && go test -v
```

## 🎯 Use Cases

### Algorithmic Trading
```go
// High-frequency trading with WebSocket streams
client.HandleTradeEvent(func(event *models.TradeEvent) error {
    // Implement your trading strategy
    return processTradeSignal(event)
})
```

### Portfolio Management
```go
// Monitor account balances and positions
account, _, err := client.SpotTradingAPI.GetAccountV3(ctx).Execute()
if err != nil {
    return err
}
return updatePortfolio(account)
```

### Market Data Analysis
```go
// Collect real-time market data
client.HandleKlineEvent(func(event *models.KlineEvent) error {
    return storeCandlestickData(event)
})
```

### Risk Management
```go
// Monitor liquidations and funding rates
client.HandleLiquidationEvent(func(event *models.LiquidationEvent) error {
    return assessMarketRisk(event)
})
```

## ⚠️ Rate Limits

Please respect Binance API rate limits:

- **REST API**: Varies by endpoint (typically 1200 requests/minute)
- **WebSocket**: 5 messages per second per connection
- **Streams**: 1024 streams per connection

The SDK automatically handles rate limiting for REST APIs. For WebSocket APIs, implement your own rate limiting logic.

## 🔧 Advanced Features

### Error Handling

```go
// REST API errors
if err != nil {
    if apiErr, ok := err.(*spot.GenericOpenAPIError); ok {
        fmt.Printf("API Error: %s\n", apiErr.Error())
        fmt.Printf("Response: %s\n", string(apiErr.Body()))
    }
}

// WebSocket errors
client.HandleError(func(err error) {
    log.Printf("WebSocket error: %v", err)
})
```

### Proxy Support

```go
import "net/http"
import "net/url"

// Configure proxy for REST APIs
proxyURL, _ := url.Parse("http://proxy:8080")
config := spot.NewConfiguration()
config.HTTPClient = &http.Client{
    Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
}

// For WebSocket APIs, set environment variable
os.Setenv("HTTP_PROXY", "http://proxy:8080")
```

### Custom Timeouts

```go
// REST API timeouts
config := spot.NewConfiguration()
config.HTTPClient.Timeout = 30 * time.Second

// WebSocket connection timeouts
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
err := client.Connect(ctx)
```

## 🛡️ Security Best Practices

1. **Never hardcode API keys** - Use environment variables
2. **Use testnet for development** - Avoid accidental real trades
3. **Implement proper error handling** - Don't ignore API errors
4. **Respect rate limits** - Implement backoff strategies
5. **Validate signatures** - Ensure secure authentication
6. **Use read-only keys when possible** - Limit API key permissions

## 🤝 Contributing

This SDK is auto-generated from official Binance API specifications. To contribute:

1. **Report Issues**: [GitHub Issues](https://github.com/openxapi/binance-go/issues)
2. **Update Specifications**: Submit PRs to [OpenXAPI](https://github.com/openxapi/openxapi)
3. **Improve Templates**: Enhance the [code generation templates](https://github.com/openxapi/openxapi/tree/main/templates)

**Do not edit the generated code directly** - changes will be overwritten on the next generation.

## 📋 Requirements

- **Go**: 1.19 or later
- **Dependencies**: Managed via Go modules

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## ⚡ Performance Tips

1. **Reuse clients** - Don't create new clients for every request
2. **Use WebSocket for real-time data** - More efficient than polling REST APIs
3. **Batch operations** - Use batch endpoints when available
4. **Connection pooling** - Configure HTTP client for optimal performance
5. **Local caching** - Cache static data like exchange info

## 🆘 Support

- **Documentation**: [Binance API Docs](https://developers.binance.com/)
- **Issues**: [GitHub Issues](https://github.com/openxapi/binance-go/issues)
- **Discussions**: [GitHub Discussions](https://github.com/openxapi/binance-go/discussions)
- **OpenXAPI Project**: [OpenXAPI GitHub](https://github.com/openxapi/openxapi)

## ⚖️ Disclaimer

This is an unofficial SDK. Use at your own risk. Always refer to the official Binance API documentation for the most up-to-date information. The authors are not responsible for any financial losses incurred through the use of this SDK.

## 🏷️ Version Compatibility

| SDK Version | Binance API Version | Go Version |
|-------------|-------------------|------------|
| v1.x | 2024+ | 1.19+ |

---

**Happy Trading! 🚀**