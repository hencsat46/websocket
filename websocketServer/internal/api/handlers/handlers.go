package handlers

import (
	"errors"
	"log"
	"net/http"
	"websocket/internal/models"
	customerrors "websocket/internal/pkg/customErrors"

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

type handler struct {
	usecase UsecaseInterfaces
}

type UsecaseInterfaces interface {
	SignIn(string, string) (int, error)
	SignUp(string, string) error
	MessageDB(string, string) error
	GetMessages() ([]models.Messages, error)
}

func NewHandler(usecase UsecaseInterfaces) *handler {
	return &handler{usecase: usecase}
}

func (h *handler) CreateRoutes(e *echo.Echo) {
	e.GET("/ws", h.Websocket)
	e.POST("/signin", h.SignIn)
	e.POST("/signup", h.SignUp)
	e.GET("/getmessages", h.GetMessages)
}

func (h *handler) reciever(ws *websocket.Conn, in chan<- models.MessageSender) {
	for {
		var request = MessageDTO{"", ""}
		if err := ws.ReadJSON(&request); err != nil {
			log.Println(err)
			log.Println(request)
			break
		}
		h.usecase.MessageDB(request.UserId, request.Message)
		in <- models.MessageSender{Connection: ws, Message: request.Message}
	}
}

func (h *handler) sender(out <-chan models.MessageSender) {

	for i := range out {
		for k := range connMap {
			if k != i.Connection.(*websocket.Conn) {
				if err := k.WriteJSON(MessageDTO{UserId: "4", Message: i.Message}); err != nil {
					log.Println(err)
				}
			}
		}
	}
}

func (h *handler) Websocket(ctx echo.Context) error {
	ws, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	//defer ws.Close()

	connMap[ws] = struct{}{}

	data := make(chan models.MessageSender)
	go h.reciever(ws, data)
	go h.sender(data)

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

func (h *handler) GetMessages(ctx echo.Context) error {
	messagesArr, err := h.usecase.GetMessages()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: []models.Messages{{Message_id: 0, Message_text: "Internal Server Error", Message_owner: -1, Message_date: "0-0-0 0:0:0"}}})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: messagesArr})
}

func (h *handler) SignUp(ctx echo.Context) error {
	var request SignDTO = SignDTO{"", ""}

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Wrong json"})
	}

	if err := h.usecase.SignUp(request.Username, request.Password); err != nil {
		if errors.Is(err, customerrors.ErrRepeat) {
			return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "This user already exists"})
		}

		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: "Sign Up ok"})
}

func (h *handler) SignIn(ctx echo.Context) error {

	var request SignDTO = SignDTO{"", ""}
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Wrong json"})
	}

	userId, err := h.usecase.SignIn(request.Username, request.Password)

	if err != nil {
		if errors.Is(err, customerrors.ErrEmpty) {
			return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "This user doesn't exists"})
		}
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: userId})

}
