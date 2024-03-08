package main

import (
	"github.com/frcrod/watt-companion-api/api/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", handler.HandleLanding)
	e.Static("static", "web/static")
	e.Logger.Fatal(e.Start(":8080"))
}
