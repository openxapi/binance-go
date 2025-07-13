package models
import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
)

// MessageType represents the type of a WebSocket message
type MessageType string

// OneOfType represents a type that can hold one of multiple possible types
type OneOfType interface {
	GetType() string
	GetValue() interface{}
}

// ResponseRegistry manages all possible response types
type ResponseRegistry struct {
	types   map[string]reflect.Type
	creators map[string]func() interface{}
	mu      sync.RWMutex
}

// NewResponseRegistry creates a new response registry
func NewResponseRegistry() *ResponseRegistry {
	return &ResponseRegistry{
		types:   make(map[string]reflect.Type),
		creators: make(map[string]func() interface{}),
	}
}

// RegisterType registers a response type for dynamic parsing
func (r *ResponseRegistry) RegisterType(typeName string, typeInstance interface{}, creator func() interface{}) {
	r.mu.Lock()
	r.types[typeName] = reflect.TypeOf(typeInstance)
	r.creators[typeName] = creator
	r.mu.Unlock()
}

// CreateInstance creates a new instance of the registered type
func (r *ResponseRegistry) CreateInstance(typeName string) (interface{}, error) {
	r.mu.RLock()
	creator, exists := r.creators[typeName]
	r.mu.RUnlock()
	
	if !exists {
		return nil, fmt.Errorf("unknown type: %s", typeName)
	}
	
	return creator(), nil
}

// GetType returns the reflect.Type for a registered type name
func (r *ResponseRegistry) GetType(typeName string) (reflect.Type, bool) {
	r.mu.RLock()
	typ, exists := r.types[typeName]
	r.mu.RUnlock()
	return typ, exists
}

// ListTypes returns all registered type names
func (r *ResponseRegistry) ListTypes() []string {
	r.mu.RLock()
	names := make([]string, 0, len(r.types))
	for name := range r.types {
		names = append(names, name)
	}
	r.mu.RUnlock()
	return names
}

// Global response registry instance
var GlobalRegistry = NewResponseRegistry()

// Common message parsing utilities
var (
	// MessageTypeMap maps message IDs to their corresponding Go types
	MessageTypeMap = make(map[string]reflect.Type)
)

// RegisterMessageType registers a message type for dynamic parsing
func RegisterMessageType(messageID string, messageType interface{}) {
	MessageTypeMap[messageID] = reflect.TypeOf(messageType)
}

// ParseMessage parses a JSON message into the specified struct type
func ParseMessage[T any](data []byte, target *T) error {
	return json.Unmarshal(data, target)
}

// ParseDynamicMessage parses a message based on its ID
func ParseDynamicMessage(messageID string, data []byte) (interface{}, error) {
	msgType, exists := MessageTypeMap[messageID]
	if !exists {
		return nil, fmt.Errorf("unknown message type: %s", messageID)
	}
	
	// Create a new instance of the message type
	msg := reflect.New(msgType).Interface()
	err := json.Unmarshal(data, msg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse message: %w", err)
	}
	
	return msg, nil
}

// ParseOneOfResult attempts to parse a oneOf result based on type detection
func ParseOneOfResult(data []byte) (interface{}, string, error) {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, "", err
	}
	
	// Check for event type field (common in Binance WebSocket events)
	if eventType, exists := raw["e"]; exists {
		if eventTypeStr, ok := eventType.(string); ok {
			// Map event types to struct types
			responseType := mapEventTypeToStruct(eventTypeStr)
			if responseType != "" {
				instance, err := GlobalRegistry.CreateInstance(responseType)
				if err != nil {
					return nil, "", err
				}
				
				if err := json.Unmarshal(data, instance); err != nil {
					return nil, "", err
				}
				
				return instance, responseType, nil
			}
		}
	}
	
	return nil, "", fmt.Errorf("unable to determine oneOf type")
}

// mapEventTypeToStruct maps Binance event types to Go struct types
func mapEventTypeToStruct(eventType string) string {
	// This function will be populated based on actual event types in the spec
	// For APIs that don't define event types (like umfutures), this returns empty string
	return ""
}

// RegisterAllEventTypes registers all known event types with the global registry
func RegisterAllEventTypes() {
	// Event types will be registered here based on what's actually defined in the AsyncAPI spec
	// For APIs that don't define event types (like umfutures), this is a no-op function
}

// MessageValidator interface for messages that can validate themselves
type MessageValidator interface {
	Validate() error
}

// ValidateMessage validates a message if it implements MessageValidator
func ValidateMessage(msg interface{}) error {
	if validator, ok := msg.(MessageValidator); ok {
		return validator.Validate()
	}
	return nil
}

// ResponseTypeDetector interface for responses that can identify their own type
type ResponseTypeDetector interface {
	DetectType() string
}

// DetectResponseType attempts to detect the type of a response
func DetectResponseType(msg interface{}) string {
	if detector, ok := msg.(ResponseTypeDetector); ok {
		return detector.DetectType()
	}
	return fmt.Sprintf("%T", msg)
}

// init function to register all event types on package load
func init() {
	// This will be called when the package is loaded
	// Individual model files will call their registration functions
}

