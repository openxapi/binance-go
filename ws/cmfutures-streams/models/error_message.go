package models

// ErrorMessage represents global message '#/components/messages/errorMessage'
type ErrorMessage struct {
	Error struct {
		Code int `json:"code,omitempty"` // Error code
		Msg string `json:"msg,omitempty"` // Error message
	} `json:"error,omitempty"`
	Id int64 `json:"id,omitempty"`
}


