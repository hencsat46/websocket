package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
)

func main() {
	e := echo.New()
	e.GET("/ws", webst)
	e.Start(":3000")

}

func webst(ctx echo.Context) error {
	ws, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	defer ws.Close()

	for {
		if err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Websocket!")); err != nil {
			log.Println(err)
		}

		messageType, message, err := ws.ReadMessage()

		if err != nil {
			log.Println(err)
			break
		}

		fmt.Printf("%d: %s", messageType, message)
	}

	return nil

}
