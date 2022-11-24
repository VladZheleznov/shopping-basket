package service

import (
	"github.com/VladZheleznov/shopping-basket/internal/model"
	"github.com/VladZheleznov/shopping-basket/internal/repository"

	"context"
)

type Service struct {
	rps repository.Repository
}

// NewService create new service connection
func NewService(pool repository.Repository) *Service {
	return &Service{rps: pool}
}

// AddItem
func (se *Service) AddItem(ctx context.Context, p *model.Product) (string, error) {
	return se.rps.AddItem(ctx, p)
}

// GetItem _
func (se *Service) GetItem(ctx context.Context, id string) (*model.Product, error) {
	return se.rps.GetItemByID(ctx, id)
}

// GetAllItems _
func (se *Service) GetAllItems(ctx context.Context) ([]*model.Product, error) {
	return se.rps.GetAllItems(ctx)
}

// DeleteItem _
func (se *Service) DeleteItem(ctx context.Context, id string) error {
	return se.rps.DeleteItem(ctx, id)
}

// UpdateItem _
func (se *Service) UpdateItem(ctx context.Context, id string, item *model.Product) error {
	return se.rps.UpdateItem(ctx, id, item)
}
