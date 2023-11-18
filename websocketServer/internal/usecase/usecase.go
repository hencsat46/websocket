package usecase

import (
	"log"
	"strconv"
	"websocket/internal/database"
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

func MessageDB(userId, message string) error {
	userIdInt, err := strconv.Atoi(userId)

	if err != nil {
		log.Println(err)
		return err
	}

	if err := database.WriteMessage(userIdInt, message); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
