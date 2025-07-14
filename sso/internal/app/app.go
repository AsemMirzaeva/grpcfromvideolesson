package app

import (
	"log/slog"
	authrpc "sso/internal/grpc/auth"

	"google.golang.org/grpc"
)

type App struct {
	log *slog.Logger
	gRPCServer *grpc.Server
	port string
}

func New(
	log *slog.Logger,
	port int,
) *App {
	gRPCServer := grpc.NewServer()

	authrpc.
}