package service

import (
	"github.com/mokoshi/simple-chat/go-api-server/db"
	"github.com/mokoshi/simple-chat/go-api-server/models"
)

const MaxInt = 4294967295

func GetOlderMessages(endId uint, limit uint) (messages []models.Message, err error) {

	if endId <= 0 {
		endId = MaxInt
	}
	if limit <= 0 {
		limit = 15
	}

	if err = db.GetConnection().Preload("User").Where("id <= ?", endId).Order("ID desc").Limit(limit).Find(&messages).Error; err != nil {
		return
	}

	return
}

func GetNewerMessages(startId uint, limit uint) (messages []models.Message, err error) {

	if startId <= 0 {
		startId = 0
	}
	if limit <= 0 {
		limit = 15
	}

	if err = db.GetConnection().Preload("User").Where("id >= ?", startId).Order("ID asc").Limit(limit).Find(&messages).Error; err != nil {
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
