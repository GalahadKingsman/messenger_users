package userservice

import (
	"context"
	"fmt"
	"github.com/GalahadKingsman/messenger_users/internal/jwt"
	pb "github.com/GalahadKingsman/messenger_users/pkg/messenger_users_api"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	// Генерация токена
	token, err := jwt.GenerateToken(fmt.Sprintf("%d", user.ID))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ошибка генерации токена: %v", err)
	}

	return &pb.LoginResponse{
		Message: "Успешный вход",
		UserId:  int32(user.ID),
		Token:   token,
	}, nil
}
