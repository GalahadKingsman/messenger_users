package userservice

import (
	"context"
	pb "github.com/GalahadKingsman/messenger_users/pkg/messenger_users_api"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.userRepo.GetUserByLogin(req.Login)
	if err != nil || user.Password == "" {
		return &pb.LoginResponse{Message: "Некорректный логин или пароль"}, nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return &pb.LoginResponse{Message: "Некорректный логин или пароль"}, nil
	}
	return &pb.LoginResponse{Message: "Ваш ID", UserId: int32(user.ID)}, nil
}
