package controllers

import (
	"go-api/services/auth"
	"net/http"
	"fmt"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authSevice auth.AuthService
}

func NewAuthController(authSevice auth.AuthService) *AuthController {
	return &AuthController{authSevice: authSevice}
}

func (ac *AuthController) Login(context echo.Context) error {
	payload := context.Get("valid_payload")
	fmt.Println(payload)

	// data, err := ac.authSevice.AuthenticateWithThirdParty(payload)
	// if err != nil {
	// 	return err
	// }

	// return context.JSON(http.StatusOK, data)
	return context.JSON(http.StatusOK, "data")
}
