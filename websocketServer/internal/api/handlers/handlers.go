package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"websocket/internal/models"
	customerrors "websocket/internal/pkg/customErrors"
	"websocket/internal/usecase"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
)

func Websocket(ctx echo.Context) error {
	ws, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	defer ws.Close()

	for {

		_, message, err := ws.ReadMessage()

		if err != nil {
			log.Println(err)
			break
		}

		if err := ws.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Println(err)
		}
		fmt.Printf("%s\n", message)
	}

	return nil

}

func SignUp(ctx echo.Context) error {
	var request SignDTO = SignDTO{"", ""}

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Wrong json"})
	}

	if err := usecase.SignUp(request.Username, request.Password); err != nil {
		if errors.Is(err, customerrors.ErrRepeat) {
			return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "This user already exists"})
		}

		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: "Sign Up ok"})
}

func SignIn(ctx echo.Context) error {

	var request SignDTO = SignDTO{"", ""}
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Wrong json"})
	}

	userId, err := usecase.SingIn(request.Username, request.Password)

	if err != nil {
		if errors.Is(err, customerrors.ErrEmpty) {
			return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "This user doesn't exists"})
		}
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: userId})

}
