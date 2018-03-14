package service


import (
	"golang.org/x/crypto/bcrypt"
	"github.com/mokoshi/simple-chat/go-api-server/models"
	"github.com/mokoshi/simple-chat/go-api-server/db"
)

func CreateUser(name string, password string) (user models.User, err error) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	hash := string(bytes)

	user = models.User{Name: name, Password: hash}
	if err = db.GetConnection().Create(&user).Error; err != nil {
		return
	}

	return
}

func VerifyUser(name string, password string) (user models.User, err error) {
	if err = db.GetConnection().Where("name = ?", name).First(&user).Error; err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return
	}

	return
}
