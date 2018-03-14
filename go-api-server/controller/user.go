package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/mokoshi/simple-chat/go-api-server/service"
	"github.com/mokoshi/simple-chat/go-api-server/middleware"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/mokoshi/simple-chat/go-api-server/models"
)

func Signup(c echo.Context) (e error) {
	req := struct {
		Name     string `form:"name"`
		Password string `form:"password"`
	}{}

	if e = c.Bind(&req); e != nil {
		return
	}

	user, _ := service.CreateUser(req.Name, req.Password)

	t, err := GenerateJWT(user)
	if err != nil {
		return c.NoContent(http.StatusForbidden)
	}

	return c.JSON(http.StatusOK, echo.Map{"user": user, "jwt": t})
}

func Login(c echo.Context) (e error) {
	req := struct {
		Name     string `form:"name"`
		Password string `form:"password"`
	}{}

	if e = c.Bind(&req); e != nil {
		return
	}

	user, err := service.VerifyUser(req.Name, req.Password)

	if err != nil {
		return c.NoContent(http.StatusForbidden)
	}

	t, err := GenerateJWT(user)
	if err != nil {
		return c.NoContent(http.StatusForbidden)
	}

	return c.JSON(http.StatusOK, echo.Map{"user": user, "jwt": t})
}

func GenerateJWT(user models.User) (string, error) {
	claims := &middleware.JWTClaims{UserId: user.ID}
	claims.ExpiresAt = time.Now().Add(time.Hour * 72).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secret"))
}
