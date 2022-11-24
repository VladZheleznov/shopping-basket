package repository

import (
	"context"
	"fmt"

	"github.com/VladZheleznov/shopping-basket/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// PRepository p
type PRepository struct {
	Pool *pgxpool.Pool
}

// AddItem add item to db
func (p *PRepository) AddItem(ctx context.Context, product *model.Product) (string, error) {
	newID := uuid.New().String()
	_, err := p.Pool.Exec(ctx, "insert into products(id,name,price,quantity) values($1,$2,$3,$4)",
		newID, &product.Name, &product.Price, &product.Quantity)
	if err != nil {
		panic(err)
	}
	return newID, nil
}

// GetItemByID select item by id
func (p *PRepository) GetItemByID(ctx context.Context, idProduct string) (*model.Product, error) {
	u := model.Product{}
	err := p.Pool.QueryRow(ctx, "select id,name,price,quantity from products where id=$1", idProduct).Scan(
		&u.ID, &u.Name, &u.Price, &u.Quantity)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &model.Product{}, fmt.Errorf("user with this id doesnt exist: %v", err)
		}
		panic(err)
	}
	return &u, nil
}

// GetAllItems select all items from db
func (p *PRepository) GetAllItems(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product
	rows, err := p.Pool.Query(ctx, "select id,name,price,quantity from products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		pro := model.Product{}
		err = rows.Scan(&pro.ID, &pro.Name, &pro.Price, &pro.Quantity)
		if err != nil {
			panic(err)
		}
		products = append(products, &pro)
	}

	return products, nil
}

// DeleteItem delete item by id
func (p *PRepository) DeleteItem(ctx context.Context, id string) error {
	a, err := p.Pool.Exec(ctx, "delete from products where id=$1", id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("user with this id doesnt exist")
	}
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("user with this id doesnt exist: %v", err)
		}
		panic(err)
	}
	return nil
}

// UpdateItem update parameters for item
func (p *PRepository) UpdateItem(ctx context.Context, id string, pro *model.Product) error {
	a, err := p.Pool.Exec(ctx, "update products set name=$1,price=$2,quantity=$3 where id=$4", &pro.Name, &pro.Price, &pro.Quantity, id)
	if a.RowsAffected() == 0 {
		return fmt.Errorf("user with this id doesnt exist")
	}
	if err != nil {
		panic(err)
	}
	return nil
}
