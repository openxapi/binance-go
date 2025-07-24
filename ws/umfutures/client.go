package umfutures
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
	"github.com/openxapi/binance-go/ws/umfutures/models"
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
		URL:         "wss://ws-fapi.binance.com/ws-fapi/v1",
		Host:        "ws-fapi.binance.com",
		Pathname:    "/ws-fapi/v1",
		Protocol:    "wss",
		Title:       "Binance Mainnet Server",
		Summary:     "Binance Umfutures WebSocket API Server (Mainnet)",
		Description: "WebSocket server for binance exchange umfutures API (mainnet environment)",
		Active:      false,
	}
	sm.servers["testnet1"] = &ServerInfo{
		Name:        "testnet1",
		URL:         "wss://testnet.binancefuture.com/ws-fapi/v1",
		Host:        "testnet.binancefuture.com",
		Pathname:    "/ws-fapi/v1",
		Protocol:    "wss",
		Title:       "Binance Testnet Server",
		Summary:     "Binance Umfutures WebSocket API Server (Testnet)",
		Description: "WebSocket server for binance exchange umfutures API (testnet environment)",
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

// Client represents a high-performance WebSocket client for 
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
		handlers:      eventHandlers{},        // Initialize event handlers registry
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
	
	// Start the message reading loop
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

// SendAccountBalance sends a account.balance request using typed request/response structs
// Authentication required: USER_DATA
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendAccountBalance(ctx context.Context, request *models.AccountBalanceRequest, responseHandler func(*models.AccountBalanceResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "account.balance"

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
	RegisterTypedResponseHandler[models.AccountBalanceResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendAccountPosition sends a account.position request using typed request/response structs
// Authentication required: USER_DATA
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendAccountPosition(ctx context.Context, request *models.AccountPositionRequest, responseHandler func(*models.AccountPositionResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "account.position"

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
	RegisterTypedResponseHandler[models.AccountPositionResponse](c, reqID, responseHandler)

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
		return fmt.Errorf("method order.cancel requires parameters but none provided: %v", []string{"symbol"})
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


// SendOrderModify sends a order.modify request using typed request/response structs
// Authentication required: TRADE
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendOrderModify(ctx context.Context, request *models.OrderModifyRequest, responseHandler func(*models.OrderModifyResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "order.modify"

	// Validate required parameters (excluding auth parameters that are auto-added)
	if request.Params == nil {
		return fmt.Errorf("method order.modify requires parameters but none provided: %v", []string{"symbol", "side", "quantity", "price"})
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
	RegisterTypedResponseHandler[models.OrderModifyResponse](c, reqID, responseHandler)

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


// SendV2AccountBalance sends a v2/account.balance request using typed request/response structs
// Authentication required: USER_DATA
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendV2AccountBalance(ctx context.Context, request *models.V2AccountBalanceRequest, responseHandler func(*models.V2AccountBalanceResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "v2/account.balance"

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
	RegisterTypedResponseHandler[models.V2AccountBalanceResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendV2AccountPosition sends a v2/account.position request using typed request/response structs
// Authentication required: USER_DATA
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendV2AccountPosition(ctx context.Context, request *models.V2AccountPositionRequest, responseHandler func(*models.V2AccountPositionResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "v2/account.position"

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
	RegisterTypedResponseHandler[models.V2AccountPositionResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// SendV2AccountStatus sends a v2/account.status request using typed request/response structs
// Authentication required: USER_DATA
// If request.Id is empty, a new request ID will be generated automatically
func (c *Client) SendV2AccountStatus(ctx context.Context, request *models.V2AccountStatusRequest, responseHandler func(*models.V2AccountStatusResponse, error) error) error {
	// Use existing request ID or generate a new one
	var reqID string
	if request.Id != "" {
		reqID = request.Id
	} else {
		reqID = GenerateRequestID()
		request.Id = reqID
	}
	request.Method = "v2/account.status"

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
	RegisterTypedResponseHandler[models.V2AccountStatusResponse](c, reqID, responseHandler)

	// Send request
	return c.sendRequest(requestMap)
}


// HandleAccountConfigUpdate registers a handler for Account Configuration Update Event events
// This method allows you to handle real-time Account Configuration Update Event events from the WebSocket stream
func (c *Client) HandleAccountConfigUpdate(handler func(*models.AccountConfigUpdate) error) {
	c.eventHandler.RegisterHandler("accountConfigUpdate", func(data interface{}) error {
		// Parse the event data - handle both nested and direct event structures
		var event models.AccountConfigUpdate
		
		if jsonData, ok := data.([]byte); ok {
			// Direct JSON data parsing
			if err := json.Unmarshal(jsonData, &event); err != nil {
				return fmt.Errorf("failed to parse Account Configuration Update Event event: %w", err)
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
				return fmt.Errorf("failed to marshal Account Configuration Update Event event data: %w", err)
			}
			if err := json.Unmarshal(jsonBytes, &event); err != nil {
				return fmt.Errorf("failed to parse Account Configuration Update Event event: %w", err)
			}
		} else {
			return fmt.Errorf("unexpected data type for Account Configuration Update Event event: %T", data)
		}
		
		// Call the user-provided handler
		return handler(&event)
	})
}


// HandleAccountUpdate registers a handler for Account Update Event events
// This method allows you to handle real-time Account Update Event events from the WebSocket stream
func (c *Client) HandleAccountUpdate(handler func(*models.AccountUpdate) error) {
	c.eventHandler.RegisterHandler("accountUpdate", func(data interface{}) error {
		// Parse the event data - handle both nested and direct event structures
		var event models.AccountUpdate
		
		if jsonData, ok := data.([]byte); ok {
			// Direct JSON data parsing
			if err := json.Unmarshal(jsonData, &event); err != nil {
				return fmt.Errorf("failed to parse Account Update Event event: %w", err)
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
				return fmt.Errorf("failed to marshal Account Update Event event data: %w", err)
			}
			if err := json.Unmarshal(jsonBytes, &event); err != nil {
				return fmt.Errorf("failed to parse Account Update Event event: %w", err)
			}
		} else {
			return fmt.Errorf("unexpected data type for Account Update Event event: %T", data)
		}
		
		// Call the user-provided handler
		return handler(&event)
	})
}


// HandleConditionalOrderTriggerReject registers a handler for Conditional Order Trigger Reject Event events
// This method allows you to handle real-time Conditional Order Trigger Reject Event events from the WebSocket stream
func (c *Client) HandleConditionalOrderTriggerReject(handler func(*models.ConditionalOrderTriggerReject) error) {
	c.eventHandler.RegisterHandler("conditionalOrderTriggerReject", func(data interface{}) error {
		// Parse the event data - handle both nested and direct event structures
		var event models.ConditionalOrderTriggerReject
		
		if jsonData, ok := data.([]byte); ok {
			// Direct JSON data parsing
			if err := json.Unmarshal(jsonData, &event); err != nil {
				return fmt.Errorf("failed to parse Conditional Order Trigger Reject Event event: %w", err)
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
				return fmt.Errorf("failed to marshal Conditional Order Trigger Reject Event event data: %w", err)
			}
			if err := json.Unmarshal(jsonBytes, &event); err != nil {
				return fmt.Errorf("failed to parse Conditional Order Trigger Reject Event event: %w", err)
			}
		} else {
			return fmt.Errorf("unexpected data type for Conditional Order Trigger Reject Event event: %T", data)
		}
		
		// Call the user-provided handler
		return handler(&event)
	})
}


// HandleGridUpdate registers a handler for Grid Update Event events
// This method allows you to handle real-time Grid Update Event events from the WebSocket stream
func (c *Client) HandleGridUpdate(handler func(*models.GridUpdate) error) {
	c.eventHandler.RegisterHandler("gridUpdate", func(data interface{}) error {
		// Parse the event data - handle both nested and direct event structures
		var event models.GridUpdate
		
		if jsonData, ok := data.([]byte); ok {
			// Direct JSON data parsing
			if err := json.Unmarshal(jsonData, &event); err != nil {
				return fmt.Errorf("failed to parse Grid Update Event event: %w", err)
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
				return fmt.Errorf("failed to marshal Grid Update Event event data: %w", err)
			}
			if err := json.Unmarshal(jsonBytes, &event); err != nil {
				return fmt.Errorf("failed to parse Grid Update Event event: %w", err)
			}
		} else {
			return fmt.Errorf("unexpected data type for Grid Update Event event: %T", data)
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


// HandleMarginCall registers a handler for Margin Call Event events
// This method allows you to handle real-time Margin Call Event events from the WebSocket stream
func (c *Client) HandleMarginCall(handler func(*models.MarginCall) error) {
	c.eventHandler.RegisterHandler("marginCall", func(data interface{}) error {
		// Parse the event data - handle both nested and direct event structures
		var event models.MarginCall
		
		if jsonData, ok := data.([]byte); ok {
			// Direct JSON data parsing
			if err := json.Unmarshal(jsonData, &event); err != nil {
				return fmt.Errorf("failed to parse Margin Call Event event: %w", err)
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
				return fmt.Errorf("failed to marshal Margin Call Event event data: %w", err)
			}
			if err := json.Unmarshal(jsonBytes, &event); err != nil {
				return fmt.Errorf("failed to parse Margin Call Event event: %w", err)
			}
		} else {
			return fmt.Errorf("unexpected data type for Margin Call Event event: %T", data)
		}
		
		// Call the user-provided handler
		return handler(&event)
	})
}


// HandleOrderTradeUpdate registers a handler for Order Trade Update Event events
// This method allows you to handle real-time Order Trade Update Event events from the WebSocket stream
func (c *Client) HandleOrderTradeUpdate(handler func(*models.OrderTradeUpdate) error) {
	c.eventHandler.RegisterHandler("orderTradeUpdate", func(data interface{}) error {
		// Parse the event data - handle both nested and direct event structures
		var event models.OrderTradeUpdate
		
		if jsonData, ok := data.([]byte); ok {
			// Direct JSON data parsing
			if err := json.Unmarshal(jsonData, &event); err != nil {
				return fmt.Errorf("failed to parse Order Trade Update Event event: %w", err)
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
				return fmt.Errorf("failed to marshal Order Trade Update Event event data: %w", err)
			}
			if err := json.Unmarshal(jsonBytes, &event); err != nil {
				return fmt.Errorf("failed to parse Order Trade Update Event event: %w", err)
			}
		} else {
			return fmt.Errorf("unexpected data type for Order Trade Update Event event: %T", data)
		}
		
		// Call the user-provided handler
		return handler(&event)
	})
}


// HandleStrategyUpdate registers a handler for Strategy Update Event events
// This method allows you to handle real-time Strategy Update Event events from the WebSocket stream
func (c *Client) HandleStrategyUpdate(handler func(*models.StrategyUpdate) error) {
	c.eventHandler.RegisterHandler("strategyUpdate", func(data interface{}) error {
		// Parse the event data - handle both nested and direct event structures
		var event models.StrategyUpdate
		
		if jsonData, ok := data.([]byte); ok {
			// Direct JSON data parsing
			if err := json.Unmarshal(jsonData, &event); err != nil {
				return fmt.Errorf("failed to parse Strategy Update Event event: %w", err)
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
				return fmt.Errorf("failed to marshal Strategy Update Event event data: %w", err)
			}
			if err := json.Unmarshal(jsonBytes, &event); err != nil {
				return fmt.Errorf("failed to parse Strategy Update Event event: %w", err)
			}
		} else {
			return fmt.Errorf("unexpected data type for Strategy Update Event event: %T", data)
		}
		
		// Call the user-provided handler
		return handler(&event)
	})
}


// HandleTradeLite registers a handler for Trade Lite Event events
// This method allows you to handle real-time Trade Lite Event events from the WebSocket stream
func (c *Client) HandleTradeLite(handler func(*models.TradeLite) error) {
	c.eventHandler.RegisterHandler("tradeLite", func(data interface{}) error {
		// Parse the event data - handle both nested and direct event structures
		var event models.TradeLite
		
		if jsonData, ok := data.([]byte); ok {
			// Direct JSON data parsing
			if err := json.Unmarshal(jsonData, &event); err != nil {
				return fmt.Errorf("failed to parse Trade Lite Event event: %w", err)
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
				return fmt.Errorf("failed to marshal Trade Lite Event event data: %w", err)
			}
			if err := json.Unmarshal(jsonBytes, &event); err != nil {
				return fmt.Errorf("failed to parse Trade Lite Event event: %w", err)
			}
		} else {
			return fmt.Errorf("unexpected data type for Trade Lite Event event: %T", data)
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



