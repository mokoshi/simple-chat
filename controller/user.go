package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/mokoshi/go-simple-chat/service"
	"github.com/mokoshi/go-simple-chat/middleware"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func Signup(c echo.Context) (e error) {
	req := struct {
		Name     string `form:"name"`
		Password string `form:"password"`
	}{}

	if e = c.Bind(&req); e != nil {
		return
	}

	userId, _ := service.CreateUser(req.Name, req.Password)

	return c.JSON(http.StatusOK, echo.Map{"user_id": userId})
}

func Login(c echo.Context) (e error) {
	req := struct {
		Name     string `form:"name"`
		Password string `form:"password"`
	}{}

	if e = c.Bind(&req); e != nil {
		return
	}

	userId, err := service.VerifyUser(req.Name, req.Password)

	if err != nil {
		return c.NoContent(http.StatusForbidden)
	}

	claims := &middleware.JWTClaims{UserId: userId}
	claims.ExpiresAt = time.Now().Add(time.Hour * 72).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.NoContent(http.StatusForbidden)
	}

	return c.JSON(http.StatusOK, echo.Map{"token": t})
}
