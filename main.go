package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ryoukata/go-api/handler"
	"github.com/ryoukata/go-api/intercepter"
)

func main() {

	router := NewRouter()

	// ミドルウェアの設定
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(intercepter.BasicAuth())

	// サーバ起動
	router.Logger.Fatal(router.Start(":8081"))
}

// NewRouter return *echo.Echo
func NewRouter() *echo.Echo {
	// RESTによるAPIサーバを構築するためechoを使用
	e := echo.New()

	e.GET("/", handler.Hello())

	// パスパラメータを使用する場合
	e.GET("/:username", handler.ParamShow())

	// JSONデータを返却
	e.GET("/json", handler.JsonGet())

	// 天気APIから天気のデータを取得して表示
	e.GET("/weather/:city", handler.GetWeatherByCity())

	return e
}
