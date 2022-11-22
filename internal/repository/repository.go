package repository

import (
	"github.com/VladZheleznov/shopping-basket/internal/model"

	"context"
)

type Repository interface {
	AddItem(ctx context.Context, p *model.Product) (string, error)
	GetItemByID(ctx context.Context, idProduct string) (*model.Product, error)
	GetAllItems(ctx context.Context) ([]*model.Product, error)
	DeleteItem(ctx context.Context, id string) error
	UpdateItem(ctx context.Context, id string, pro *model.Product) error
}
