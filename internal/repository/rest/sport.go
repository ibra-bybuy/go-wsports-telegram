package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ibra-bybuy/go-wsports-telegram/pkg/model"
)

func (r *repository) GetBySport(ctx context.Context, sport string, limit int, page int) *model.PaginationEvents {
	result := model.SuccessEventsResponse{}
	url := r.baseURL + fmt.Sprintf("/events?limit=%d&page=%d&sport=%s", limit, page, sport)
	response, err := http.Get(url)
	if err != nil {
		return &result.Data
	}
	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(&result)

	return &result.Data
}
