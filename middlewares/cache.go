package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/coocood/freecache"
)

var globalCache *freecache.Cache

func Cache() echo.MiddlewareFunc {

	if globalCache == nil {
		globalCache = freecache.NewCache(1024 * 1024 * 10) // 10Mb allocated
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("cache", globalCache)

			return next(c)
		}
	}
}
