package migrations

import (
	"log"
	"os"
	"websocket/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	var err error
	DB, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = DB.Migrator().AutoMigrate(&models.Users{}, &models.Chatrooms{}, &models.Chatroom_users{}, &models.Messages{}); err != nil {
		log.Println(err)
		return nil, err
	}

	return DB, nil

}
