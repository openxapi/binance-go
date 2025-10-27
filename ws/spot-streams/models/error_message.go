package models

// ErrorMessage represents global message '#/components/messages/errorMessage'
type ErrorMessage struct {
	Error struct {
		Code int `json:"code,omitempty"` // Error code
		Msg string `json:"msg,omitempty"` // Error message
	} `json:"error,omitempty"`
	Id MessageID `json:"id,omitempty"` // The id is used as an identifier to uniquely identify the messages going back and forth. Accepted formats: - 64-bit signed integer - Alphanumeric strings; max length 36 - null 
}


