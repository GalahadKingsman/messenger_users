package userservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/GalahadKingsman/messenger_users/internal/models"
	pb "github.com/GalahadKingsman/messenger_users/pkg/messenger_users_api"
)

func (s *Service) CreateUser(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	// Валидация входных данных
	if req.GetFirstName() == "" || req.GetEmail() == "" {
		return nil, errors.New("имя и email обязательны для заполнения")
	}

	// Создание пользователя
	user := models.User{
		Login:     req.GetLogin(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Email:     req.GetEmail(),
		Phone:     req.GetPhone(),
	}

	// Сохранение пользователя в базе данных
	id, err := s.userRepo.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании пользователя: %v", err)
	}

	// Формирование ответа
	return &pb.CreateResponse{
		Success: fmt.Sprintf("Пользователь успешно создан с ID: %s", id),
	}, nil
}
