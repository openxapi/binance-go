package models

import (
	"encoding/json"
)

// UserDataStreamPingRequestParams contains the parameters for UserDataStreamPingRequest
type UserDataStreamPingRequestParams struct {
	ApiKey string `json:"apiKey"`
	ListenKey string `json:"listenKey"`
}

// UserDataStreamPingRequest - Send a userDataStream.ping request
// Message name: Ping user data stream (USER_STREAM) (Deprecated) Request
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
	r.Params.ApiKey = value
	return r
}

// SetListenKey sets the listenKey parameter and returns the struct for method chaining
func (r *UserDataStreamPingRequest) SetListenKey(value string) *UserDataStreamPingRequest {
	if r.Params == nil {
		r.Params = &UserDataStreamPingRequestParams{}
	}
	r.Params.ListenKey = value
	return r
}


