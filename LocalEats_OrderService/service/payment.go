package service

import (
	pb "Orders/genproto/orders"
	"context"
)

func (s *OrderService) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.Payment, error) {
	payment, err := s.Db.CreatePayment(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return payment, nil
}

func (s *OrderService) GetPayments(ctx context.Context, req *pb.GetPaymentsRequest) (*pb.GetPaymentsResponse, error) {
	payments, err := s.Db.GetPayments(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return payments, nil
}