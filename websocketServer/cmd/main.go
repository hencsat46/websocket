package main

import (
	"websocket/internal/api/server"
	"websocket/internal/migrations"
	"websocket/internal/pkg/env"
)

func main() {

	env.InitEnv()
	migrations.InitDB()

	server.Run()
}
