package env

import (
	"log"

	dotenv "github.com/joho/godotenv"
)

func InitEnv() {
	if err := dotenv.Load("../.env"); err != nil {
		log.Println("Cannot find .env file")
	}
}
