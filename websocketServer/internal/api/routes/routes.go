package routes

import (
	"websocket/internal/api/handlers"

	"github.com/labstack/echo/v4"
)

func CreateRoutes(e *echo.Echo) {
	e.GET("/ws", handlers.Websocket)
	e.POST("/signin", handlers.SignIn)
	e.POST("/signup", handlers.SignUp)
}
