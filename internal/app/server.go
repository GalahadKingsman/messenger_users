package app

import (
	"fmt"
	"github.com/GalahadKingsman/messenger_users/internal/config"
	"github.com/GalahadKingsman/messenger_users/internal/database"
	"google.golang.org/grpc"
	"net"
)

func Run(cfg *config.Config) error {
	if cfg == nil {
		return fmt.Errorf("config is nil")
	}

	// 1. Подключаемся к БД
	db, err := database.Init(cfg)
	if err != nil {
		return fmt.Errorf("failed to init DB: %v", err)
	}
	defer db.Close()

	// 2. Запускаем gRPC-сервер
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	return grpcServer.Serve(lis)
}
