package database

import (
	"fmt"
	"time"
	"websocket/internal/migrations"
	"websocket/internal/models"
	customerrors "websocket/internal/pkg/customErrors"
	"websocket/internal/usecase"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository() usecase.RepositoryInterfaces {
	db, err := migrations.InitDB()
	if err != nil {
		return &repository{db: nil}
	}

	return &repository{db: db}
}

func (r *repository) SignIn(username, password string) (int, error) {

	var user_id uint

	count, err := checkRepeat(username, r.db)
	if err != nil {
		return -1, err
	}

	if count < 1 {
		return -1, customerrors.ErrEmpty
	}

	if err := r.db.Model(&models.Users{}).Select("user_id").Where("username = ? AND passwd = ?", username, password).Find(&user_id).Error; err != nil {
		return -1, err
	}

	return int(user_id), nil
}

func (r *repository) SignUp(username, password string) error {

	count, err := checkRepeat(username, r.db)

	if err != nil {
		return err
	}

	if count != 0 {
		return customerrors.ErrRepeat
	}

	if err = r.db.Create(&models.Users{Username: username, Passwd: password}).Error; err != nil {
		return err
	}

	return nil

}

func checkRepeat(username string, db *gorm.DB) (int, error) {

	var userCount int64

	if err := db.Model(&models.Users{}).Where("username = ?", username).Count(&userCount).Error; err != nil {
		return -1, err
	}

	return int(userCount), nil
}

func (r *repository) GetMessages() ([]models.Messages, error) {

	var recordsCount int64

	if err := r.db.Model(&models.Messages{}).Count(&recordsCount).Error; err != nil {
		return nil, err
	}

	messagesArr := make([]models.Messages, recordsCount)

	if err := r.db.Order("message_date ASC").Find(&messagesArr).Error; err != nil {
		return nil, err
	}

	return messagesArr, nil
}

func (r *repository) WriteMessage(userId int, message string) error {

	currentDate := time.Now()

	formatTime := fmt.Sprintf("%d-%d-%d %d:%d:%d",
		currentDate.Year(),
		currentDate.Month(),
		currentDate.Day(),
		currentDate.Hour(),
		currentDate.Minute(),
		currentDate.Second())

	if err := r.db.Create(&models.Messages{Message_text: message, Message_date: formatTime, Message_owner: userId}).Error; err != nil {
		return err
	}

	return nil
}
