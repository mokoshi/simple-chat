package models

import "time"

type Message struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	UserId    uint       `gorm:"type:int unsigned" json:"user_id"`
	Text      string     `gorm:"type:varchar(255)" json:"text"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

