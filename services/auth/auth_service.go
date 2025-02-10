package auth

import (
	"go-api/requests"

	"github.com/labstack/echo/v4"
)

type AuthService interface {
	Register(echo.Context, *requests.RegisterRequest) (map[string]string, error)
	VerifyEmail(echo.Context, *requests.VerifyEmailRequest) (
		map[string]string,
		error,
	)
	ResendOtp(echo.Context,*requests.ResendOtpRequest) (
		map[string]string,
		error,
	)
	Login(echo.Context, *requests.LoginRequest) (map[string]string, error)
	SignOn(echo.Context, *requests.SignOnRequest) (
		map[string]string,
		error,
	)
}
