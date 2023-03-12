package controller

import (
	"auto_reload_data/helper"
	"auto_reload_data/model"
)

func GetWeatherStatus() model.Weather {
	var weather model.Weather

	weather.WaterVal = helper.GenerateRandValue()
	weather.WindVal = helper.GenerateRandValue()

	if weather.WaterVal <= 5 {
		weather.WaterStatus = "Aman"
	} else if weather.WaterVal >= 6 && weather.WaterVal <= 8 {
		weather.WaterStatus = "Siaga"
	} else if weather.WaterVal > 8 {
		weather.WaterStatus = "Bahaya"
	} else {
		weather.WaterStatus = "Error"
	}

	if weather.WindVal <= 6 {
		weather.WindStatus = "Aman"
	} else if weather.WindVal >= 7 && weather.WindVal <= 15 {
		weather.WindStatus = "Siaga"
	} else if weather.WindVal > 15 {
		weather.WindStatus = "Bahaya"
	} else {
		weather.WindStatus = "Error"
	}

	return weather
}
