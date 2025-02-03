package controllers

import (
	"go-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db: db}
}

func (uc *UserController) GetAllUsers(context echo.Context) error {
	users := []models.User{}
	uc.db.Find(&users)

	return context.JSON(http.StatusOK, users)
}

func (uc *UserController) GetUser(context echo.Context) error {
	userId := context.Param("id")
	user := models.User{}
	uc.db.First(&user, userId)

	return context.JSON(http.StatusOK, user)
}

func (uc *UserController) CreateUser(context echo.Context) error {
	type UserPayload struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required"`
	}
	payload := UserPayload{}

	if err := context.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := context.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := models.User{Name: payload.Name, Email: payload.Email}
	preferences := models.UserPreference{
		City: "Karachi",
		Country: "Pakistan",
		State: "Sindh",
		ZipCode: "75350",
		Address: "Gulshan",
	}

	uc.db.Create(&user)
	uc.db.Create(&preferences)
	user.UserPreferenceID = preferences.ID
	uc.db.Save(&user)
	// uc.db.Select("name", "email").Create(&user)

	return context.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(context echo.Context) error {
	userId := context.Param("id")
	user := models.User{}
	uc.db.First(&user, userId)

	type UserPayload struct {
		Name string `json:"name"`
	}
	payload := UserPayload{}

	if err := context.Bind(&payload); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	if payload.Name != "" && payload.Name != user.Name {
		user.Name = payload.Name
		uc.db.Save(&user)
	}

	return context.JSON(http.StatusOK, user)
}

func (uc *UserController) DeleteUser(context echo.Context) error {
	userId := context.Param("id")
	uc.db.Delete(&models.User{}, userId)

	return context.JSON(http.StatusOK, map[string]interface{}{"message": "User successfully deleted"})
}
