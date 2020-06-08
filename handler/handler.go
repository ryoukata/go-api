package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Hello return echo.HandlerFunc
func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}

// ParamShow return echo.HandlerFunc
func ParamShow() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		return c.String(http.StatusOK, "Hello "+username)
	}
}

// JsonGet return echo.HandlerFunc
func JsonGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		jsonMap := map[string]string{
			"foo":  "bar",
			"hoge": "fuga",
		}
		return c.JSON(http.StatusOK, jsonMap)
	}
}
