package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/mokoshi/simple-chat/go-api-server/controller"
	"github.com/mokoshi/simple-chat/go-api-server/db"
	"github.com/mokoshi/simple-chat/go-api-server/models"
	m "github.com/mokoshi/simple-chat/go-api-server/middleware"
)

func main() {
	e := echo.New()

	db.GetConnection().AutoMigrate(&models.User{})
	db.GetConnection().AutoMigrate(&models.Message{})

	jwtHandler := middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &m.JWTClaims{},
		SigningKey: []byte("secret"),
	})

	e.POST("/signup", controller.Signup)

	e.POST("/login", controller.Login)

	e.GET("/messages", controller.ListMessages)

	e.POST("/messages", controller.PostMessages, jwtHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
