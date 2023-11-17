package usecase

import (
	"log"
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
