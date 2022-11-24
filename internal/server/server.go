package server

import (
	"github.com/VladZheleznov/shopping-basket/internal/model"
	"github.com/VladZheleznov/shopping-basket/internal/service"
	pb "github.com/VladZheleznov/shopping-basket/proto"

	"context"
)

type Server struct {
	pb.UnimplementedCRUDServer
	se *service.Service
}

// NewServer create new server connection
func NewServer(serv *service.Service) *Server {
	return &Server{se: serv}
}

func (s *Server) Registration(ctx context.Context, request *pb.RegistrationRequest) (*pb.RegistrationResponse, error) {
	p := model.Product{
		Name:     request.Name,
		Price:    request.Price,
		Quantity: request.Quantity,
	}
	newID, err := s.se.AddItem(ctx, &p)
	if err != nil {
		return nil, err
	}
	return &pb.RegistrationResponse{Id: newID}, nil
}

// GetItem get item by id from db
func (s *Server) GetItem(ctx context.Context, request *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	idProduct := request.GetId()
	productDB, err := s.se.GetItem(ctx, idProduct)
	if err != nil {
		return nil, err
	}
	productProto := &pb.GetItemResponse{
		Product: &pb.Product{
			Id:       productDB.ID,
			Name:     productDB.Name,
			Price:    productDB.Price,
			Quantity: productDB.Quantity,
		},
	}
	return productProto, nil
}

// GetAllItems get all items from db
func (s *Server) GetAllItems(ctx context.Context, _ *pb.GetAllItemsRequest) (*pb.GetAllItemsResponse, error) {
	products, err := s.se.GetAllItems(ctx)
	if err != nil {
		return nil, err
	}
	var list []*pb.Product
	for _, product := range products {
		productProto := new(pb.Product)
		productProto.Id = product.ID
		productProto.Name = product.Name
		productProto.Price = product.Price
		productProto.Quantity = product.Quantity
		list = append(list, productProto)
	}
	return &pb.GetAllItemsResponse{Product: list}, nil
}

// DeleteItem delete item by id
func (s *Server) DeleteItem(ctx context.Context, request *pb.DeleteItemRequest) (*pb.Response, error) {
	idItem := request.GetId()
	err := s.se.DeleteItem(ctx, idItem)
	if err != nil {
		return nil, err
	}
	return new(pb.Response), nil
}

// UpdateItem update item with new parameters
func (s *Server) UpdateItem(ctx context.Context, request *pb.UpdateItemRequest) (*pb.Response, error) {
	item := &model.Product{
		Name:     request.Product.Name,
		Quantity: request.Product.Quantity,
		Price:    request.Product.Price,
	}
	idItem := request.GetId()
	err := s.se.UpdateItem(ctx, idItem, item)
	if err != nil {
		return nil, err
	}
	return new(pb.Response), nil
}
