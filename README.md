# Go API client for Binance

[![Go Reference](https://pkg.go.dev/badge/github.com/openxapi/binance-go.svg)](https://pkg.go.dev/github.com/openxapi/binance-go)

OpenAPI specification for Binance cryptocurrency exchange.

## Overview

This API client was automatically generated by [OpenXAPI](https://github.com/openxapi/openxapi) using [OpenAPI Generator](https://openapi-generator.tech).

**Please do not edit the generated code manually.**

If you need to update the API client, please submit an issue or a PR to [OpenXAPI](https://github.com/openxapi/openxapi) to update the [OpenAPI specification](https://github.com/openxapi/openxapi/tree/main/specs/binance/openapi) or the [Go client templates](https://github.com/openxapi/openxapi/tree/main/templates/binance/openapi/go), and regenerate the API client.

## Supported Products

- [x] [Spot Trading](./spot)
- [x] Derivatives Trading
  - [x] [USDS-M Futures](./derivatives/umfutures)
  - [x] [COIN-M Futures](./derivatives/cmfutures)
  - [x] [Options](./derivatives/options)
  - [x] [Portfolio Margin](./derivatives/pmargin)
  - [x] [Portfolio Margin Pro](./derivatives/pmarginpro)
  - [x] [Futures Data](./derivatives/futuresdata)
- [x] [Margin Trading](./margin)
- [x] Algo Trading
- [x] Wallet
- [x] Copy Trading
- [x] Convert
- [x] Sub Account
- [ ] More to come...

Feel free to contribute to this project by adding support for other products.

With [OpenXAPI](https://github.com/openxapi/openxapi), you can add support for other products and generate API clients in different languages in a breeze.

Following is an example of how to use the spot API client, please refer to the sub-folders for other products for more details.

## Import

```go
import (
    spot "github.com/openxapi/binance-go/spot"
)
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Usage

Example

```go
conf := spot.NewConfiguration()
client := spot.NewAPIClient(conf)
ctx := context.Background()

info, res, err := client.GeneralAPI.SpotGetExchangeInfoV3(ctx).Symbol("BTCUSDT").Execute()
if err != nil {
    fmt.Println(err)
}
fmt.Printf("%+v\n", info)
fmt.Printf("%+v\n", res)
```

## Authentication

The API client supports HMAC, RSA and Ed25519 authentication.

The authentication is calculated per request, so you can use different keys for different requests.

### HMAC

Example

```go
    conf := spot.NewConfiguration()
    client := spot.NewAPIClient(conf)
    ctx := context.Background()

    // get API key from env
    apiKey := os.Getenv("BINANCE_API_KEY")
    auth := spot.NewAuth(apiKey)    
    auth.SetSecretKey(os.Getenv("BINANCE_SECRET_KEY"))
    ctx, err = auth.ContextWithValue(ctx)
    if err != nil {
        fmt.Println(err)
    }
    // Get current time in millisecond
    user, res, err := client.AccountAPI.SpotGetAccountV3(ctx).Timestamp(time.Now().UnixMilli()).Execute()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("%+v\n", user)
    fmt.Printf("%+v\n", res)
```

### RSA or Ed25519

Example

```go
    conf := spot.NewConfiguration()
    client := spot.NewAPIClient(conf)
    ctx := context.Background()

    // get API key from env
    apiKey := os.Getenv("BINANCE_API_KEY")
    auth := spot.NewAuth(apiKey)
    // Key type will be auto-detected, you can use RSA or Ed25519 key here
    auth.PrivateKeyPath = "/local/.keys/binance_api_test_private_key.pem"
    ctx, err = auth.ContextWithValue(ctx)
    if err != nil {
        fmt.Println(err)
    }
    // Get current time in millisecond
    user, res, err := client.AccountAPI.SpotGetAccountV3(ctx).Timestamp(time.Now().UnixMilli()).Execute()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("%+v\n", user)
    fmt.Printf("%+v\n", res)
```

## Configuration

### Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `spot.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), spot.ContextServerIndex, 1)
```
