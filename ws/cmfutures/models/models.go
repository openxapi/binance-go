package models
import (
	"bytes"
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
	// This function is populated from event types declared in the AsyncAPI spec.
	// When the spec omits explicit event identifiers, return empty string.
	return ""
}

// RegisterAllEventTypes registers all known event types with the global registry
func RegisterAllEventTypes() {
	// Event types are registered here when provided by the specification.
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

// MessageID represents a request/response id that can be either:
// - 64-bit signed integer
// - alphanumeric string (max length 36)
// - null
type MessageID struct {
	i64   *int64
	str   *string
	isNull bool
}

// NewMessageIDInt64 creates a MessageID from int64
func NewMessageIDInt64(v int64) MessageID { return MessageID{i64: &v} }

// NewMessageIDString creates a MessageID from string (<=36 chars)
func NewMessageIDString(v string) MessageID { return MessageID{str: &v} }

// NewMessageIDNull creates a null MessageID
func NewMessageIDNull() MessageID { return MessageID{isNull: true} }

// String returns a canonical string form used for correlation maps
func (m MessageID) String() string {
	if m.isNull { return "" }
	if m.str != nil { return *m.str }
	if m.i64 != nil { return fmt.Sprintf("%d", *m.i64) }
	return ""
}

// MarshalJSON encodes MessageID as number, string, or null
func (m MessageID) MarshalJSON() ([]byte, error) {
	if m.isNull { return []byte("null"), nil }
	if m.str != nil { return json.Marshal(*m.str) }
	if m.i64 != nil { return json.Marshal(*m.i64) }
	return []byte("null"), nil
}

// UnmarshalJSON decodes MessageID from number, string, or null
func (m *MessageID) UnmarshalJSON(b []byte) error {
	// reset
	*m = MessageID{}
	// Handle null
	if len(b) == 0 || string(b) == "null" { m.isNull = true; return nil }
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	var v interface{}
	if err := dec.Decode(&v); err != nil { return err }
	switch t := v.(type) {
	case json.Number:
		i, err := t.Int64()
		if err != nil { return fmt.Errorf("invalid id number: %w", err) }
		m.i64 = &i
		return nil
	case string:
		if len(t) > 36 { return fmt.Errorf("id string too long: %d", len(t)) }
		m.str = &t
		return nil
	default:
		return fmt.Errorf("invalid id type: %T", v)
	}
}

// ValInt64 returns the int64 value and true if MessageID holds an integer; otherwise returns 0, false
func (m MessageID) ValInt64() (int64, bool) {
    if m.i64 != nil {
        return *m.i64, true
    }
    return 0, false
}

// ValString returns the string value and true if MessageID holds a string; otherwise returns "", false
func (m MessageID) ValString() (string, bool) {
    if m.str != nil {
        return *m.str, true
    }
    return "", false
}

// ValNull reports whether MessageID is explicitly null
func (m MessageID) ValNull() bool { return m.isNull }


