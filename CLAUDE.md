# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is the **binance-go** SDK - a comprehensive Go client for Binance cryptocurrency exchange APIs. The codebase is **auto-generated** from OpenAPI and AsyncAPI specifications using the [OpenXAPI project](https://github.com/openxapi/openxapi).

**⚠️ Important**: This is generated code. Do NOT edit the generated files manually as changes will be overwritten during regeneration.

## Project Structure

The codebase is organized into two main modules:

### REST APIs (`/rest/`)
- **Module**: `github.com/openxapi/binance-go/rest`
- **Generator**: OpenAPI Generator
- **Products**: `spot/`, `umfutures/`, `cmfutures/`, `options/`, `pmargin/`

### WebSocket APIs (`/ws/`)
- **Module**: `github.com/openxapi/binance-go/ws`  
- **Generator**: OpenXAPI AsyncAPI generator
- **Products**: `spot/`, `umfutures/`, `cmfutures/` (APIs) + `spot-streams/`, `umfutures-streams/`, `cmfutures-streams/` (Streams)

## Build and Development Commands

### Basic Build Commands
```bash
# Build REST APIs
cd rest/spot && go build ./...
cd rest/umfutures && go build ./...
cd rest/cmfutures && go build ./...

# Build WebSocket APIs  
cd ws/spot && go build ./...
cd ws/spot-streams && go build ./...
cd ws/umfutures && go build ./...
```

### Testing

**Auto-Generated Unit Tests** (all skipped by default):
```bash
cd rest/spot && go test ./...
cd ws/spot && go test ./...
```

**Integration Tests** (comprehensive, requires API keys):
```bash
# Set up environment variables
export BINANCE_API_KEY="your_api_key"
export BINANCE_SECRET_KEY="your_secret_key"

# Test REST APIs
cd ../integration-tests/src/binance/go/rest/spot
go test -v

# Test WebSocket APIs  
cd ../integration-tests/src/binance/go/ws/spot
go test -v
```

## Code Architecture

### REST API Pattern
Generated clients follow standard OpenAPI patterns:
- **Configuration**: `configuration.go` - Client setup
- **Client**: `client.go` - HTTP client wrapper  
- **APIs**: `api_*.go` - Endpoint implementations
- **Models**: `model_*.go` - Request/response structs
- **Auth**: `signing.go` - Authentication utilities

**Usage Pattern**:
```go
config := spot.NewConfiguration()
client := spot.NewAPIClient(config)

// Authentication
auth := spot.NewAuth(apiKey)
auth.SetSecretKey(secretKey)
ctx, err := auth.ContextWithValue(ctx)

// API calls
result, httpRes, err := client.SpotTradingAPI.GetAccountV3(ctx).Execute()
```

### WebSocket API Pattern
Event-driven architecture with handlers:
- **Client**: `client.go` - WebSocket client
- **Models**: `models/` - Event structures
- **Auth**: `signing.go` - Authentication

**Usage Pattern**:
```go
client := spotws.NewClient()

// Authentication
auth := &spotws.Auth{APIKey: apiKey}
auth.SetSecretKey(secretKey)
client.SetAuth(auth)

// Register handlers
client.HandleExecutionReport(func(event *models.ExecutionReport) error {
    return nil
})

// Connect
err := client.Connect(ctx)
```

### Authentication Architecture
Supports three authentication methods:
- **HMAC-SHA256**: API key + secret (most common)
- **RSA**: Private key file
- **Ed25519**: Ed25519 private key

Authentication is context-based for REST APIs and client-based for WebSocket APIs.

## Module Dependencies

### REST Module (`rest/go.mod`)
```
github.com/stretchr/testify v1.10.0
gopkg.in/validator.v2 v2.0.1
```

### WebSocket Module (`ws/go.mod`)
```
github.com/google/uuid v1.6.0
github.com/gorilla/websocket v1.5.3
```

## Development Workflow

### For Contributing to Generated Code
1. Update specifications in [OpenXAPI](https://github.com/openxapi/openxapi)
2. Regenerate code using OpenXAPI tools
3. Replace generated files in this repository

### For Testing Changes
1. Use integration tests at `../integration-tests/src/binance/go/`
2. Set up Binance API credentials
3. Run comprehensive test suites for each product

### For Adding New Products
1. Add AsyncAPI/OpenAPI specs to OpenXAPI project
2. Configure code generation templates
3. Generate new product modules
4. Update documentation

## Key Limitations

- No custom build system (uses standard Go tooling)
- Generated unit tests are skipped placeholders
- Cannot modify generated files directly
- Must work through OpenXAPI project for changes

## Testing Focus

Focus on integration tests rather than unit tests:
- Unit tests are auto-generated placeholders (all skipped)
- Integration tests provide comprehensive real-world testing
- Located in separate integration-tests directory
- Require actual Binance API credentials for testing