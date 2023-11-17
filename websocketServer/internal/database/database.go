package database

import (
	"websocket/internal/migrations"
	"websocket/internal/models"
	customerrors "websocket/internal/pkg/customErrors"
)

func SignIn(username, password string) (int, error) {

	var user_id uint

	count, err := checkRepeat(username)
	if err != nil {
		return -1, err
	}

	if count < 1 {
		return -1, customerrors.ErrEmpty
	}

	if err := migrations.DB.Model(&models.Users{}).Select("user_id").Where("username = ? AND passwd = ?", username, password).Find(&user_id).Error; err != nil {
		return -1, err
	}

	return int(user_id), nil
}

func SignUp(username, password string) error {

	count, err := checkRepeat(username)

	if err != nil {
		return err
	}

	if count != 0 {
		return customerrors.ErrRepeat
	}

	if err = migrations.DB.Create(&models.Users{Username: username, Passwd: password}).Error; err != nil {
		return err
	}

	return nil

}

func checkRepeat(username string) (int, error) {

	var userCount int64

	if err := migrations.DB.Model(&models.Users{}).Where("username = ?", username).Count(&userCount).Error; err != nil {
		return -1, err
	}

	return int(userCount), nil
}
