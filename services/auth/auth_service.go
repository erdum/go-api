package auth

import (
	"go-api/requests"

	"github.com/labstack/echo/v4"
)

type AuthService interface {
	DevLogin(echo.Context, *requests.DevLoginRequest) (
		map[string]string,
		error,
	)
	Register(echo.Context, *requests.RegisterRequest) (map[string]string, error)
	Login(echo.Context, *requests.LoginRequest) (map[string]string, error)
	SignOn(echo.Context, *requests.SignOnRequest) (
		map[string]string,
		error,
	)
	ForgetPassword(echo.Context, *requests.ResendOtpRequest) (
		map[string]string,
		error,
	)
	UpdatePassword(echo.Context, *requests.UpdatePasswordRequest) (
		map[string]string,
		error,
	)
	ResendOtp(echo.Context,*requests.ResendOtpRequest) (
		map[string]string,
		error,
	)
	VerifyOtp(echo.Context, *requests.VerifyOtpRequest) (
		map[string]string,
		error,
	)
}
