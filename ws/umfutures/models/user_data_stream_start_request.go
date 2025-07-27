package models

import (
	"encoding/json"
)

// UserDataStreamStartRequestParams contains the parameters for UserDataStreamStartRequest
type UserDataStreamStartRequestParams struct {
	// apiKey property
	ApiKey *string `json:"apiKey,omitempty"`
}

// UserDataStreamStartRequest - Send a userDataStream.start request
// Message name: Start User Data Stream (USER_STREAM) Request
type UserDataStreamStartRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *UserDataStreamStartRequestParams `json:"params,omitempty"`
}

// NewUserDataStreamStartRequest creates a new UserDataStreamStartRequest instance
func NewUserDataStreamStartRequest() *UserDataStreamStartRequest {
	return &UserDataStreamStartRequest{}
}

// String returns string representation of UserDataStreamStartRequest
func (s UserDataStreamStartRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *UserDataStreamStartRequest) SetParams(value UserDataStreamStartRequestParams) *UserDataStreamStartRequest {
	r.Params = &value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *UserDataStreamStartRequest) SetApiKey(value string) *UserDataStreamStartRequest {
	if r.Params == nil {
		r.Params = &UserDataStreamStartRequestParams{}
	}
	r.Params.ApiKey = &value
	return r
}


