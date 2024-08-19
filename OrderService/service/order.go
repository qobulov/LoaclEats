package service

import (
	pb "Orders/genproto/orders"
	"context"
)

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	order, err := s.Db.CreateOrder(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return order, nil
}

func (s *OrderService) GetOrderByID(ctx context.Context, req *pb.GetOrderByIDRequest) (*pb.GetOrderByIDResponse, error) {
	order, err := s.Db.GetOrderByID(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return order, nil
}

func (s *OrderService) GetUserOrders(ctx context.Context, req *pb.GetUserOrdersRequest) (*pb.GetUserOrdersResponse, error) {
	orders, err := s.Db.GetUserOrders(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return orders, nil
}

func (s *OrderService) GetKitchenOrders(ctx context.Context, req *pb.GetKitchenOrdersRequest) (*pb.GetKitchenOrdersResponse, error) {
	orders, err := s.Db.GetKitchenOrders(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return orders, nil
}

func (s *OrderService) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest) (*pb.Status, error) {
	order, err := s.Db.UpdateOrderStatus(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return order, nil
}