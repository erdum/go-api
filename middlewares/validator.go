package middlewares

import (
	"net/http"
	// "fmt"

	"github.com/labstack/echo/v4"
)

func Validate(payload interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if err := c.Bind(payload); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			if err := c.Validate(payload); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			c.Set("valid_payload", payload)

			return next(c)
		}
	}
}
