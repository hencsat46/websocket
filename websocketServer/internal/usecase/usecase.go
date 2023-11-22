package usecase

import (
	"log"
	"strconv"
	"websocket/internal/database"
	"websocket/internal/models"
)

func SingIn(username, password string) (int, error) {

	userId, err := database.SignIn(username, password)
	log.Println(userId)
	if err != nil {
		log.Println(err)
	}
	return userId, err

}

func SignUp(username, password string) error {

	if err := database.SignUp(username, password); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func MessageDB(userId string, message string) error {

	userInt, _ := strconv.Atoi(userId)

	if err := database.WriteMessage(userInt, message); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetMessages() ([]models.Messages, error) {
	messagesArray, err := database.GetMessages()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return messagesArray, nil
}
