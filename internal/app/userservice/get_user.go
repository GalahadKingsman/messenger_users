package userservice

import (
	"context"
	"github.com/GalahadKingsman/messenger_users/internal/models"
	pb "github.com/GalahadKingsman/messenger_users/pkg/messenger_users_api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	filter := &models.GetUserFilter{
		Id:        req.Id,
		Login:     req.Login,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
	}

	users, err := s.userRepo.GetUsers(ctx, filter)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pbUsers := make([]*pb.GetUserResponse_User, 0, len(users))
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.GetUserResponse_User{
			Id:        user.ID,
			Login:     user.Login,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
		})
	}
	return &pb.GetUserResponse{
		Users: pbUsers,
	}, nil
}
