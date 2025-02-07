package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/coocood/freecache"
)

func Cache() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cache := freecache.NewCache(1024 * 1024 * 10) // 10Mb allocated
			c.Set("cache", cache)

			return next(c)
		}
	}
}
