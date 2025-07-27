package umfuturesstreams
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
	"github.com/openxapi/binance-go/ws/umfutures-streams/models"
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
		URL:         "wss://fstream.binance.com/{streamPath}",
		Host:        "fstream.binance.com",
		Pathname:    "/{streamPath}",
		Protocol:    "wss",
		Title:       "Binance USD-S Margined Futures Server",
		Summary:     "Binance USD-S Margined Futures WebSocket Streams Server (Mainnet)",
		Description: "WebSocket server for binance exchange USD-S margined futures market data streams (mainnet environment)",
		Active:      false,
	}
	sm.servers["testnet1"] = &ServerInfo{
		Name:        "testnet1",
		URL:         "wss://fstream.binancefuture.com/{streamPath}",
		Host:        "fstream.binancefuture.com",
		Pathname:    "/{streamPath}",
		Protocol:    "wss",
		Title:       "Binance USD-S Margined Futures Testnet Server",
		Summary:     "Binance USD-S Margined Futures WebSocket Streams Server (Testnet)",
		Description: "WebSocket server for binance exchange USD-S margined futures market data streams (testnet environment)",
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
	Id      string `json:"id"`      // Request ID that caused the error
}

// Error implements the error interface
func (e APIError) Error() string {
	return fmt.Sprintf("binance api error: status=%d, code=%d, message=%s, id=%s", e.Status, e.Code, e.Message, e.Id)
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
	RequestId string
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

// Client represents a high-performance WebSocket client for Binance USD-S Margined Futures WebSocket Streams
type Client struct {
	conn             *websocket.Conn
	connMu           sync.RWMutex     // Protects connection access
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
	// For streams modules, use ConnectToSingleStreams by default
	return c.ConnectToSingleStreams(ctx, "")
}

// ConnectToServer establishes a WebSocket connection to a specific server
func (c *Client) ConnectToServer(ctx context.Context, serverName string) error {
	if err := c.SetActiveServer(serverName); err != nil {
		return fmt.Errorf("failed to set active server: %w", err)
	}
	
	return c.Connect(ctx)
}

// ConnectWithVariables establishes a WebSocket connection using provided server variables
// This method resolves server URL template variables like {streamPath}
func (c *Client) ConnectWithVariables(ctx context.Context, streamPath string) error {
	activeServer := c.serverManager.GetActiveServer()
	if activeServer == nil {
		return fmt.Errorf("no active server configured")
	}
	
	// Resolve the URL with the provided variables
	variables := map[string]string{
		"streamPath": streamPath,
	}
	
	resolvedURL, err := c.serverManager.ResolveServerURL(activeServer.Name, variables)
	if err != nil {
		return fmt.Errorf("failed to resolve server URL: %w", err)
	}
	
	u, err := url.Parse(resolvedURL)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}

	dialer := websocket.DefaultDialer
	dialer.HandshakeTimeout = 10 * time.Second

	conn, _, err := dialer.DialContext(ctx, u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w", resolvedURL, err)
	}

	c.conn = conn
	c.isConnected = true
	go c.readMessages()
	return nil
}

// ConnectToServerWithVariables establishes a WebSocket connection to a specific server using provided server variables
func (c *Client) ConnectToServerWithVariables(ctx context.Context, serverName string, streamPath string) error {
	if err := c.SetActiveServer(serverName); err != nil {
		return fmt.Errorf("failed to set active server: %w", err)
	}
	
	return c.ConnectWithVariables(ctx, streamPath)
}

// ConnectWithStreamPath establishes a WebSocket connection using the provided streamPath
// This is a convenience method for the streamPath variable
func (c *Client) ConnectWithStreamPath(ctx context.Context, streamPath string) error {
	return c.ConnectWithVariables(ctx, streamPath)
}

// ConnectToServerWithStreamPath establishes a WebSocket connection to a specific server using the provided streamPath
func (c *Client) ConnectToServerWithStreamPath(ctx context.Context, serverName string, streamPath string) error {
	if err := c.SetActiveServer(serverName); err != nil {
		return fmt.Errorf("failed to set active server: %w", err)
	}
	
	return c.ConnectWithStreamPath(ctx, streamPath)
}

// Disconnect closes the WebSocket connection safely and resets state for reconnection
func (c *Client) Disconnect() error {
	// Signal all goroutines to stop first
	c.isConnected = false
	
	// Safely close the done channel only once
	select {
	case <-c.done:
		// Channel already closed, do nothing
	default:
		close(c.done)
	}
	
	// Wait a brief moment for goroutines to see the done signal
	time.Sleep(10 * time.Millisecond)
	
	// Now safely handle the connection with proper locking
	c.connMu.Lock()
	defer c.connMu.Unlock()
	
	var err error
	if c.conn != nil {
		err = c.conn.Close()
		c.conn = nil  // Reset connection to nil for clean reconnection
	}
	
	// Reset connection state for reconnection
	c.done = make(chan struct{})  // Recreate done channel
	
	return err
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
		return c.processStreamMessage(data)
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
			Id     interface{} `json:"id"`
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
				Id:      requestID,
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

	// For stream modules, always try processStreamMessage for any non-request message
		// This handles both standard events (with "e" field), special events (BookTicker, PartialDepth), and combined streams
		return c.processStreamMessage(data)
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
	return c.processStreamMessage(data)
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


// Subscribe to market data streams
func (c *Client) Subscribe(ctx context.Context, streams []string) error {
	if !c.isConnected {
		return fmt.Errorf("websocket not connected")
	}

	request := map[string]interface{}{
		"method": "SUBSCRIBE",
		"params": streams,
		"id":     GenerateRequestID(),
	}

	return c.sendRequest(request)
}

// Unsubscribe from market data streams  
func (c *Client) Unsubscribe(ctx context.Context, streams []string) error {
	if !c.isConnected {
		return fmt.Errorf("websocket not connected")
	}

	request := map[string]interface{}{
		"method": "UNSUBSCRIBE", 
		"params": streams,
		"id":     GenerateRequestID(),
	}

	return c.sendRequest(request)
}

// List active subscriptions
func (c *Client) ListSubscriptions(ctx context.Context) error {
	if !c.isConnected {
		return fmt.Errorf("websocket not connected")
	}

	request := map[string]interface{}{
		"method": "LIST_SUBSCRIPTIONS",
		"id":     GenerateRequestID(),
	}

	return c.sendRequest(request)
}

// ConnectToSingleStreams connects to single stream endpoint with optional timeUnit parameter
func (c *Client) ConnectToSingleStreams(ctx context.Context, timeUnit string) error {
	if c.isConnected {
		return fmt.Errorf("already connected to websocket")
	}
	
	// Set server variable to single stream path
	if err := c.setStreamPath("ws"); err != nil {
		return fmt.Errorf("failed to set stream path: %w", err)
	}
	
	// Build endpoint URL with timeUnit parameter
	endpoint := "/ws"
	if timeUnit != "" {
		endpoint += timeUnit // timeUnit should be formatted like "?timeUnit=MICROSECOND"
	}
	
	return c.connect(ctx, endpoint, false) // false = not combined stream
}

// ConnectToCombinedStreams connects to combined stream endpoint with optional timeUnit parameter
func (c *Client) ConnectToCombinedStreams(ctx context.Context, timeUnit string) error {
	if c.isConnected {
		return fmt.Errorf("already connected to websocket")
	}
	
	// Set server variable to combined stream path
	if err := c.setStreamPath("stream"); err != nil {
		return fmt.Errorf("failed to set stream path: %w", err)
	}
	
	// Build endpoint URL with timeUnit parameter
	endpoint := "/stream"
	if timeUnit != "" {
		endpoint += timeUnit // timeUnit should be formatted like "?timeUnit=MICROSECOND"
	}
	
	return c.connect(ctx, endpoint, true) // true = combined stream
}

// ConnectToSingleStreamsMicrosecond connects to single stream endpoint with microsecond precision
func (c *Client) ConnectToSingleStreamsMicrosecond(ctx context.Context) error {
	return c.ConnectToSingleStreams(ctx, "?timeUnit=MICROSECOND")
}

// ConnectToCombinedStreamsMicrosecond connects to combined stream endpoint with microsecond precision
func (c *Client) ConnectToCombinedStreamsMicrosecond(ctx context.Context) error {
	return c.ConnectToCombinedStreams(ctx, "?timeUnit=MICROSECOND")
}

// setStreamPath sets the server variable for stream path selection
func (c *Client) setStreamPath(streamPath string) error {
	activeServer := c.serverManager.GetActiveServer()
	if activeServer == nil {
		return fmt.Errorf("no active server configured")
	}
	
	// Update the server's pathname to use the specified stream path
	updatedPathname := "/" + streamPath
	return c.serverManager.UpdateServerPathname(activeServer.Name, updatedPathname)
}

// connect establishes a WebSocket connection to a specific endpoint (for umfutures-streams)
func (c *Client) connect(ctx context.Context, endpoint string, isCombined bool) error {
	// Check connection state with proper locking
	c.connMu.RLock()
	isConnected := c.isConnected
	c.connMu.RUnlock()
	
	if isConnected {
		return fmt.Errorf("websocket already connected")
	}
	
	activeServer := c.serverManager.GetActiveServer()
	if activeServer == nil {
		return fmt.Errorf("no active server configured")
	}
	
	// Build the WebSocket URL with the specific endpoint
	// For streams, we connect directly to the endpoint (like /ws or /stream)
	// The Host field should be clean hostname without template variables
	host := activeServer.Host
	if strings.Contains(host, "{streamPath}") {
		// Extract just the hostname part, removing template variables
		host = strings.Split(host, "/")[0]
		host = strings.ReplaceAll(host, "{streamPath}", "")
	}
	serverURL := fmt.Sprintf("%s://%s%s", activeServer.Protocol, host, endpoint)
	
	u, err := url.Parse(serverURL)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}

	dialer := websocket.DefaultDialer
	dialer.HandshakeTimeout = 10 * time.Second

	conn, _, err := dialer.DialContext(ctx, u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w", serverURL, err)
	}

	// Safely assign connection with proper locking
	c.connMu.Lock()
	c.conn = conn
	c.isConnected = true
	c.connMu.Unlock()
	
	// Start appropriate message reading routine
	if isCombined {
		go c.readCombinedStreamMessages()
	} else {
		go c.readSingleStreamMessages()
	}
	
	return nil
}

// readSingleStreamMessages processes messages from single stream connections
func (c *Client) readSingleStreamMessages() {
	defer func() {
		c.isConnected = false
	}()

	for {
		select {
		case <-c.done:
			return
		default:
			// Safely access connection with read lock
			c.connMu.RLock()
			conn := c.conn
			c.connMu.RUnlock()
			
			// Check if connection is nil (race condition protection)
			if conn == nil {
				return
			}
			
			_, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					return
				}
				if strings.Contains(err.Error(), "use of closed network connection") {
					return
				}
				log.Printf("Error reading message: %v", err)
				return
			}

			if err := c.processStreamMessage(message); err != nil {
				log.Printf("Error processing stream message: %v", err)
			}
		}
	}
}

// readCombinedStreamMessages processes messages from combined stream connections
func (c *Client) readCombinedStreamMessages() {
	defer func() {
		c.isConnected = false
	}()

	for {
		select {
		case <-c.done:
			return
		default:
			// Safely access connection with read lock
			c.connMu.RLock()
			conn := c.conn
			c.connMu.RUnlock()
			
			// Check if connection is nil (race condition protection)
			if conn == nil {
				return
			}
			
			_, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					return
				}
				if strings.Contains(err.Error(), "use of closed network connection") {
					return
				}
				log.Printf("Error reading message: %v", err)
				return
			}

			if err := c.processStreamMessage(message); err != nil {
				log.Printf("Error processing stream message: %v", err)
			}
		}
	}
}


// Stream event handler functions
type (
	// AggregateTradeEvent Handler
	AggregateTradeHandler func(*models.AggregateTradeEvent) error
	
	// MarkPriceEvent Handler
	MarkPriceHandler func(*models.MarkPriceEvent) error
	
	// KlineEvent Handler
	KlineHandler func(*models.KlineEvent) error
	
	// ContinuousKlineEvent Handler
	ContinuousKlineHandler func(*models.ContinuousKlineEvent) error
	
	// MiniTickerEvent Handler
	MiniTickerHandler func(*models.MiniTickerEvent) error
	
	// TickerEvent Handler
	TickerHandler func(*models.TickerEvent) error
	
	// BookTickerEvent Handler
	BookTickerHandler func(*models.BookTickerEvent) error
	
	// LiquidationEvent Handler
	LiquidationHandler func(*models.LiquidationEvent) error
	
	// DiffDepthEvent Handler
	DiffDepthHandler func(*models.DiffDepthEvent) error
	
	// CompositeIndexEvent Handler
	CompositeIndexHandler func(*models.CompositeIndexEvent) error
	
	// ContractInfoEvent Handler
	ContractInfoHandler func(*models.ContractInfoEvent) error
	
	// AssetIndexEvent Handler
	AssetIndexHandler func(*models.AssetIndexEvent) error
	
	// Combined Stream Handler
	CombinedStreamHandler func(*models.CombinedStreamEvent) error
	
	// Subscription Response Handler
	SubscriptionResponseHandler func(*models.SubscriptionResponse) error
	
	// Error Handler
	StreamErrorHandler func(*models.ErrorResponse) error
)

// Event handler registry
type eventHandlers struct {
	aggregateTrade AggregateTradeHandler
	markPriceUpdate MarkPriceHandler
	kline KlineHandler
	continuouskline ContinuousKlineHandler
	miniTicker MiniTickerHandler
	ticker TickerHandler
	bookTicker BookTickerHandler
	forceOrder LiquidationHandler
	depth DiffDepthHandler
	compositeIndex CompositeIndexHandler
	contractInfo ContractInfoHandler
	assetIndexUpdate AssetIndexHandler
	assetIndex AssetIndexHandler
	combinedStream     CombinedStreamHandler
	subscriptionResponse SubscriptionResponseHandler
	error              StreamErrorHandler
}

// Register event handlers
func (c *Client) OnAggregateTradeEvent(handler AggregateTradeHandler) {
	c.handlers.aggregateTrade = handler
}

func (c *Client) OnMarkPriceEvent(handler MarkPriceHandler) {
	c.handlers.markPriceUpdate = handler
}

func (c *Client) OnKlineEvent(handler KlineHandler) {
	c.handlers.kline = handler
}

func (c *Client) OnContinuousKlineEvent(handler ContinuousKlineHandler) {
	c.handlers.continuouskline = handler
}

func (c *Client) OnMiniTickerEvent(handler MiniTickerHandler) {
	c.handlers.miniTicker = handler
}

func (c *Client) OnTickerEvent(handler TickerHandler) {
	c.handlers.ticker = handler
}

func (c *Client) OnBookTickerEvent(handler BookTickerHandler) {
	c.handlers.bookTicker = handler
}

func (c *Client) OnLiquidationEvent(handler LiquidationHandler) {
	c.handlers.forceOrder = handler
}

func (c *Client) OnDiffDepthEvent(handler DiffDepthHandler) {
	c.handlers.depth = handler
}

func (c *Client) OnCompositeIndexEvent(handler CompositeIndexHandler) {
	c.handlers.compositeIndex = handler
}

func (c *Client) OnContractInfoEvent(handler ContractInfoHandler) {
	c.handlers.contractInfo = handler
}

func (c *Client) OnAssetIndexEvent(handler AssetIndexHandler) {
	c.handlers.assetIndexUpdate = handler
	c.handlers.assetIndex = handler
}

func (c *Client) OnCombinedStreamEvent(handler CombinedStreamHandler) {
	c.handlers.combinedStream = handler
}

func (c *Client) OnSubscriptionResponse(handler SubscriptionResponseHandler) {
	c.handlers.subscriptionResponse = handler
}

func (c *Client) OnStreamError(handler StreamErrorHandler) {
	c.handlers.error = handler
}

// processStreamMessage processes incoming stream messages
func (c *Client) processStreamMessage(data []byte) error {
	// First check if this is an array stream (like !assetIndex@arr)
	// by trying to parse as an array first
	var arrayData []interface{}
	if err := json.Unmarshal(data, &arrayData); err == nil {
		// This is an array stream - process as array of events
		return c.processArrayStreamEvent(data, arrayData)
	}
	
	// Parse message as object to determine type
	var baseMsg map[string]interface{}
	if err := json.Unmarshal(data, &baseMsg); err != nil {
		return fmt.Errorf("failed to parse message: %w", err)
	}

	// Check for subscription response
	if _, hasID := baseMsg["id"]; hasID {
		var response models.SubscriptionResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return fmt.Errorf("failed to parse subscription response: %w", err)
		}
		if c.handlers.subscriptionResponse != nil {
			return c.handlers.subscriptionResponse(&response)
		}
		return nil
	}

	// Check for error response
	if errorData, hasError := baseMsg["error"]; hasError && errorData != nil {
		var errorResp models.ErrorResponse
		if err := json.Unmarshal(data, &errorResp); err != nil {
			return fmt.Errorf("failed to parse error response: %w", err)
		}
		if c.handlers.error != nil {
			return c.handlers.error(&errorResp)
		}
		return nil
	}

	// Check for combined stream format
	if _, hasStream := baseMsg["stream"]; hasStream {
		var combined models.CombinedStreamEvent
		if err := json.Unmarshal(data, &combined); err != nil {
			return fmt.Errorf("failed to parse combined stream: %w", err)
		}
		if c.handlers.combinedStream != nil {
			return c.handlers.combinedStream(&combined)
		}
		// Also try to process the nested data
		if dataBytes, err := json.Marshal(combined.StreamData); err == nil {
			return c.processSingleStreamEvent(dataBytes)
		}
		return nil
	}

	// Process as single stream event
	return c.processSingleStreamEvent(data)
}

// processArrayStreamEvent processes array stream events (like !assetIndex@arr)
func (c *Client) processArrayStreamEvent(data []byte, arrayData []interface{}) error {
	// Array streams contain multiple events of the same type
	// Process each element in the array individually
	if len(arrayData) == 0 {
		return nil // Empty array, nothing to process
	}
	
	// Process each element in the array
	for i, element := range arrayData {
		elementBytes, err := json.Marshal(element)
		if err != nil {
			log.Printf("Failed to marshal array element %d: %v", i, err)
			continue
		}
		
		if err := c.processSingleStreamEvent(elementBytes); err != nil {
			log.Printf("Failed to process array element %d: %v", i, err)
			// Continue processing other elements even if one fails
		}
	}
	
	return nil
}

// processSingleStreamEvent processes individual stream events
func (c *Client) processSingleStreamEvent(data []byte) error {
	// First try to parse the message as a generic map to handle flexible event type field
	var genericMsg map[string]interface{}
	if err := json.Unmarshal(data, &genericMsg); err != nil {
		return fmt.Errorf("failed to parse message: %w", err)
	}
	
	// Extract event type, handling both string and number formats
	var eventType string
	if eValue, hasEventType := genericMsg["e"]; hasEventType {
		switch v := eValue.(type) {
		case string:
			eventType = v
		case float64:
			// Handle numeric event types by converting to string
			eventType = fmt.Sprintf("%.0f", v)
		case int:
			// Handle integer event types by converting to string
			eventType = fmt.Sprintf("%d", v)
		default:
			return fmt.Errorf("unknown event type format: %T", v)
		}
	} else {
		return fmt.Errorf("no event type field found")
	}

	switch eventType {
	case "aggTrade":
		if c.handlers.aggregateTrade != nil {
			var event models.AggregateTradeEvent
			if err := json.Unmarshal(data, &event); err != nil {
				return fmt.Errorf("failed to parse aggTrade event: %w", err)
			}
			return c.handlers.aggregateTrade(&event)
		}
	case "markPriceUpdate":
		if c.handlers.markPriceUpdate != nil {
			var event models.MarkPriceEvent
			if err := json.Unmarshal(data, &event); err != nil {
				return fmt.Errorf("failed to parse markPriceUpdate event: %w", err)
			}
			return c.handlers.markPriceUpdate(&event)
		}
	case "kline":
		if c.handlers.kline != nil {
			var event models.KlineEvent
			if err := json.Unmarshal(data, &event); err != nil {
				return fmt.Errorf("failed to parse kline event: %w", err)
			}
			return c.handlers.kline(&event)
		}
	case "continuous_kline":
		if c.handlers.continuouskline != nil {
			var event models.ContinuousKlineEvent
			if err := json.Unmarshal(data, &event); err != nil {
				return fmt.Errorf("failed to parse continuous_kline event: %w", err)
			}
			return c.handlers.continuouskline(&event)
		}
	case "24hrMiniTicker":
		if c.handlers.miniTicker != nil {
			var event models.MiniTickerEvent
			if err := json.Unmarshal(data, &event); err != nil {
				return fmt.Errorf("failed to parse 24hrMiniTicker event: %w", err)
			}
			return c.handlers.miniTicker(&event)
		}
	case "24hrTicker":
		if c.handlers.ticker != nil {
			var event models.TickerEvent
			if err := json.Unmarshal(data, &event); err != nil {
				return fmt.Errorf("failed to parse 24hrTicker event: %w", err)
			}
			return c.handlers.ticker(&event)
		}
	case "bookTicker":
		if c.handlers.bookTicker != nil {
			var event models.BookTickerEvent
			if err := json.Unmarshal(data, &event); err != nil {
				return fmt.Errorf("failed to parse bookTicker event: %w", err)
			}
			return c.handlers.bookTicker(&event)
		}
	case "forceOrder":
		if c.handlers.forceOrder != nil {
			var event models.LiquidationEvent
			if err := json.Unmarshal(data, &event); err != nil {
				return fmt.Errorf("failed to parse forceOrder event: %w", err)
			}
			return c.handlers.forceOrder(&event)
		}
	case "depthUpdate":
		if c.handlers.depth != nil {
			var event models.DiffDepthEvent
			if err := json.Unmarshal(data, &event); err != nil {
				return fmt.Errorf("failed to parse depthUpdate event: %w", err)
			}
			return c.handlers.depth(&event)
		}
	case "compositeIndex":
		if c.handlers.compositeIndex != nil {
			var event models.CompositeIndexEvent
			if err := json.Unmarshal(data, &event); err != nil {
				return fmt.Errorf("failed to parse compositeIndex event: %w", err)
			}
			return c.handlers.compositeIndex(&event)
		}
	case "contractInfo":
		if c.handlers.contractInfo != nil {
			var event models.ContractInfoEvent
			if err := json.Unmarshal(data, &event); err != nil {
				return fmt.Errorf("failed to parse contractInfo event: %w", err)
			}
			return c.handlers.contractInfo(&event)
		}
	case "assetIndexUpdate":
		if c.handlers.assetIndexUpdate != nil {
			var event models.AssetIndexEvent
			if err := json.Unmarshal(data, &event); err != nil {
				return fmt.Errorf("failed to parse assetIndexUpdate event: %w", err)
			}
			return c.handlers.assetIndexUpdate(&event)
		}
	case "assetIndex":
		if c.handlers.assetIndex != nil {
			var event models.AssetIndexEvent
			if err := json.Unmarshal(data, &event); err != nil {
				return fmt.Errorf("failed to parse assetIndex event: %w", err)
			}
			return c.handlers.assetIndex(&event)
		}
	default:
		// Log unknown event types with full message for debugging
		log.Printf("Unknown event type '%s' in umfutures-streams. Message: %s", eventType, string(data))
		// Don't return error for unknown event types, just ignore them
	}
	return nil
}


// Additional stream utility functions



