package models

import "time"

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `gorm:"type:varchar(32);unique_index" json:"name"`
	Password  string     `gorm:"type:varchar(255)" json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

