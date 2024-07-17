package service

import (
	postgres "Orders/Storage"
	pb "Orders/genproto/orders"
	"Orders/pkg/logger"
	"context"
	"database/sql"
	"log/slog"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	Logger *slog.Logger
	Db     *postgres.OrderService
}

func NewOrderService(db *sql.DB) *OrderService {
	return &OrderService{
		Logger: logger.NewLogger(),
		Db:     postgres.NewOrdersRepo(db),
	}
}

func (s *OrderService) GetDishes(ctx context.Context, req *pb.GetDishesRequest) (*pb.GetDishesResponse, error) {
	dish, err := s.Db.GetDishes(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return dish, nil
}

func (s *OrderService) CreateDish(ctx context.Context, req *pb.CreateDishRequest) (*pb.Dish, error) {
	dish, err := s.Db.CreateDish(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return dish, nil
}

func (s *OrderService) UpdateDish(ctx context.Context, req *pb.UpdateDishRequest) (*pb.Dish, error) {
	dish, err := s.Db.UpdateDish(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return dish, nil
}

func (s *OrderService) DeleteDish(ctx context.Context, req *pb.DeleteDishRequest) (*pb.Status, error) {
	status, err := s.Db.DeleteDish(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return status, nil
}

func (s *OrderService) GetDish(ctx context.Context, req *pb.GetDishRequest) (*pb.Dish, error) {
	dish, err := s.Db.GetDish(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return dish, nil
}