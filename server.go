package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/trigger_delay", func(c echo.Context) error {
    q := c.QueryParam("q")
    return c.String(http.StatusOK, "「"+q+"」の検索結果")
	})
	e.POST("/clicked", func(c echo.Context) error {
		return c.String(http.StatusOK, "Clicked!")
	})
	e.POST("/mouse_entered", func(c echo.Context) error {
		return c.String(http.StatusOK, "Mouse entered!")
	})
	e.Static("/", "static")
	e.Logger.Fatal(e.Start(":1323"))
}
