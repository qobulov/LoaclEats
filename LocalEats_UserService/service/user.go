package service

import (
	postgres "AuthService/Storage"
	pb "AuthService/genproto/proto"
	"AuthService/pkg/logger"
	"context"
	"database/sql"
	"log/slog"
)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
	Logger *slog.Logger
	User   *postgres.UserRepo
}

func NewAuthServiceServer(db *sql.DB) *AuthServiceServer {
	return &AuthServiceServer{
		User:   postgres.NewUserRepo(db),
		Logger: logger.NewLogger(),
	}
}

func (s *AuthServiceServer) GetProfile(ctx context.Context, req *pb.UserId) (*pb.GetProfileResponse, error) {
	user, err := s.User.GetProfile(req)
	if err != nil {
		s.Logger.Error("Error with getting user", slog.Any("error", err))
		return nil, err
	}
	return user, nil
}


func (s *AuthServiceServer) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.Status, error) {

	status, err := s.User.UpdateProfile(req)
	if err != nil {
		s.Logger.Error("Error with updating user", slog.Any("error", err))
		return nil, err
	}
	return status, nil
}

func (s *AuthServiceServer) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.Status, error) {

	status, err := s.User.ResetPassword(req)
	if err != nil {
		s.Logger.Error("Error with resetting password", slog.Any("error", err))
		return nil, err
	}
	return status, nil
}

func (s *AuthServiceServer) DeleteProfile(ctx context.Context, req *pb.UserId) (*pb.Status, error) {

	status, err := s.User.DeleteProfile(req.Id)
	if err != nil {
		s.Logger.Error("Error with deleting user", slog.Any("error", err))
		return nil, err
	}
	return status, nil
}
