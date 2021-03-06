package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/JasurbekUz/ToDo-service/config"
	pb "github.com/JasurbekUz/ToDo-service/genproto"
	"github.com/JasurbekUz/ToDo-service/pkg/db"
	"github.com/JasurbekUz/ToDo-service/pkg/logger"
	"github.com/JasurbekUz/ToDo-service/service"
	"github.com/JasurbekUz/ToDo-service/storage"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "todo-service")
	defer func(l logger.Logger) {
		err := logger.Cleanup(l)
		if err != nil {
			log.Fatal("failed cleanup logger", logger.Error(err))
		}
	}(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectionToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	pgStorage := storage.NewStoragePg(connDB)

	todoService := service.NewTodoService(pgStorage, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, todoService)
	reflection.Register(s)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
