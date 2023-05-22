package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ibra-bybuy/go-wsports-telegram/pkg/model"
)

func (r *repository) GetByID(ctx context.Context, id string) *model.SuccessEvent {
	result := model.SuccessEvent{
		Data: model.Event{
			Streams: []model.Stream{},
		},
	}
	url := r.baseURL + fmt.Sprintf("/events?id=%s", id)
	response, err := http.Get(url)
	if err != nil {
		return &result
	}
	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(&result)

	return &result
}
