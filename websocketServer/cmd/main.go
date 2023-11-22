package main

import (
	"context"
	"os/signal"
	"syscall"
	"websocket/internal/api/server"
	"websocket/internal/migrations"
	"websocket/internal/pkg/env"
)

func main() {

	env.InitEnv()
	migrations.InitDB()

	server.Run()
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT)

	defer stop()

	<-ctx.Done()
}
