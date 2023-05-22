package rest

type repository struct {
	baseURL string
}

func New(baseURL string) *repository {
	return &repository{baseURL}
}
