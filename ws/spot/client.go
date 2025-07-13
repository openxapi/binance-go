package spot
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/openxapi/binance-go/ws/spot/models"
)

// Context keys for authentication and configuration
type contextKey string

const (
	// ContextBinanceAuth takes Auth as authentication for the request
	ContextBinanceAuth = contextKey("binanceAuth")
)

// ServerInfo represents a WebSocket server configuration
type ServerInfo struct {
	Name        string `json:"name"`        // Server identifier (e.g., "mainnet", "testnet")
	URL         string `json:"url"`         // Full WebSocket URL
	Host        string `json:"host"`        // Server host
	Pathname    string `json:"pathname"`    // URL path
	Protocol    string `json:"protocol"`    // ws or wss
	Title       string `json:"title"`       // Human-readable title
	Summary     string `json:"summary"`     // Short description
	Description string `json:"description"` // Detailed description
	Active      bool   `json:"active"`      // Whether this server is currently active
}

// ServerManager manages multiple WebSocket servers
type ServerManager struct {
	servers      map[string]*ServerInfo
	activeServer string
	mu           sync.RWMutex
}

// NewServerManager creates a new server manager with default servers
func NewServerManager() *ServerManager {
	sm := &ServerManager{
		servers: make(map[string]*ServerInfo),
	}
	
	// Initialize with predefined servers from AsyncAPI spec
	// Using direct assignment since this is initialization (no risk of conflicts)
	sm.servers["mainnet1"] = &ServerInfo{
		Name:        "mainnet1",
		URL:         "wss://ws-api.binance.com/ws-api/v3",
		Host:        "ws-api.binance.com",
		Pathname:    "/ws-api/v3",
		Protocol:    "wss",
		Title:       "Binance Mainnet Server",
		Summary:     "Binance Spot WebSocket API Server (Mainnet)",
		Description: "WebSocket server for binance exchange spot API (mainnet environment)",
		Active:      false,
	}
	sm.servers["mainnet2"] = &ServerInfo{
		Name:        "mainnet2",
		URL:         "wss://ws-api.binance.com:9443/ws-api/v3",
		Host:        "ws-api.binance.com:9443",
		Pathname:    "/ws-api/v3",
		Protocol:    "wss",
		Title:       "Binance Mainnet Server",
		Summary:     "Binance Spot WebSocket API Server (Mainnet)",
		Description: "WebSocket server for binance exchange spot API (mainnet environment)",
		Active:      false,
	}
	sm.servers["testnet1"] = &ServerInfo{
		Name:        "testnet1",
		URL:         "wss://ws-api.testnet.binance.vision/ws-api/v3",
		Host:        "ws-api.testnet.binance.vision",
		Pathname:    "/ws-api/v3",
		Protocol:    "wss",
		Title:       "Binance Testnet Server",
		Summary:     "Binance Spot WebSocket API Server (Testnet)",
		Description: "WebSocket server for binance exchange spot API (testnet environment)",
		Active:      false,
	}
	sm.servers["testnet2"] = &ServerInfo{
		Name:        "testnet2",
		URL:         "wss://ws-api.testnet.binance.vision:9443/ws-api/v3",
		Host:        "ws-api.testnet.binance.vision:9443",
		Pathname:    "/ws-api/v3",
		Protocol:    "wss",
		Title:       "Binance Testnet Server",
		Summary:     "Binance Spot WebSocket API Server (Testnet)",
		Description: "WebSocket server for binance exchange spot API (testnet environment)",
		Active:      false,
	}
	
	// Set first server as active by default
	sm.SetActiveServer("mainnet1")
	
	return sm
}

// AddServer adds a new server to the manager
func (sm *ServerManager) AddServer(name string, server *ServerInfo) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	if server == nil {
		return fmt.Errorf("server info cannot be nil")
	}
	
	// Check if server name already exists
	if _, exists := sm.servers[name]; exists {
		return fmt.Errorf("server with name '%s' already exists, use UpdateServer() to modify it", name)
	}
	
	// Validate URL
	if _, err := url.Parse(server.URL); err != nil {
		return fmt.Errorf("invalid server URL '%s': %w", server.URL, err)
	}
	
	server.Name = name
	sm.servers[name] = server
	
	// If this is the first server, make it active
	if sm.activeServer == "" {
		sm.activeServer = name
		server.Active = true
	}
	
	return nil
}

// AddOrUpdateServer adds a new server or updates an existing one
func (sm *ServerManager) AddOrUpdateServer(name string, server *ServerInfo) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	if server == nil {
		return fmt.Errorf("server info cannot be nil")
	}
	
	// Validate URL
	if _, err := url.Parse(server.URL); err != nil {
		return fmt.Errorf("invalid server URL '%s': %w", server.URL, err)
	}
	
	// Preserve active status if server already exists
	server.Name = name
	if existingServer, exists := sm.servers[name]; exists {
		server.Active = existingServer.Active
	}
	
	sm.servers[name] = server
	
	// If this is the first server, make it active
	if sm.activeServer == "" {
		sm.activeServer = name
		server.Active = true
	}
	
	return nil
}

// RemoveServer removes a server from the manager
func (sm *ServerManager) RemoveServer(name string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	if _, exists := sm.servers[name]; !exists {
		return fmt.Errorf("server '%s' not found", name)
	}
	
	// Don't allow removing the active server
	if sm.activeServer == name {
		return fmt.Errorf("cannot remove active server '%s', switch to another server first", name)
	}
	
	delete(sm.servers, name)
	return nil
}

// UpdateServer updates an existing server's configuration
func (sm *ServerManager) UpdateServer(name string, server *ServerInfo) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	if _, exists := sm.servers[name]; !exists {
		return fmt.Errorf("server '%s' not found", name)
	}
	
	if server == nil {
		return fmt.Errorf("server info cannot be nil")
	}
	
	// Validate URL
	if _, err := url.Parse(server.URL); err != nil {
		return fmt.Errorf("invalid server URL '%s': %w", server.URL, err)
	}
	
	// Preserve active status and name
	server.Name = name
	server.Active = (sm.activeServer == name)
	sm.servers[name] = server
	
	return nil
}

// UpdateServerPathname updates the pathname of an existing server
func (sm *ServerManager) UpdateServerPathname(name string, pathname string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	server, exists := sm.servers[name]
	if !exists {
		return fmt.Errorf("server '%s' not found", name)
	}
	
	// Update pathname and rebuild URL
	server.Pathname = pathname
	server.URL = fmt.Sprintf("%s://%s%s", server.Protocol, server.Host, pathname)
	
	// Validate the new URL (skip validation if it contains template variables)
	if !strings.Contains(server.URL, "{") {
		if _, err := url.Parse(server.URL); err != nil {
			return fmt.Errorf("invalid server URL '%s': %w", server.URL, err)
		}
	}
	
	return nil
}

// ResolveServerURL resolves template variables in server URL
func (sm *ServerManager) ResolveServerURL(name string, variables map[string]string) (string, error) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	server, exists := sm.servers[name]
	if !exists {
		return "", fmt.Errorf("server '%s' not found", name)
	}
	
	resolvedURL := server.URL
	
	// Replace template variables
	for key, value := range variables {
		placeholder := fmt.Sprintf("{%s}", key)
		resolvedURL = strings.ReplaceAll(resolvedURL, placeholder, value)
	}
	
	// Validate the resolved URL
	if _, err := url.Parse(resolvedURL); err != nil {
		return "", fmt.Errorf("invalid resolved URL '%s': %w", resolvedURL, err)
	}
	
	return resolvedURL, nil
}

// SetActiveServer sets the active server
func (sm *ServerManager) SetActiveServer(name string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	if _, exists := sm.servers[name]; !exists {
		return fmt.Errorf("server '%s' not found", name)
	}
	
	// Deactivate current active server
	if sm.activeServer != "" {
		if currentActive := sm.servers[sm.activeServer]; currentActive != nil {
			currentActive.Active = false
		}
	}
	
	// Activate new server
	sm.activeServer = name
	sm.servers[name].Active = true
	
	return nil
}

// GetActiveServer returns the currently active server
func (sm *ServerManager) GetActiveServer() *ServerInfo {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	if sm.activeServer == "" {
		return nil
	}
	
	server := sm.servers[sm.activeServer]
	if server == nil {
		return nil
	}
	
	// Return a copy to prevent external modification
	return &ServerInfo{
		Name:        server.Name,
		URL:         server.URL,
		Host:        server.Host,
		Pathname:    server.Pathname,
		Protocol:    server.Protocol,
		Title:       server.Title,
		Summary:     server.Summary,
		Description: server.Description,
		Active:      server.Active,
	}
}

// GetServer returns a specific server by name
func (sm *ServerManager) GetServer(name string) *ServerInfo {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	server := sm.servers[name]
	if server == nil {
		return nil
	}
	
	// Return a copy to prevent external modification
	return &ServerInfo{
		Name:        server.Name,
		URL:         server.URL,
		Host:        server.Host,
		Pathname:    server.Pathname,
		Protocol:    server.Protocol,
		Title:       server.Title,
		Summary:     server.Summary,
		Description: server.Description,
		Active:      server.Active,
	}
}

// ListServers returns all available servers
func (sm *ServerManager) ListServers() map[string]*ServerInfo {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	result := make(map[string]*ServerInfo, len(sm.servers))
	for name, server := range sm.servers {
		result[name] = &ServerInfo{
			Name:        server.Name,
			URL:         server.URL,
			Host:        server.Host,
			Pathname:    server.Pathname,
			Protocol:    server.Protocol,
			Title:       server.Title,
			Summary:     server.Summary,
			Description: server.Description,
			Active:      server.Active,
		}
	}
	
	return result
}

// GetActiveServerURL returns the URL of the currently active server
func (sm *ServerManager) GetActiveServerURL() string {
	if server := sm.GetActiveServer(); server != nil {
		return server.URL
	}
	return ""
}

// APIError represents an error returned by the Binance WebSocket API
type APIError struct {
	Status  int    `json:"status"`  // HTTP-like status code from the response
	Code    int    `json:"code"`    // Binance-specific error code
	Message string `json:"msg"`     // Error message
	ID      string `json:"id"`      // Request ID that caused the error
}

// Error implements the error interface
func (e APIError) Error() string {
	return fmt.Sprintf("binance api error: status=%d, code=%d, message=%s, id=%s", e.Status, e.Code, e.Message, e.ID)
}

// IsAPIError checks if an error is an APIError
func IsAPIError(err error) (*APIError, bool) {
	if apiErr, ok := err.(APIError); ok {
		return &apiErr, true
	}
	if apiErr, ok := err.(*APIError); ok {
		return apiErr, true
	}
	return nil, false
}

// ResponseHandler represents a high-performance handler for WebSocket responses
type ResponseHandler struct {
	RequestID string
	Handler   func([]byte, error) error
}

// TypedResponseHandler is a generic interface for type-safe response handling
type TypedResponseHandler[T any] interface {
	Handle(*T, error) error
}

// HandlerFunc is a function type that implements TypedResponseHandler
type HandlerFunc[T any] func(*T, error) error

// Handle implements TypedResponseHandler interface
func (f HandlerFunc[T]) Handle(response *T, err error) error {
	return f(response, err)
}

// EventHandler handles all possible response types with optimized lookup
type EventHandler struct {
	handlers sync.Map // Using sync.Map for better concurrent performance
}

// NewEventHandler creates a new optimized event handler
func NewEventHandler() *EventHandler {
	return &EventHandler{}
}

// RegisterHandler registers a handler for a specific response type
func (e *EventHandler) RegisterHandler(responseType string, handler func(interface{}) error) {
	e.handlers.Store(responseType, handler)
}

// HandleResponse processes a event based on its type with optimized lookup
func (e *EventHandler) HandleResponse(eventType string, data []byte) error {
	if handler, ok := e.handlers.Load(eventType); ok {
		if h, ok := handler.(func(interface{}) error); ok {
			return h(data)
		}
	}
	
	log.Printf("No handler found for event type: %s", eventType)
	return nil
}

// Client represents a high-performance WebSocket client for Binance Spot WebSocket API
type Client struct {
	conn             *websocket.Conn
	serverManager    *ServerManager   // Manages multiple servers
	responseHandlers sync.Map         // Using sync.Map for better concurrent performance
	eventHandler     *EventHandler
	responseList     []interface{}    // Global list of all received responses
	responseListMu   sync.RWMutex     // Separate mutex for response list
	auth             *Auth            // Authentication configuration
	done             chan struct{}
	isConnected      bool             // Connection status flag
	handlers         eventHandlers    // Event handlers registry
	
	// Pre-allocated buffer for JSON parsing to reduce allocations
	jsonBuffer []byte
	bufferMu   sync.Mutex
}

// NewClient creates a new high-performance WebSocket client
func NewClient() *Client {
	return &Client{
		serverManager: NewServerManager(),
		eventHandler:  NewEventHandler(),
		responseList:  make([]interface{}, 0, 100), // Pre-allocate with capacity
		done:          make(chan struct{}),
		jsonBuffer:    make([]byte, 0, 1024), // Pre-allocate JSON buffer
	}
}

// NewClientWithAuth creates a new high-performance WebSocket client with authentication
func NewClientWithAuth(auth *Auth) *Client {
	client := NewClient()
	client.auth = auth
	return client
}

// SetAuth sets authentication for the client
func (c *Client) SetAuth(auth *Auth) {
	c.auth = auth
}

// Server Management Methods

// AddServer adds a new server to the client
func (c *Client) AddServer(name string, serverURL string, title string, description string) error {
	if c.conn != nil {
		return fmt.Errorf("cannot add server while connected - please disconnect first")
	}
	
	// Parse URL to extract components
	parsedURL, err := url.Parse(serverURL)
	if err != nil {
		return fmt.Errorf("invalid URL format: %w", err)
	}
	
	server := &ServerInfo{
		Name:        name,
		URL:         serverURL,
		Host:        parsedURL.Host,
		Pathname:    parsedURL.Path,
		Protocol:    parsedURL.Scheme,
		Title:       title,
		Summary:     fmt.Sprintf("WebSocket API Server (%s)", name),
		Description: description,
		Active:      false,
	}
	
	return c.serverManager.AddServer(name, server)
}

// AddOrUpdateServer adds a new server or updates an existing one
func (c *Client) AddOrUpdateServer(name string, serverURL string, title string, description string) error {
	if c.conn != nil {
		return fmt.Errorf("cannot add/update server while connected - please disconnect first")
	}
	
	// Parse URL to extract components
	parsedURL, err := url.Parse(serverURL)
	if err != nil {
		return fmt.Errorf("invalid URL format: %w", err)
	}
	
	server := &ServerInfo{
		Name:        name,
		URL:         serverURL,
		Host:        parsedURL.Host,
		Pathname:    parsedURL.Path,
		Protocol:    parsedURL.Scheme,
		Title:       title,
		Summary:     fmt.Sprintf("WebSocket API Server (%s)", name),
		Description: description,
		Active:      false,
	}
	
	return c.serverManager.AddOrUpdateServer(name, server)
}

// RemoveServer removes a server from the client
func (c *Client) RemoveServer(name string) error {
	if c.conn != nil {
		return fmt.Errorf("cannot remove server while connected - please disconnect first")
	}
	
	return c.serverManager.RemoveServer(name)
}

// UpdateServer updates an existing server's configuration
func (c *Client) UpdateServer(name string, serverURL string, title string, description string) error {
	if c.conn != nil {
		return fmt.Errorf("cannot update server while connected - please disconnect first")
	}
	
	// Parse URL to extract components
	parsedURL, err := url.Parse(serverURL)
	if err != nil {
		return fmt.Errorf("invalid URL format: %w", err)
	}
	
	server := &ServerInfo{
		Name:        name,
		URL:         serverURL,
		Host:        parsedURL.Host,
		Pathname:    parsedURL.Path,
		Protocol:    parsedURL.Scheme,
		Title:       title,
		Summary:     fmt.Sprintf("WebSocket API Server (%s)", name),
		Description: description,
		Active:      false, // Will be set correctly by UpdateServer
	}
	
	return c.serverManager.UpdateServer(name, server)
}

// SetActiveServer sets the active server by name
func (c *Client) SetActiveServer(name string) error {
	if c.conn != nil {
		return fmt.Errorf("cannot change active server while connected - please disconnect first")
	}
	
	return c.serverManager.SetActiveServer(name)
}

// GetActiveServer returns the currently active server information
func (c *Client) GetActiveServer() *ServerInfo {
	return c.serverManager.GetActiveServer()
}

// GetServer returns information about a specific server
func (c *Client) GetServer(name string) *ServerInfo {
	return c.serverManager.GetServer(name)
}

// ListServers returns all available servers
func (c *Client) ListServers() map[string]*ServerInfo {
	return c.serverManager.ListServers()
}

// GetCurrentURL returns the URL of the currently active server
func (c *Client) GetCurrentURL() string {
	return c.serverManager.GetActiveServerURL()
}

// Legacy SetURL method for backward compatibility
// Deprecated: Use SetActiveServer() or UpdateServer() instead
func (c *Client) SetURL(newURL string) error {
	if c.conn != nil {
		return fmt.Errorf("cannot change URL while connected - please disconnect first")
	}
	
	// Find if this URL matches any existing server
	servers := c.serverManager.ListServers()
	for name, server := range servers {
		if server.URL == newURL {
			return c.serverManager.SetActiveServer(name)
		}
	}
	
	// If URL doesn't match any existing server, update the active server
	activeServer := c.serverManager.GetActiveServer()
	if activeServer != nil {
		return c.UpdateServer(activeServer.Name, newURL, activeServer.Title, activeServer.Description)
	}
	
	// If no active server, add as new server
	return c.AddServer("custom", newURL, "Custom Server", "Custom WebSocket server")
}

// Connect establishes a WebSocket connection to the active server
func (c *Client) Connect(ctx context.Context) error {
	currentURL := c.serverManager.GetActiveServerURL()
	if currentURL == "" {
		return fmt.Errorf("no active server configured")
	}
	
	u, err := url.Parse(currentURL)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}

	dialer := websocket.DefaultDialer
	dialer.HandshakeTimeout = 10 * time.Second

	conn, _, err := dialer.DialContext(ctx, u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w", currentURL, err)
	}

	c.conn = conn
	c.isConnected = true
	go c.readMessages()
	return nil
}

// ConnectToServer establishes a WebSocket connection to a specific server
func (c *Client) ConnectToServer(ctx context.Context, serverName string) error {
	if err := c.SetActiveServer(serverName); err != nil {
		return fmt.Errorf("failed to set active server: %w", err)
	}
	
	return c.Connect(ctx)
}

// Disconnect closes the WebSocket connection safely
func (c *Client) Disconnect() error {
	c.isConnected = false
	
	// Safely close the done channel only once
	select {
	case <-c.done:
		// Channel already closed, do nothing
	default:
		close(c.done)
	}
	
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// GenerateRequestID generates a unique UUID v4 request ID (global function)
func GenerateRequestID() string {
	return uuid.New().String()
}

// readMessages reads messages from the WebSocket connection
func (c *Client) readMessages() {
	defer func() {
		c.isConnected = false
		if c.conn != nil {
			c.conn.Close()
		}
	}()

	for {
		select {
		case <-c.done:
			return
		default:
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				// Check if the error is due to connection being closed
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					// Normal close - no need to log as error
					return
				}
				// Check if this is a network connection closed error (also normal during shutdown)
				if strings.Contains(err.Error(), "use of closed network connection") {
					// Normal network close - no need to log as error
					return
				}
				log.Printf("Error reading message: %v", err)
				return
			}

			if err := c.handleMessage(message); err != nil {
				log.Printf("Error handling message: %v", err)
			}
		}
	}
}

// handleMessage processes incoming WebSocket messages
func (c *Client) handleMessage(data []byte) error {
	// First check if this is an array stream (like !assetIndex@arr)
	// by trying to parse as an array first
	var arrayData []interface{}
	if err := json.Unmarshal(data, &arrayData); err == nil {
		// This is an array stream - delegate to stream processing logic
		return c.handleEventMessage("arrayStream", data)
	}
	
	// Parse the message to determine its type
	var genericMessage map[string]interface{}
	if err := json.Unmarshal(data, &genericMessage); err != nil {
		return fmt.Errorf("failed to parse message: %w", err)
	}

	// Check if this is a response to a request (has "id" field)
	if id, hasID := genericMessage["id"]; hasID {
		// Parse response structure to check for errors
		var response struct {
			ID     interface{} `json:"id"`
			Status int         `json:"status"`
			Result interface{} `json:"result,omitempty"`
			Error  *struct {
				Code int    `json:"code"`
				Msg  string `json:"msg"`
			} `json:"error,omitempty"`
			RateLimits []interface{} `json:"rateLimits,omitempty"`
		}

		if err := json.Unmarshal(data, &response); err != nil {
			return fmt.Errorf("failed to parse response structure: %w", err)
		}

		requestID := fmt.Sprintf("%v", id)

		// Check if status indicates an error
		var responseError error
		if response.Status != 200 && response.Error != nil {
			responseError = &APIError{
				Status:  response.Status,
				Code:    response.Error.Code,
				Message: response.Error.Msg,
				ID:      requestID,
			}
		}

		return c.handleRequestResponse(requestID, data, responseError)
	}

	// Check for event structure with nested "event" object (Binance event messages)
	if eventObj, hasEventObj := genericMessage["event"]; hasEventObj {
		if eventMap, ok := eventObj.(map[string]interface{}); ok {
			if eventType, hasEventType := eventMap["e"]; hasEventType {
				return c.handleEventMessage(eventType.(string), data)
			}
		}
	}

	// Check for direct event type field (stream messages)
		if eventType, hasEventType := genericMessage["e"]; hasEventType {
			if eventTypeStr, ok := eventType.(string); ok {
				return c.handleEventMessage(eventTypeStr, data)
			}
		}

		// If we can't determine the message type, log it
		log.Printf("Unknown message type: %s", string(data))
		return nil
}

// handleRequestResponse handles responses to specific requests
func (c *Client) handleRequestResponse(requestID string, data []byte, err error) error {
	if handler, ok := c.responseHandlers.Load(requestID); ok {
		c.responseHandlers.Delete(requestID) // Clean up after handling
		
		if h, ok := handler.(ResponseHandler); ok {
			return h.Handler(data, err)
		}
	}
	
	// Store in global response list for debugging/monitoring
	c.responseListMu.Lock()
	defer c.responseListMu.Unlock()
	
	var response interface{}
	if err == nil {
		json.Unmarshal(data, &response)
	}
	c.responseList = append(c.responseList, response)
	
	return nil
}

// handleEventMessage handles event messages (like balance updates, execution reports, etc.)
func (c *Client) handleEventMessage(eventType string, data []byte) error {
	// For stream modules, use the dedicated stream processing logic
	// Handle with event handler
	return c.eventHandler.HandleResponse(eventType, data)
}

// sendRequest sends a request over the WebSocket connection
func (c *Client) sendRequest(request map[string]interface{}) error {
	if c.conn == nil {
		return fmt.Errorf("not connected")
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	return c.conn.WriteMessage(websocket.TextMessage, data)
}

// GetResponseList returns a copy of all received responses (for debugging)
func (c *Client) GetResponseList() []interface{} {
	c.responseListMu.RLock()
	defer c.responseListMu.RUnlock()
	
	result := make([]interface{}, len(c.responseList))
	copy(result, c.responseList)
	return result
}

// ClearResponseList clears the response list
func (c *Client) ClearResponseList() {
	c.responseListMu.Lock()
	defer c.responseListMu.Unlock()
	c.responseList = c.responseList[:0]
}

// Health check and utility methods
func (c *Client) IsConnected() bool {
	return c.isConnected && c.conn != nil
}

// Deprecated: Use GetCurrentURL() instead
func (c *Client) GetURL() string {
	return c.GetCurrentURL()
}


// Event handler registry placeholder type
type eventHandlers struct {
	// This struct will be populated by module-specific handlers
	// It serves as a placeholder to satisfy the Client struct definition
}

// SendAccountCommission sends a account.commission request using typed request/response structs
// Authentication required: USER_DATA
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendAccountCommission(ctx context.Context, request *models.AccountCommissionRequest, responseHandler func(*models.AccountCommissionResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "account.commission"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method account.commission requires parameters but none provided: %v", []string{"symbol"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for USER_DATA request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.AccountCommissionResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendAccountRateLimitsOrders sends a account.rateLimits.orders request using typed request/response structs
// Authentication required: USER_DATA
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendAccountRateLimitsOrders(ctx context.Context, request *models.AccountRateLimitsOrdersRequest, responseHandler func(*models.AccountRateLimitsOrdersResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "account.rateLimits.orders"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for USER_DATA request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.AccountRateLimitsOrdersResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendAccountStatus sends a account.status request using typed request/response structs
// Authentication required: USER_DATA
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendAccountStatus(ctx context.Context, request *models.AccountStatusRequest, responseHandler func(*models.AccountStatusResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "account.status"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for USER_DATA request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.AccountStatusResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendAllOrderLists sends a allOrderLists request using typed request/response structs
// Authentication required: USER_DATA
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendAllOrderLists(ctx context.Context, request *models.AllOrderListsRequest, responseHandler func(*models.AllOrderListsResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "allOrderLists"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for USER_DATA request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.AllOrderListsResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendAllOrders sends a allOrders request using typed request/response structs
// Authentication required: USER_DATA
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendAllOrders(ctx context.Context, request *models.AllOrdersRequest, responseHandler func(*models.AllOrdersResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "allOrders"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method allOrders requires parameters but none provided: %v", []string{"symbol"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for USER_DATA request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.AllOrdersResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendAvgPrice sends a avgPrice request using typed request/response structs
// Authentication required: NONE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendAvgPrice(ctx context.Context, request *models.AvgPriceRequest, responseHandler func(*models.AvgPriceResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "avgPrice"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method avgPrice requires parameters but none provided: %v", []string{"symbol"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.AvgPriceResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendDepth sends a depth request using typed request/response structs
// Authentication required: NONE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendDepth(ctx context.Context, request *models.DepthRequest, responseHandler func(*models.DepthResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "depth"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method depth requires parameters but none provided: %v", []string{"symbol"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.DepthResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendExchangeInfo sends a exchangeInfo request using typed request/response structs
// Authentication required: NONE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendExchangeInfo(ctx context.Context, request *models.ExchangeInfoRequest, responseHandler func(*models.ExchangeInfoResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "exchangeInfo"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.ExchangeInfoResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendKlines sends a klines request using typed request/response structs
// Authentication required: NONE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendKlines(ctx context.Context, request *models.KlinesRequest, responseHandler func(*models.KlinesResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "klines"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method klines requires parameters but none provided: %v", []string{"symbol", "interval"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.KlinesResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendMyAllocations sends a myAllocations request using typed request/response structs
// Authentication required: USER_DATA
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendMyAllocations(ctx context.Context, request *models.MyAllocationsRequest, responseHandler func(*models.MyAllocationsResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "myAllocations"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method myAllocations requires parameters but none provided: %v", []string{"symbol"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for USER_DATA request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.MyAllocationsResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendMyPreventedMatches sends a myPreventedMatches request using typed request/response structs
// Authentication required: USER_DATA
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendMyPreventedMatches(ctx context.Context, request *models.MyPreventedMatchesRequest, responseHandler func(*models.MyPreventedMatchesResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "myPreventedMatches"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method myPreventedMatches requires parameters but none provided: %v", []string{"symbol"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for USER_DATA request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.MyPreventedMatchesResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendMyTrades sends a myTrades request using typed request/response structs
// Authentication required: USER_DATA
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendMyTrades(ctx context.Context, request *models.MyTradesRequest, responseHandler func(*models.MyTradesResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "myTrades"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method myTrades requires parameters but none provided: %v", []string{"symbol"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for USER_DATA request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.MyTradesResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendOpenOrderListsStatus sends a openOrderLists.status request using typed request/response structs
// Authentication required: USER_DATA
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendOpenOrderListsStatus(ctx context.Context, request *models.OpenOrderListsStatusRequest, responseHandler func(*models.OpenOrderListsStatusResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "openOrderLists.status"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for USER_DATA request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.OpenOrderListsStatusResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendOpenOrdersCancelAll sends a openOrders.cancelAll request using typed request/response structs
// Authentication required: TRADE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendOpenOrdersCancelAll(ctx context.Context, request *models.OpenOrdersCancelAllRequest, responseHandler func(*models.OpenOrdersCancelAllResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "openOrders.cancelAll"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method openOrders.cancelAll requires parameters but none provided: %v", []string{"symbol"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for TRADE request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.OpenOrdersCancelAllResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendOpenOrdersStatus sends a openOrders.status request using typed request/response structs
// Authentication required: USER_DATA
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendOpenOrdersStatus(ctx context.Context, request *models.OpenOrdersStatusRequest, responseHandler func(*models.OpenOrdersStatusResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "openOrders.status"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for USER_DATA request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.OpenOrdersStatusResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendOrderAmendKeepPriority sends a order.amend.keepPriority request using typed request/response structs
// Authentication required: TRADE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendOrderAmendKeepPriority(ctx context.Context, request *models.OrderAmendKeepPriorityRequest, responseHandler func(*models.OrderAmendKeepPriorityResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "order.amend.keepPriority"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method order.amend.keepPriority requires parameters but none provided: %v", []string{"symbol", "newQty"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for TRADE request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.OrderAmendKeepPriorityResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendOrderAmendments sends a order.amendments request using typed request/response structs
// Authentication required: USER_DATA
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendOrderAmendments(ctx context.Context, request *models.OrderAmendmentsRequest, responseHandler func(*models.OrderAmendmentsResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "order.amendments"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method order.amendments requires parameters but none provided: %v", []string{"symbol", "orderId"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for USER_DATA request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.OrderAmendmentsResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendOrderCancel sends a order.cancel request using typed request/response structs
// Authentication required: TRADE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendOrderCancel(ctx context.Context, request *models.OrderCancelRequest, responseHandler func(*models.OrderCancelResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "order.cancel"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method order.cancel requires parameters but none provided: %v", []string{"symbol", "orderId"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for TRADE request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.OrderCancelResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendOrderCancelReplace sends a order.cancelReplace request using typed request/response structs
// Authentication required: TRADE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendOrderCancelReplace(ctx context.Context, request *models.OrderCancelReplaceRequest, responseHandler func(*models.OrderCancelReplaceResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "order.cancelReplace"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method order.cancelReplace requires parameters but none provided: %v", []string{"symbol", "cancelReplaceMode", "cancelOrderId", "side", "type"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for TRADE request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.OrderCancelReplaceResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendOrderListCancel sends a orderList.cancel request using typed request/response structs
// Authentication required: TRADE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendOrderListCancel(ctx context.Context, request *models.OrderListCancelRequest, responseHandler func(*models.OrderListCancelResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "orderList.cancel"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method orderList.cancel requires parameters but none provided: %v", []string{"symbol", "orderListId"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for TRADE request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.OrderListCancelResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendOrderListPlace sends a orderList.place request using typed request/response structs
// Authentication required: TRADE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendOrderListPlace(ctx context.Context, request *models.OrderListPlaceRequest, responseHandler func(*models.OrderListPlaceResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "orderList.place"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method orderList.place requires parameters but none provided: %v", []string{"symbol", "side", "price", "quantity", "stopPrice", "trailingDelta"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for TRADE request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.OrderListPlaceResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendOrderListPlaceOco sends a orderList.place.oco request using typed request/response structs
// Authentication required: TRADE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendOrderListPlaceOco(ctx context.Context, request *models.OrderListPlaceOcoRequest, responseHandler func(*models.OrderListPlaceOcoResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "orderList.place.oco"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method orderList.place.oco requires parameters but none provided: %v", []string{"symbol", "side", "quantity", "aboveType", "belowType"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for TRADE request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.OrderListPlaceOcoResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendOrderListPlaceOto sends a orderList.place.oto request using typed request/response structs
// Authentication required: TRADE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendOrderListPlaceOto(ctx context.Context, request *models.OrderListPlaceOtoRequest, responseHandler func(*models.OrderListPlaceOtoResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "orderList.place.oto"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method orderList.place.oto requires parameters but none provided: %v", []string{"symbol", "workingType", "workingSide", "workingPrice", "workingQuantity", "pendingType", "pendingSide", "pendingQuantity"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for TRADE request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.OrderListPlaceOtoResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendOrderListPlaceOtoco sends a orderList.place.otoco request using typed request/response structs
// Authentication required: TRADE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendOrderListPlaceOtoco(ctx context.Context, request *models.OrderListPlaceOtocoRequest, responseHandler func(*models.OrderListPlaceOtocoResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "orderList.place.otoco"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method orderList.place.otoco requires parameters but none provided: %v", []string{"symbol", "workingType", "workingSide", "workingPrice", "workingQuantity", "pendingSide", "pendingQuantity", "pendingAboveType"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for TRADE request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.OrderListPlaceOtocoResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendOrderListStatus sends a orderList.status request using typed request/response structs
// Authentication required: USER_DATA
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendOrderListStatus(ctx context.Context, request *models.OrderListStatusRequest, responseHandler func(*models.OrderListStatusResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "orderList.status"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for USER_DATA request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.OrderListStatusResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendOrderPlace sends a order.place request using typed request/response structs
// Authentication required: TRADE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendOrderPlace(ctx context.Context, request *models.OrderPlaceRequest, responseHandler func(*models.OrderPlaceResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "order.place"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method order.place requires parameters but none provided: %v", []string{"symbol", "side", "type"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for TRADE request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.OrderPlaceResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendOrderStatus sends a order.status request using typed request/response structs
// Authentication required: USER_DATA
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendOrderStatus(ctx context.Context, request *models.OrderStatusRequest, responseHandler func(*models.OrderStatusResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "order.status"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method order.status requires parameters but none provided: %v", []string{"symbol"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for USER_DATA request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeUserData); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.OrderStatusResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendOrderTest sends a order.test request using typed request/response structs
// Authentication required: TRADE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendOrderTest(ctx context.Context, request *models.OrderTestRequest, responseHandler func(*models.OrderTestResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "order.test"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for TRADE request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.OrderTestResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendPing sends a ping request using typed request/response structs
// Authentication required: NONE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendPing(ctx context.Context, request *models.PingRequest, responseHandler func(*models.PingResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "ping"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.PingResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendSessionLogon sends a session.logon request using typed request/response structs
// Authentication required: SIGNED
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendSessionLogon(ctx context.Context, request *models.SessionLogonRequest, responseHandler func(*models.SessionLogonResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "session.logon"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for SIGNED request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeSigned); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeSigned); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.SessionLogonResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendSessionLogout sends a session.logout request using typed request/response structs
// Authentication required: NONE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendSessionLogout(ctx context.Context, request *models.SessionLogoutRequest, responseHandler func(*models.SessionLogoutResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "session.logout"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.SessionLogoutResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendSessionStatus sends a session.status request using typed request/response structs
// Authentication required: NONE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendSessionStatus(ctx context.Context, request *models.SessionStatusRequest, responseHandler func(*models.SessionStatusResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "session.status"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.SessionStatusResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendSorOrderPlace sends a sor.order.place request using typed request/response structs
// Authentication required: TRADE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendSorOrderPlace(ctx context.Context, request *models.SorOrderPlaceRequest, responseHandler func(*models.SorOrderPlaceResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "sor.order.place"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method sor.order.place requires parameters but none provided: %v", []string{"symbol", "side", "type", "quantity"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for TRADE request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.SorOrderPlaceResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendSorOrderTest sends a sor.order.test request using typed request/response structs
// Authentication required: TRADE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendSorOrderTest(ctx context.Context, request *models.SorOrderTestRequest, responseHandler func(*models.SorOrderTestResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "sor.order.test"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for TRADE request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeTrade); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.SorOrderTestResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendTicker sends a ticker request using typed request/response structs
// Authentication required: NONE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendTicker(ctx context.Context, request *models.TickerRequest, responseHandler func(*models.TickerResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "ticker"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.TickerResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendTicker24hr sends a ticker.24hr request using typed request/response structs
// Authentication required: NONE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendTicker24hr(ctx context.Context, request *models.Ticker24hrRequest, responseHandler func(*models.Ticker24hrResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "ticker.24hr"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.Ticker24hrResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendTickerBook sends a ticker.book request using typed request/response structs
// Authentication required: NONE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendTickerBook(ctx context.Context, request *models.TickerBookRequest, responseHandler func(*models.TickerBookResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "ticker.book"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.TickerBookResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendTickerPrice sends a ticker.price request using typed request/response structs
// Authentication required: NONE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendTickerPrice(ctx context.Context, request *models.TickerPriceRequest, responseHandler func(*models.TickerPriceResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "ticker.price"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.TickerPriceResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendTickerTradingDay sends a ticker.tradingDay request using typed request/response structs
// Authentication required: NONE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendTickerTradingDay(ctx context.Context, request *models.TickerTradingDayRequest, responseHandler func(*models.TickerTradingDayResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "ticker.tradingDay"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.TickerTradingDayResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendTime sends a time request using typed request/response structs
// Authentication required: NONE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendTime(ctx context.Context, request *models.TimeRequest, responseHandler func(*models.TimeResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "time"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.TimeResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendTradesAggregate sends a trades.aggregate request using typed request/response structs
// Authentication required: NONE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendTradesAggregate(ctx context.Context, request *models.TradesAggregateRequest, responseHandler func(*models.TradesAggregateResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "trades.aggregate"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method trades.aggregate requires parameters but none provided: %v", []string{"symbol"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.TradesAggregateResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendTradesHistorical sends a trades.historical request using typed request/response structs
// Authentication required: NONE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendTradesHistorical(ctx context.Context, request *models.TradesHistoricalRequest, responseHandler func(*models.TradesHistoricalResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "trades.historical"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method trades.historical requires parameters but none provided: %v", []string{"symbol"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.TradesHistoricalResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendTradesRecent sends a trades.recent request using typed request/response structs
// Authentication required: NONE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendTradesRecent(ctx context.Context, request *models.TradesRecentRequest, responseHandler func(*models.TradesRecentResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "trades.recent"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method trades.recent requires parameters but none provided: %v", []string{"symbol"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.TradesRecentResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendUiKlines sends a uiKlines request using typed request/response structs
// Authentication required: NONE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendUiKlines(ctx context.Context, request *models.UiKlinesRequest, responseHandler func(*models.UiKlinesResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "uiKlines"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method uiKlines requires parameters but none provided: %v", []string{"symbol", "interval"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.UiKlinesResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendUserDataStreamPing sends a userDataStream.ping request using typed request/response structs
// Authentication required: USER_STREAM
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendUserDataStreamPing(ctx context.Context, request *models.UserDataStreamPingRequest, responseHandler func(*models.UserDataStreamPingResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "userDataStream.ping"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method userDataStream.ping requires parameters but none provided: %v", []string{"listenKey"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for USER_STREAM request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeUserStream); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeUserStream); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.UserDataStreamPingResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendUserDataStreamStart sends a userDataStream.start request using typed request/response structs
// Authentication required: USER_STREAM
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendUserDataStreamStart(ctx context.Context, request *models.UserDataStreamStartRequest, responseHandler func(*models.UserDataStreamStartResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "userDataStream.start"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for USER_STREAM request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeUserStream); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeUserStream); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.UserDataStreamStartResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendUserDataStreamStop sends a userDataStream.stop request using typed request/response structs
// Authentication required: USER_STREAM
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendUserDataStreamStop(ctx context.Context, request *models.UserDataStreamStopRequest, responseHandler func(*models.UserDataStreamStopResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "userDataStream.stop"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method userDataStream.stop requires parameters but none provided: %v", []string{"listenKey"})
	}

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// Get authentication from context or fall back to client auth
	var auth *Auth
	if contextAuth, ok := ctx.Value(ContextBinanceAuth).(Auth); ok {
		auth = &contextAuth
	} else if c.auth != nil {
		auth = c.auth
	} else {
		return fmt.Errorf("authentication required for USER_STREAM request but no auth provided in context or client")
	}

	// Create signer and sign the request parameters
	signer := NewRequestSigner(auth)
	if params, ok := requestMap["params"].(map[string]interface{}); ok {
		if err := signer.SignRequest(params, AuthTypeUserStream); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	} else {
		// Create params if it doesn't exist
		params := make(map[string]interface{})
		if err := signer.SignRequest(params, AuthTypeUserStream); err != nil {
			return fmt.Errorf("failed to sign request: %w", err)
		}
		requestMap["params"] = params
	}

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.UserDataStreamStopResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendUserDataStreamSubscribe sends a userDataStream.subscribe request using typed request/response structs
// Authentication required: USER_STREAM
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendUserDataStreamSubscribe(ctx context.Context, request *models.UserDataStreamSubscribeRequest, responseHandler func(*models.UserDataStreamSubscribeResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "userDataStream.subscribe"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// userDataStream methods don't need params - authentication is handled at the WebSocket connection level

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.UserDataStreamSubscribeResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendUserDataStreamUnsubscribe sends a userDataStream.unsubscribe request using typed request/response structs
// Authentication required: USER_STREAM
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendUserDataStreamUnsubscribe(ctx context.Context, request *models.UserDataStreamUnsubscribeRequest, responseHandler func(*models.UserDataStreamUnsubscribeResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "userDataStream.unsubscribe"

	// Convert struct to map for WebSocket sending
	requestMap, err := structToMap(request)
	if err != nil {
		return fmt.Errorf("failed to convert request to map: %w", err)
	}

	// userDataStream methods don't need params - authentication is handled at the WebSocket connection level

	// Register typed response handler with automatic JSON parsing
	RegisterTypedResponseHandler[models.UserDataStreamUnsubscribeResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// HandleBalanceUpdate registers a handler for Balance Update Event events
// This method allows you to handle real-time Balance Update Event events from the WebSocket stream
func (c *Client) HandleBalanceUpdate(handler func(*models.BalanceUpdate) error) {
	c.eventHandler.RegisterHandler("balanceUpdate", func(data interface{}) error {
		// Parse the event data - handle both nested and direct event structures
		var event models.BalanceUpdate
		
		if jsonData, ok := data.([]byte); ok {
			// Direct JSON data parsing
			if err := json.Unmarshal(jsonData, &event); err != nil {
				return fmt.Errorf("failed to parse Balance Update Event event: %w", err)
			}
		} else if mapData, ok := data.(map[string]interface{}); ok {
			// Map data - check if this is the nested event object or the full message
			var eventDataToUnmarshal interface{}
			
			// Check if this map contains an "event" field (nested structure)
			if _, hasEvent := mapData["event"]; hasEvent {
				// This is a full message with nested event object
				// Use the entire message structure for parsing
				eventDataToUnmarshal = mapData
			} else {
				// This might be the event data itself
				eventDataToUnmarshal = mapData
			}
			
			// Convert to JSON and back to struct
			jsonBytes, err := json.Marshal(eventDataToUnmarshal)
			if err != nil {
				return fmt.Errorf("failed to marshal Balance Update Event event data: %w", err)
			}
			if err := json.Unmarshal(jsonBytes, &event); err != nil {
				return fmt.Errorf("failed to parse Balance Update Event event: %w", err)
			}
		} else {
			return fmt.Errorf("unexpected data type for Balance Update Event event: %T", data)
		}
		
		// Call the user-provided handler
		return handler(&event)
	})
}


// HandleEventStreamTerminated registers a handler for Event Stream Terminated Event events
// This method allows you to handle real-time Event Stream Terminated Event events from the WebSocket stream
func (c *Client) HandleEventStreamTerminated(handler func(*models.EventStreamTerminated) error) {
	c.eventHandler.RegisterHandler("eventStreamTerminated", func(data interface{}) error {
		// Parse the event data - handle both nested and direct event structures
		var event models.EventStreamTerminated
		
		if jsonData, ok := data.([]byte); ok {
			// Direct JSON data parsing
			if err := json.Unmarshal(jsonData, &event); err != nil {
				return fmt.Errorf("failed to parse Event Stream Terminated Event event: %w", err)
			}
		} else if mapData, ok := data.(map[string]interface{}); ok {
			// Map data - check if this is the nested event object or the full message
			var eventDataToUnmarshal interface{}
			
			// Check if this map contains an "event" field (nested structure)
			if _, hasEvent := mapData["event"]; hasEvent {
				// This is a full message with nested event object
				// Use the entire message structure for parsing
				eventDataToUnmarshal = mapData
			} else {
				// This might be the event data itself
				eventDataToUnmarshal = mapData
			}
			
			// Convert to JSON and back to struct
			jsonBytes, err := json.Marshal(eventDataToUnmarshal)
			if err != nil {
				return fmt.Errorf("failed to marshal Event Stream Terminated Event event data: %w", err)
			}
			if err := json.Unmarshal(jsonBytes, &event); err != nil {
				return fmt.Errorf("failed to parse Event Stream Terminated Event event: %w", err)
			}
		} else {
			return fmt.Errorf("unexpected data type for Event Stream Terminated Event event: %T", data)
		}
		
		// Call the user-provided handler
		return handler(&event)
	})
}


// HandleExecutionReport registers a handler for Execution Report Event events
// This method allows you to handle real-time Execution Report Event events from the WebSocket stream
func (c *Client) HandleExecutionReport(handler func(*models.ExecutionReport) error) {
	c.eventHandler.RegisterHandler("executionReport", func(data interface{}) error {
		// Parse the event data - handle both nested and direct event structures
		var event models.ExecutionReport
		
		if jsonData, ok := data.([]byte); ok {
			// Direct JSON data parsing
			if err := json.Unmarshal(jsonData, &event); err != nil {
				return fmt.Errorf("failed to parse Execution Report Event event: %w", err)
			}
		} else if mapData, ok := data.(map[string]interface{}); ok {
			// Map data - check if this is the nested event object or the full message
			var eventDataToUnmarshal interface{}
			
			// Check if this map contains an "event" field (nested structure)
			if _, hasEvent := mapData["event"]; hasEvent {
				// This is a full message with nested event object
				// Use the entire message structure for parsing
				eventDataToUnmarshal = mapData
			} else {
				// This might be the event data itself
				eventDataToUnmarshal = mapData
			}
			
			// Convert to JSON and back to struct
			jsonBytes, err := json.Marshal(eventDataToUnmarshal)
			if err != nil {
				return fmt.Errorf("failed to marshal Execution Report Event event data: %w", err)
			}
			if err := json.Unmarshal(jsonBytes, &event); err != nil {
				return fmt.Errorf("failed to parse Execution Report Event event: %w", err)
			}
		} else {
			return fmt.Errorf("unexpected data type for Execution Report Event event: %T", data)
		}
		
		// Call the user-provided handler
		return handler(&event)
	})
}


// HandleExternalLockUpdate registers a handler for External Lock Update Event events
// This method allows you to handle real-time External Lock Update Event events from the WebSocket stream
func (c *Client) HandleExternalLockUpdate(handler func(*models.ExternalLockUpdate) error) {
	c.eventHandler.RegisterHandler("externalLockUpdate", func(data interface{}) error {
		// Parse the event data - handle both nested and direct event structures
		var event models.ExternalLockUpdate
		
		if jsonData, ok := data.([]byte); ok {
			// Direct JSON data parsing
			if err := json.Unmarshal(jsonData, &event); err != nil {
				return fmt.Errorf("failed to parse External Lock Update Event event: %w", err)
			}
		} else if mapData, ok := data.(map[string]interface{}); ok {
			// Map data - check if this is the nested event object or the full message
			var eventDataToUnmarshal interface{}
			
			// Check if this map contains an "event" field (nested structure)
			if _, hasEvent := mapData["event"]; hasEvent {
				// This is a full message with nested event object
				// Use the entire message structure for parsing
				eventDataToUnmarshal = mapData
			} else {
				// This might be the event data itself
				eventDataToUnmarshal = mapData
			}
			
			// Convert to JSON and back to struct
			jsonBytes, err := json.Marshal(eventDataToUnmarshal)
			if err != nil {
				return fmt.Errorf("failed to marshal External Lock Update Event event data: %w", err)
			}
			if err := json.Unmarshal(jsonBytes, &event); err != nil {
				return fmt.Errorf("failed to parse External Lock Update Event event: %w", err)
			}
		} else {
			return fmt.Errorf("unexpected data type for External Lock Update Event event: %T", data)
		}
		
		// Call the user-provided handler
		return handler(&event)
	})
}


// HandleListStatus registers a handler for List Status Event events
// This method allows you to handle real-time List Status Event events from the WebSocket stream
func (c *Client) HandleListStatus(handler func(*models.ListStatus) error) {
	c.eventHandler.RegisterHandler("listStatus", func(data interface{}) error {
		// Parse the event data - handle both nested and direct event structures
		var event models.ListStatus
		
		if jsonData, ok := data.([]byte); ok {
			// Direct JSON data parsing
			if err := json.Unmarshal(jsonData, &event); err != nil {
				return fmt.Errorf("failed to parse List Status Event event: %w", err)
			}
		} else if mapData, ok := data.(map[string]interface{}); ok {
			// Map data - check if this is the nested event object or the full message
			var eventDataToUnmarshal interface{}
			
			// Check if this map contains an "event" field (nested structure)
			if _, hasEvent := mapData["event"]; hasEvent {
				// This is a full message with nested event object
				// Use the entire message structure for parsing
				eventDataToUnmarshal = mapData
			} else {
				// This might be the event data itself
				eventDataToUnmarshal = mapData
			}
			
			// Convert to JSON and back to struct
			jsonBytes, err := json.Marshal(eventDataToUnmarshal)
			if err != nil {
				return fmt.Errorf("failed to marshal List Status Event event data: %w", err)
			}
			if err := json.Unmarshal(jsonBytes, &event); err != nil {
				return fmt.Errorf("failed to parse List Status Event event: %w", err)
			}
		} else {
			return fmt.Errorf("unexpected data type for List Status Event event: %T", data)
		}
		
		// Call the user-provided handler
		return handler(&event)
	})
}


// HandleListenKeyExpired registers a handler for Listen Key Expired Event events
// This method allows you to handle real-time Listen Key Expired Event events from the WebSocket stream
func (c *Client) HandleListenKeyExpired(handler func(*models.ListenKeyExpired) error) {
	c.eventHandler.RegisterHandler("listenKeyExpired", func(data interface{}) error {
		// Parse the event data - handle both nested and direct event structures
		var event models.ListenKeyExpired
		
		if jsonData, ok := data.([]byte); ok {
			// Direct JSON data parsing
			if err := json.Unmarshal(jsonData, &event); err != nil {
				return fmt.Errorf("failed to parse Listen Key Expired Event event: %w", err)
			}
		} else if mapData, ok := data.(map[string]interface{}); ok {
			// Map data - check if this is the nested event object or the full message
			var eventDataToUnmarshal interface{}
			
			// Check if this map contains an "event" field (nested structure)
			if _, hasEvent := mapData["event"]; hasEvent {
				// This is a full message with nested event object
				// Use the entire message structure for parsing
				eventDataToUnmarshal = mapData
			} else {
				// This might be the event data itself
				eventDataToUnmarshal = mapData
			}
			
			// Convert to JSON and back to struct
			jsonBytes, err := json.Marshal(eventDataToUnmarshal)
			if err != nil {
				return fmt.Errorf("failed to marshal Listen Key Expired Event event data: %w", err)
			}
			if err := json.Unmarshal(jsonBytes, &event); err != nil {
				return fmt.Errorf("failed to parse Listen Key Expired Event event: %w", err)
			}
		} else {
			return fmt.Errorf("unexpected data type for Listen Key Expired Event event: %T", data)
		}
		
		// Call the user-provided handler
		return handler(&event)
	})
}


// HandleOutboundAccountPosition registers a handler for Outbound Account Position Event events
// This method allows you to handle real-time Outbound Account Position Event events from the WebSocket stream
func (c *Client) HandleOutboundAccountPosition(handler func(*models.OutboundAccountPosition) error) {
	c.eventHandler.RegisterHandler("outboundAccountPosition", func(data interface{}) error {
		// Parse the event data - handle both nested and direct event structures
		var event models.OutboundAccountPosition
		
		if jsonData, ok := data.([]byte); ok {
			// Direct JSON data parsing
			if err := json.Unmarshal(jsonData, &event); err != nil {
				return fmt.Errorf("failed to parse Outbound Account Position Event event: %w", err)
			}
		} else if mapData, ok := data.(map[string]interface{}); ok {
			// Map data - check if this is the nested event object or the full message
			var eventDataToUnmarshal interface{}
			
			// Check if this map contains an "event" field (nested structure)
			if _, hasEvent := mapData["event"]; hasEvent {
				// This is a full message with nested event object
				// Use the entire message structure for parsing
				eventDataToUnmarshal = mapData
			} else {
				// This might be the event data itself
				eventDataToUnmarshal = mapData
			}
			
			// Convert to JSON and back to struct
			jsonBytes, err := json.Marshal(eventDataToUnmarshal)
			if err != nil {
				return fmt.Errorf("failed to marshal Outbound Account Position Event event data: %w", err)
			}
			if err := json.Unmarshal(jsonBytes, &event); err != nil {
				return fmt.Errorf("failed to parse Outbound Account Position Event event: %w", err)
			}
		} else {
			return fmt.Errorf("unexpected data type for Outbound Account Position Event event: %T", data)
		}
		
		// Call the user-provided handler
		return handler(&event)
	})
}


// RegisterTypedResponseHandler registers a typed response handler for a request ID
func RegisterTypedResponseHandler[T any](c *Client, requestID string, handler func(*T, error) error) {
	c.responseHandlers.Store(requestID, ResponseHandler{
		RequestID: requestID,
		Handler: func(data []byte, err error) error {
			if err != nil {
				return handler(nil, err)
			}
			
			// Parse the response into the specified type
			var response T
			if err := json.Unmarshal(data, &response); err != nil {
				return handler(nil, fmt.Errorf("failed to parse response: %w", err))
			}
			
			return handler(&response, nil)
		},
	})
}

// structToMap converts a struct to a map[string]interface{} while preserving number precision
func structToMap(v interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	
	// Use a decoder that preserves number precision
	decoder := json.NewDecoder(strings.NewReader(string(data)))
	decoder.UseNumber()
	
	var result map[string]interface{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	
	// Convert json.Number to appropriate types
	convertJSONNumbers(result)
	
	return result, nil
}

// convertJSONNumbers recursively converts json.Number values to appropriate types
func convertJSONNumbers(m map[string]interface{}) {
	for k, v := range m {
		switch val := v.(type) {
		case json.Number:
			// Try to convert to int64 first, then fall back to float64
			if intVal, err := val.Int64(); err == nil {
				m[k] = intVal
			} else if floatVal, err := val.Float64(); err == nil {
				m[k] = floatVal
			} else {
				// Keep as string if conversion fails
				m[k] = string(val)
			}
		case map[string]interface{}:
			// Recursively convert nested maps
			convertJSONNumbers(val)
		}
	}
}



