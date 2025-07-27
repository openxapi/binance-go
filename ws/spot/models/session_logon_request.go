package models

import (
	"encoding/json"
)

// SessionLogonRequestParams contains the parameters for SessionLogonRequest
type SessionLogonRequestParams struct {
	ApiKey string `json:"apiKey"`
	// The value cannot be greater than 60000
	RecvWindow *int64 `json:"recvWindow,omitempty"`
	Signature string `json:"signature"`
	Timestamp int64 `json:"timestamp"`
}

// SessionLogonRequest - Send a session.logon request
// Message name: Log in with API key (SIGNED) Request
type SessionLogonRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *SessionLogonRequestParams `json:"params,omitempty"`
}

// NewSessionLogonRequest creates a new SessionLogonRequest instance
func NewSessionLogonRequest() *SessionLogonRequest {
	return &SessionLogonRequest{}
}

// String returns string representation of SessionLogonRequest
func (s SessionLogonRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *SessionLogonRequest) SetParams(value SessionLogonRequestParams) *SessionLogonRequest {
	r.Params = &value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *SessionLogonRequest) SetApiKey(value string) *SessionLogonRequest {
	if r.Params == nil {
		r.Params = &SessionLogonRequestParams{}
	}
	r.Params.ApiKey = value
	return r
}

// SetRecvWindow sets the recvWindow parameter and returns the struct for method chaining
func (r *SessionLogonRequest) SetRecvWindow(value int64) *SessionLogonRequest {
	if r.Params == nil {
		r.Params = &SessionLogonRequestParams{}
	}
	r.Params.RecvWindow = &value
	return r
}

// SetSignature sets the signature parameter and returns the struct for method chaining
func (r *SessionLogonRequest) SetSignature(value string) *SessionLogonRequest {
	if r.Params == nil {
		r.Params = &SessionLogonRequestParams{}
	}
	r.Params.Signature = value
	return r
}

// SetTimestamp sets the timestamp parameter and returns the struct for method chaining
func (r *SessionLogonRequest) SetTimestamp(value int64) *SessionLogonRequest {
	if r.Params == nil {
		r.Params = &SessionLogonRequestParams{}
	}
	r.Params.Timestamp = value
	return r
}


