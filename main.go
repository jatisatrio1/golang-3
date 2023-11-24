package main

import (
	"time"
	"weather/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	config.connect()

	go func() {
		for {
			time.Sleep(15 * time.Second)
			client.hitAPI()
		}
	}()

	e := echo.New()

	e.POST("/weather", controller.SaveWeather())
}
