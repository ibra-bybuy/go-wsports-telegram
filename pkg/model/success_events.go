package model

type SuccessEventsResponse struct {
	Success bool             `json:"success"`
	Data    PaginationEvents `json:"data,omitempty"`
}
