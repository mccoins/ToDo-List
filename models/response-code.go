package models

// ResponseCode is used to describe the details of a response
type ResponseCode struct {
	ID          string `json:"Id,omitempty"`
	Description string `json:"Description,omitempty"`
}
