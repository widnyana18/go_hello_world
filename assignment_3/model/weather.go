package model

type Weather struct {
	WaterVal    int    `json:"water_value"`
	WindVal     int    `json:"wind_value"`
	WaterStatus string `json:"water_status"`
	WindStatus  string `json:"wind_status"`
}
