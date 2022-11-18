package repository

import (
	"context"

	"shopping-basket/internal/model"
)

type Repository interface {
	CreateUser(ctx context.Context, p *model.Product) (string, error)
	GetUserByID(ctx context.Context, idProduct string) (*model.Product, error)
	GetAllUsers(ctx context.Context) ([]*model.Product, error)
	DeleteUser(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, id string, pro *model.Product) error
}
