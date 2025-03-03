package controllers

import (
	"go-api/requests"
	"go-api/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) UpdateProfile(c echo.Context) error {
	payload := c.Get("valid_payload").(*requests.UpdateProfileRequest)

	response, err := uc.userService.UpdateProfile(c, payload)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (uc *UserController) GetProfile(c echo.Context) error {
	response, err := uc.userService.GetProfile(c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (uc *UserController) UpdateLocation(c echo.Context) error {
	payload := c.Get("valid_payload").(*requests.UpdateLocationRequest)

	response, err := uc.userService.UpdateLocation(c, payload)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
