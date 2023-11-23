package main

import (
	"context"
	"os/signal"
	"syscall"
	"websocket/internal/api/handlers"
	"websocket/internal/database"
	"websocket/internal/pkg/env"
	"websocket/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	env.InitEnv()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "DELETE", "POST", "PUT"},
	}))

	repo := database.NewRepository()
	usecase := usecase.NewUsecase(repo)
	handlers := handlers.NewHandler(usecase)

	handlers.CreateRoutes(e)

	e.Start(":3000")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT)

	defer stop()

	<-ctx.Done()
}
