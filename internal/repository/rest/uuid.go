package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ibra-bybuy/go-wsports-telegram/pkg/model"
)

func (r *repository) GetByUuid(ctx context.Context, uuid string) *model.SuccessEvent {
	result := model.SuccessEvent{}
	v := url.Values{}
	v.Add("uuid", uuid)
	url := r.baseURL + fmt.Sprintf("/events?%s", v.Encode())
	response, err := http.Get(url)
	if err != nil {
		return &result
	}
	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(&result)

	return &result
}
