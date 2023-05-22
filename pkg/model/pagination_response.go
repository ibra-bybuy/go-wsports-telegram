package model

type PaginationEvents struct {
	Pagination Pagination `json:"pagination,omitempty"`
	Items      []Event    `json:"items"`
}
