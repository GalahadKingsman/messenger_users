package userservice

import (
	"github.com/GalahadKingsman/messenger_users/internal/repositories/user_repo"
	pb "github.com/GalahadKingsman/messenger_users/pkg/messenger_users_api" // Импорт сгенерированного пакета
)

type Service struct {
	pb.UnimplementedUserServiceServer

	userRepo *user_repo.Repo
}

func New(userRepo *user_repo.Repo) *Service {
	return &Service{
		userRepo: userRepo,
	}
}
