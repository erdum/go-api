package middlewares

import (
	"net/http"
	// "fmt"
	"reflect"

	"github.com/labstack/echo/v4"
)

func Validate(payload interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// test := reflect.New(reflect.TypeOf(payload).Elem()).Interface()

			if err := c.Bind(payload); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			if err := c.Validate(payload); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			c.Set("valid_payload", payload.(reflect.TypeOf(payload).Elem(payload)))

			return next(c)
		}
	}
}
