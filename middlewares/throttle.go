package middlewares

import (
	"fmt"
	"go-api/utils"
	"net/http"
	"strconv"

	"github.com/coocood/freecache"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Throttle(db *gorm.DB, num_of_reqs uint, per_sec uint) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cache, ok := c.Get("cache").(*freecache.Cache)
			if !ok || cache == nil {
				return echo.NewHTTPError(
					http.StatusInternalServerError,
					"Cache not available",
				)
			}

			user := c.Get("auth_user")

			var identifier string
			if user != nil {
				identifier = fmt.Sprintf("user:%v", user)
			} else {
				key := c.RealIP() + ":" + c.Request().UserAgent()
				identifier = utils.GetMD5Hash(key)
			}

			cacheKey := []byte("throttle:" + identifier)

			val, err := cache.Get(cacheKey)
			requestCount := 0
			if err == nil {
				// Convert stored value to integer
				requestCount, _ = strconv.Atoi(string(val))
			}

			if requestCount >= int(num_of_reqs) {
				return echo.NewHTTPError(
					http.StatusTooManyRequests,
					"Rate limit exceeded",
				)
			}

			requestCount++
			cache.Set(
				cacheKey,
				[]byte(strconv.Itoa(requestCount)),
				int(per_sec),
			)

			return next(c)
		}
	}
}
