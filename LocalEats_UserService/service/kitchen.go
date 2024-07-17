package service

import (
	pb "AuthService/genproto/users"
	"context"
	"log/slog"
)

func (k *AuthServiceServer) GetKitchenByID(ctx context.Context, req *pb.GetKitchenByIDRequest) (*pb.GetKitchenByIDResponse, error) {
	kitchen, err := k.User.GetKitchenByID(req)
	if err != nil {
		k.Logger.Error("Error with getting kitchen", slog.Any("error", err))
		return nil, err
	}
	return kitchen, nil
}

func (k *AuthServiceServer) GetKitchens(ctx context.Context, req *pb.GetKitchensRequest) (*pb.GetKitchensResponse, error) {
	kitchens, err := k.User.GetAllKitchens(req)
	if err != nil {
		k.Logger.Error("Error with getting kitchens", slog.Any("error", err))
		return nil, err
	}
	return kitchens, nil
}

func (k *AuthServiceServer) CreateKitchen(ctx context.Context, req *pb.CreateKitchenRequest) (*pb.Status, error) {
	kitchen, err := k.User.CreateKitchen(req)
	if err != nil {
		k.Logger.Error("Error with creating kitchen", slog.Any("error", err))
		return nil, err
	}
	return kitchen, nil
}

func (k *AuthServiceServer) UpdateKitchen(ctx context.Context, req *pb.UpdateKitchenRequest) (*pb.Status, error) {
	kitchen, err := k.User.UpdateKitchen(req)
	if err != nil {
		k.Logger.Error("Error with updating kitchen", slog.Any("error", err))
		return nil, err
	}
	return kitchen, nil
}

func (k *AuthServiceServer) DeleteKitchen(ctx context.Context, req *pb.GetKitchenByIDRequest) (*pb.Status, error) {
	kitchen, err := k.User.DeleteKitchen(req)
	if err != nil {
		k.Logger.Error("Error with deleting kitchen", slog.Any("error", err))
		return nil, err
	}
	return kitchen, nil
}

func (k *AuthServiceServer) SearchKitchens(ctx context.Context, req *pb.SearchKitchensRequest) (*pb.SearchKitchensResponse, error) {
	kitchens, err := k.User.SearchKitchens(req)
	if err != nil {
		k.Logger.Error("Error with searching kitchens", slog.Any("error", err))
		return nil, err
	}
	return kitchens, nil
}