package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"weather/config"
	models "weather/model"

	"github.com/labstack/echo/v4"
)

func SaveWeather(ctx echo.Context) error {
	db := config.GetDB()
	weather := models.Weather{}

	if err := ctx.Bind(&weather); err != nil {
		return err
	}

	json_map := make(map[string]interface{})
	err := json.NewDecoder(ctx.Request().Body).Decode(&json_map)
	if err != nil {
		return err
	}

	fmt.Println(json_map)

	var water int = json_map["water"].(int)
	var wind int = json_map["wind"].(int)

	if water <= 5 {
		fmt.Println("Water Status: Aman")
	} else if water >= 6 && water <= 8 {
		fmt.Println("Water Status: Siaga")
	} else {
		fmt.Println("Water status: Bahaya")
	}

	if wind <= 6 {
		fmt.Println("Wind Status: Aman")
	} else if wind >= 6 && water <= 8 {
		fmt.Println("Wind status: Siaga")
	} else {
		fmt.Println("Water status: Bahaya")
	}

	db.Debug().Create(&weather)
	fmt.Println("Insert success")

	return ctx.JSON(http.StatusOK, weather)
}
