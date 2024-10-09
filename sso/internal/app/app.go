package app

import (
	"log/slog"
	grpcapp "testgrpc/sso/internal/app/grpc"
	"testgrpc/sso/internal/services/auth"
	"testgrpc/sso/internal/storage/sqlite"
	"time"
	// "sso/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
