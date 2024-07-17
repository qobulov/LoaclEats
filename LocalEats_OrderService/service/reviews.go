package service

import (
	pb "Orders/genproto/orders"
	"context"
)

func (s *OrderService) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.Review, error) {
	review, err := s.Db.CreateReview(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return review, nil
}

func (s *OrderService) GetReviews(ctx context.Context, req *pb.GetReviewsRequest) (*pb.GetReviewsResponse, error) {
	reviews, err := s.Db.GetReviews(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return reviews, nil
}

func (s *OrderService) UpdateReview(ctx context.Context, req *pb.UpdateReviewRequest) (*pb.Status, error) {
	status, err := s.Db.UpdateReview(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return status, nil
}

func (s *OrderService) DeleteReview(ctx context.Context, req *pb.Review) (*pb.Status, error) {
	status, err := s.Db.DeleteReview(req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return status, nil
}
