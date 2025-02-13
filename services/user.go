package services

import (
	"go-api/requests"
	"go-api/config"
	"go-api/models"
	// "go-api/utils"
	// "net/http"
	// "errors"
	// "fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
	appConfig *config.Config
}

func NewUserService(
	db *gorm.DB,
) *UserService {
	return &UserService{
		db: db,
		appConfig: config.GetConfig(),
	}
}

func (u *UserService) UpdateProfile(
	c echo.Context,
	payload *requests.UpdateProfileRequest,
) (map[string]string, error) {
	user := c.Get("auth_user").(models.User)

	if payload.Name != "" {
		user.Name = payload.Name
	}

	if payload.PhoneNumber != "" {
		user.PhoneNumber = payload.PhoneNumber
	}

	if payload.Gender != "" {
		user.Gender = payload.Gender
	}

	if payload.Avatar != "" {
	}

	if payload.Address != "" {
		user.Address.Address = payload.Address
	}

	if payload.City != "" {
		user.Address.City = payload.City
	}

	if payload.State != "" {
		user.Address.State = payload.State
	}

	if payload.Country != "" {
		user.Address.Country = payload.Country
	}

	if payload.ZipCode != "" {
		user.Address.ZipCode = payload.ZipCode
	}

	if payload.Lat != 0 {
		user.Address.Lat = payload.Lat
	}

	if payload.Long != 0 {
		user.Address.Long = payload.Long
	}

	u.db.Save(&user)

	return map[string]string{
		"message": "User profile successfully updated.",
	}, nil
}
