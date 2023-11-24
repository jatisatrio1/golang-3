package main

import (
	"weather/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	config.connect()

	e := echo.New()

	e.POST("/weather", controller.SaveWeather())
}
