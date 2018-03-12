package service


import (
	"golang.org/x/crypto/bcrypt"
	"github.com/mokoshi/go-simple-chat/models"
	"github.com/mokoshi/go-simple-chat/db"
)

func CreateUser(name string, password string) (id uint, err error) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	hash := string(bytes)

	user := models.User{Name: name, Password: hash}
	if err = db.GetConnection().Create(&user).Error; err != nil {
		return
	}

	id = user.ID
	return
}

func VerifyUser(name string, password string) (id uint, err error) {
	var user models.User

	if err = db.GetConnection().Where("name = ?", name).First(&user).Error; err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return
	}

	id = user.ID
	return
}
