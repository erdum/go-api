package middlewares

import (
	"go-api/services/auth"
	"go-api/models"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Authenticate(
	tokenService auth.TokenService,
	db *gorm.DB,
) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := extractToken(c)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			claims, err := tokenService.ValidateToken(token)
			if err != nil {
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"Invalid or expired token",
				)
			}

			user := models.User{}
			db.First(&user, claims.EntityID)

			c.Set("auth_user", user)

			return next(c)
		}
	}
}

func extractToken(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	
	if authHeader != "" {
		parts := strings.Split(authHeader, " ")

		if len(parts) == 2 && parts[0] == "Bearer" {
			return parts[1], nil
		}

		return "", errors.New("invalid authorization header format")
	}

	return "", errors.New("no token found")
}
