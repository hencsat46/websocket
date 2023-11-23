package usecase

import (
	"log"
	"strconv"
	"websocket/internal/api/handlers"
	"websocket/internal/models"
)

type usecase struct {
	repo RepositoryInterfaces
}

type RepositoryInterfaces interface {
	SignIn(string, string) (int, error)
	SignUp(string, string) error
	GetMessages() ([]models.Messages, error)
	WriteMessage(int, string) error
}

func NewUsecase(repo RepositoryInterfaces) handlers.UsecaseInterfaces {
	return &usecase{repo: repo}
}

func (u *usecase) SignIn(username, password string) (int, error) {

	userId, err := u.repo.SignIn(username, password)
	log.Println(userId)
	if err != nil {
		log.Println(err)
	}
	return userId, err

}

func (u *usecase) SignUp(username, password string) error {

	if err := u.repo.SignUp(username, password); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u *usecase) MessageDB(userId string, message string) error {

	userInt, _ := strconv.Atoi(userId)

	if err := u.repo.WriteMessage(userInt, message); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u *usecase) GetMessages() ([]models.Messages, error) {
	messagesArray, err := u.repo.GetMessages()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return messagesArray, nil
}
