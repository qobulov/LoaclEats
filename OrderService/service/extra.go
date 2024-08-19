package service

import (
	pb "Orders/genproto/orders"
	"context"
)

func (s *OrderService) GetKitchenStatistics(ctx context.Context, req *pb.GetKitchenStatisticsRequest) (*pb.GetKitchenStatisticsResponse, error) {
	kitchenStatistics, err := s.Db.GetKitchenStatistics(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return kitchenStatistics, nil
}

func (s *OrderService) GetUserActivity(ctx context.Context, req *pb.GetUserActivityRequest) (*pb.GetUserActivityResponse, error) {
	userActivity, err := s.Db.GetUserActivity(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return userActivity, nil
}

func (s *OrderService) UpdateWorkingHours(ctx context.Context, req *pb.UpdateWorkingHoursRequest) (*pb.UpdateWorkingHoursResponse, error) {
	workingHours, err := s.Db.UpdateWorkingHours(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return workingHours, nil
}

func (s *OrderService) CreaterWorkingHours(ctx context.Context, req *pb.CreateWorkingHoursRequest) (*pb.CreateWorkingHoursResponse, error) {
	workingHours, err := s.Db.CreateWorkingHours(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return workingHours, nil
}