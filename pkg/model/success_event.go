package model

type SuccessEvent struct {
	Success bool  `json:"success"`
	Data    Event `json:"data,omitempty"`
}
