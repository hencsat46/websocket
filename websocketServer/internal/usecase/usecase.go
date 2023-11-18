package usecase

import (
	"log"
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

func MessageDB(userId int, message string) error {

	if err := database.WriteMessage(userId, message); err != nil {
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
