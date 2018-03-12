package controller

import (
	"github.com/labstack/echo"
	"github.com/mokoshi/go-simple-chat/service"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/mokoshi/go-simple-chat/middleware"
)

func ListMessages(c echo.Context) (e error) {
	messages, _ := service.GetMessages()

	return c.JSON(http.StatusOK, messages)
}

func PostMessages(c echo.Context) (e error) {
	// JWT
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JWTClaims)
	userId := claims.UserId

	req := struct {
		Text string `form:"text"`
	}{}

	if e = c.Bind(&req); e != nil {
		return
	}

	messageId, _ := service.CreateMessage(userId, req.Text)

	return c.JSON(http.StatusOK, echo.Map{"message_id": messageId})
}
