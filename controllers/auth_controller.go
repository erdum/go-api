package controllers

import (
	"go-api/services/auth"
	"go-api/requests"
	"go-api/models"
	"net/http"
	"reflect"
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthController struct {
	authService auth.AuthService
	tokenService auth.TokenService
	db *gorm.DB
}

func NewAuthController(
	authService auth.AuthService,
	tokenService auth.TokenService,
	db *gorm.DB,
) *AuthController {
	return &AuthController{
		authService: authService,
		tokenService: tokenService,
		db: db,
	}
}

func (ac *AuthController) Login(context echo.Context) error {
	payload := context.Get("valid_payload").(*requests.LoginRequest)

	userData, err := ac.authService.AuthenticateWithThirdParty(payload.IdToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	avatar := userData["avatar"]

	user := models.User{}
	ac.db.Where(models.User{Email: userData["email"]}).Assign(
		models.User{
			Name: userData["name"],
			UID: userData["uid"],
			Avatar: &avatar,
		},
	).FirstOrCreate(&user)

	token, err := ac.tokenService.GenerateToken(
		user.ID,
		fmt.Sprintf("%T", reflect.TypeOf(user)),
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return context.JSON(
		http.StatusOK,
		map[string]interface{}{"token": token, "user": user},
	)
}
