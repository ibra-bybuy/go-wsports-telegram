package model

type SuccessEvent struct {
	Success bool  `json:"success"`
	Data    Event `json:"event,omitempty"`
}
