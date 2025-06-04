package userservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/GalahadKingsman/messenger_users/internal/database"
	"github.com/GalahadKingsman/messenger_users/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"

	pb "github.com/GalahadKingsman/messenger_users/pkg/messenger_users_api" // Импорт сгенерированного пакета
)

type Service struct {
	pb.UnimplementedUserServiceServer
}

func (s *Service) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var (
		query = "SELECT id, login, first_name, last_name, email, phone FROM users"
		args  []interface{}
		where []string
	)

	if req.Id != nil {
		where = append(where, "id = $1")
		args = append(args, *req.Id)
	}
	if req.Login != nil {
		where = append(where, "login = $1")
		args = append(args, *req.Login)
	}
	if req.FirstName != nil {
		where = append(where, "first_name = $1")
		args = append(args, *req.FirstName)
	}
	if req.LastName != nil {
		where = append(where, "last_name = $1")
		args = append(args, *req.LastName)
	}
	if req.Email != nil {
		where = append(where, "email = $1")
		args = append(args, *req.Email)
	}
	if req.Phone != nil {
		where = append(where, "phone = $1")
		args = append(args, *req.Phone)
	}

	// Добавляем WHERE, если есть условия
	if len(where) > 0 {
		query += " WHERE " + strings.Join(where, " AND ")
	}
	if len(where) == 0 {
		return nil, status.Error(codes.InvalidArgument, "хотя бы одно поле должно быть указано")
	}
	row := database.DB.QueryRowContext(ctx, query, args...)

	var user models.User
	err := row.Scan(&user.ID, &user.Login, &user.FirstName, &user.LastName, &user.Email, &user.Phone)
	if err != nil {
		return nil, fmt.Errorf("пользователь не найден: %v", err)
	}
	return &pb.GetUserResponse{Id: int64(user.ID),
		Login:     user.Login,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone}, nil
}

func New() *Service {
	return &Service{}
}
