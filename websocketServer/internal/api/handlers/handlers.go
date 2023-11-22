package handlers

import (
	"errors"
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

	connMap = make(map[*websocket.Conn]struct{})
	//usersConnections = make([]*websocket.Conn, 0, 10)
)

func reciever(ws *websocket.Conn, in chan<- models.MessageSender) {
	for {
		var request = MessageDTO{"", ""}
		if err := ws.ReadJSON(&request); err != nil {
			log.Println(err)
			log.Println(request)
			break
		}
		usecase.MessageDB(request.UserId, request.Message)
		in <- models.MessageSender{Connection: ws, Message: request.Message}
	}
}

func sender(out <-chan models.MessageSender) {

	for i := range out {
		for k, _ := range connMap {
			if k != i.Connection.(*websocket.Conn) {
				if err := k.WriteJSON(MessageDTO{UserId: "4", Message: i.Message}); err != nil {
					log.Println(err)
				}
			}
		}
	}
}

func Websocket(ctx echo.Context) error {
	ws, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	//defer ws.Close()

	connMap[ws] = struct{}{}

	data := make(chan models.MessageSender)
	go reciever(ws, data)
	go sender(data)

	return nil
}

// func Websocket(ctx echo.Context) error {

// 	ws, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	usersConnections = append(usersConnections, ws)

// 	defer ws.Close()

// 	for {
// 		var request = MessageDTO{"", ""}

// 		if err := ws.ReadJSON(&request); err != nil {
// 			log.Println(err)
// 			log.Println(request)
// 			break
// 		}

// 		usecase.MessageDB(request.UserId, request.Message)

// 		for i := 0; i < len(usersConnections); i++ {
// 			if usersConnections[i] != ws {
// 				if err := usersConnections[i].WriteJSON(request); err != nil {
// 					log.Println(err)
// 				}
// 			}
// 		}

// 		fmt.Printf("Sender: %s, Message: %s\n", request.UserId, request.Message)
// 	}

// 	return nil

// }

func GetMessages(ctx echo.Context) error {
	messagesArr, err := usecase.GetMessages()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: []models.Messages{{Message_id: 0, Message_text: "Internal Server Error", Message_owner: -1, Message_date: "0-0-0 0:0:0"}}})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: messagesArr})
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
