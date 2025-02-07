package auth

import (
	"go-api/requests"

	"github.com/labstack/echo/v4"
)

type AuthService interface {
	Register(echo.Context, *requests.RegisterRequest) (map[string]string, error)
	Login(echo.Context, *requests.LoginRequest) (map[string]string, error)
	SignOn(echo.Context, *requests.SignOnRequest) (
		map[string]string,
		error,
	)
}
