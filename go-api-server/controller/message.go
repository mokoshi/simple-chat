package controller

import (
	"github.com/labstack/echo"
	"github.com/mokoshi/simple-chat/go-api-server/service"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/mokoshi/simple-chat/go-api-server/middleware"
	"github.com/mokoshi/simple-chat/go-api-server/models"
)

func ListMessages(c echo.Context) (e error) {
	req := struct {
		StartId int  `query:"start_id"`
		EndId   int  `query:"end_id"`
		Limit   uint `query:"limit"`
	}{}
	if e = c.Bind(&req); e != nil {
		return
	}

	var messages []models.Message
	if req.StartId > 0 {
		// start_id が指定されている場合は新しいメッセージを取得
		messages, _ = service.GetNewerMessages(req.StartId, req.Limit)
	} else {
		// それ以外の場合は古いメッセージを取得
		messages, _ = service.GetOlderMessages(req.EndId, req.Limit)
	}

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
