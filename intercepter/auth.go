package intercepter

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// BasicAuth return echo.MiddlewareFunc
func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username string, password string, c echo.Context) (bool, error) {
		if username == "ryoya" && password == "3091" {
			return true, nil
		}
		return false, nil
	})
}
