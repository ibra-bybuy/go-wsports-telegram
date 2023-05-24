package controller

import (
	"context"

	"github.com/ibra-bybuy/go-wsports-telegram/pkg/model"
)

type repository interface {
	GetBySport(ctx context.Context, sport string, limit int, page int) *model.PaginationEvents
	GetByUuid(ctx context.Context, uuid string) *model.SuccessEvent
}

type Controller struct {
	r repository
}

func New(r repository) *Controller {
	return &Controller{r}
}

func (c *Controller) GetBySport(ctx context.Context, sport string, limit int, page int) *model.PaginationEvents {
	return c.r.GetBySport(ctx, sport, limit, page)
}

func (c *Controller) GetByUuid(ctx context.Context, uuid string) *model.SuccessEvent {
	return c.r.GetByUuid(ctx, uuid)
}
