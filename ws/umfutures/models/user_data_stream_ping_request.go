package models

import (
	"encoding/json"
)

// UserDataStreamPingRequestParams contains the parameters for UserDataStreamPingRequest
type UserDataStreamPingRequestParams struct {
	// apiKey property
	ApiKey *string `json:"apiKey,omitempty"`
}

// UserDataStreamPingRequest - Send a userDataStream.ping request
// Message name: Keepalive User Data Stream (USER_STREAM) Request
type UserDataStreamPingRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *UserDataStreamPingRequestParams `json:"params,omitempty"`
}

// NewUserDataStreamPingRequest creates a new UserDataStreamPingRequest instance
func NewUserDataStreamPingRequest() *UserDataStreamPingRequest {
	return &UserDataStreamPingRequest{}
}

// String returns string representation of UserDataStreamPingRequest
func (s UserDataStreamPingRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *UserDataStreamPingRequest) SetParams(value UserDataStreamPingRequestParams) *UserDataStreamPingRequest {
	r.Params = &value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *UserDataStreamPingRequest) SetApiKey(value string) *UserDataStreamPingRequest {
	if r.Params == nil {
		r.Params = &UserDataStreamPingRequestParams{}
	}
	r.Params.ApiKey = &value
	return r
}


