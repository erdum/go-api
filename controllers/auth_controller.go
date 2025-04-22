package controllers

import (
	"go-api/services/auth"
	"go-api/requests"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService auth.AuthService
	tokenService auth.TokenService
}

func NewAuthController(
	authService auth.AuthService,
	tokenService auth.TokenService,
) *AuthController {
	return &AuthController{
		authService: authService,
		tokenService: tokenService,
	}
}

func (ac *AuthController) DevLogin(c echo.Context) error {
	payload := c.Get("valid_payload").(*requests.DevLoginRequest)

	response, err := ac.authService.DevLogin(c, payload)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (ac *AuthController) Register(c echo.Context) error {
	payload := c.Get("valid_payload").(*requests.RegisterRequest)

	response, err := ac.authService.Register(c, payload)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (ac *AuthController) Login(c echo.Context) error {
	payload := c.Get("valid_payload").(*requests.LoginRequest)

	response, err := ac.authService.Login(c, payload)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (ac *AuthController) SignOn(c echo.Context) error {
	payload := c.Get("valid_payload").(*requests.SignOnRequest)

	response, err := ac.authService.SignOn(c, payload)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (ac *AuthController) ForgetPassword(c echo.Context) error {
	payload := c.Get("valid_payload").(*requests.ResendOtpRequest)

	response, err := ac.authService.ForgetPassword(c, payload)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (ac *AuthController) UpdatePassword(c echo.Context) error {
	payload := c.Get("valid_payload").(*requests.UpdatePasswordRequest)

	response, err := ac.authService.UpdatePassword(c, payload)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (ac *AuthController) ResendOtp(c echo.Context) error {
	payload := c.Get("valid_payload").(*requests.ResendOtpRequest)

	response, err := ac.authService.ResendOtp(c, payload)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (ac *AuthController) VerifyOtp(c echo.Context) error {
	payload := c.Get("valid_payload").(*requests.VerifyOtpRequest)

	response, err := ac.authService.VerifyOtp(c, payload)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
