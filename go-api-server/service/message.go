package service

import (
	"github.com/mokoshi/simple-chat/go-api-server/db"
	"github.com/mokoshi/simple-chat/go-api-server/models"
)

func GetMessages() (messages []models.Message, err error) {

	if err = db.GetConnection().Find(&messages).Error; err != nil {
		return
	}

	return
}

func CreateMessage(userId uint, text string) (id uint, err error) {
	message := models.Message{UserId: userId, Text: text}
	if err = db.GetConnection().Create(&message).Error; err != nil {
		return
	}

	id = message.ID
	return
}
