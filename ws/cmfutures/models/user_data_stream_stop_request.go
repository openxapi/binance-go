package models

import (
	"encoding/json"
)

// UserDataStreamStopRequestParams contains the parameters for UserDataStreamStopRequest
type UserDataStreamStopRequestParams struct {
	// Required if session is not authenticated via session.logon
	ApiKey *string `json:"apiKey,omitempty"`
}

// UserDataStreamStopRequest - Send a userDataStream.stop request
// Message name: Close User Data Stream (USER_STREAM) Request
type UserDataStreamStopRequest struct {
	// id property
	Id string `json:"id,omitempty"`
	// method property
	Method string `json:"method,omitempty"`
	// params property
	Params *UserDataStreamStopRequestParams `json:"params,omitempty"`
}

// NewUserDataStreamStopRequest creates a new UserDataStreamStopRequest instance
func NewUserDataStreamStopRequest() *UserDataStreamStopRequest {
	return &UserDataStreamStopRequest{}
}

// String returns string representation of UserDataStreamStopRequest
func (s UserDataStreamStopRequest) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// SetParams sets the params field and returns the struct for method chaining
func (r *UserDataStreamStopRequest) SetParams(value UserDataStreamStopRequestParams) *UserDataStreamStopRequest {
	r.Params = &value
	return r
}

// SetApiKey sets the apiKey parameter and returns the struct for method chaining
func (r *UserDataStreamStopRequest) SetApiKey(value string) *UserDataStreamStopRequest {
	if r.Params == nil {
		r.Params = &UserDataStreamStopRequestParams{}
	}
	r.Params.ApiKey = &value
	return r
}


