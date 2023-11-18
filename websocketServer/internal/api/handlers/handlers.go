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

	usersConnections = make([]*websocket.Conn, 0, 10)
)

func Websocket(ctx echo.Context) error {

	ws, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}

	usersConnections = append(usersConnections, ws)

	defer ws.Close()

	for {
		var request = MessageDTO{-1, ""}

		if err := ws.ReadJSON(&request); err != nil {
			log.Println(err)
			break
		}

		for i := 0; i < len(usersConnections); i++ {
			if usersConnections[i] != ws {
				if err := usersConnections[i].WriteJSON(request); err != nil {
					log.Println(err)
				}
			}
		}

		fmt.Printf("Sender: %d, Message: %s\n", request.UserId, request.Message)
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
