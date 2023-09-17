package model

import "time"

type Weather struct {
	StartTime             *time.Time `json:"start_time"`
	EndTime               *time.Time `json:"end_time"`
	Country               string     `json:"country"`
	City                  string     `json:"city"`
	Timezone              string     `json:"timezone"`
	TimezoneOffsetSeconds int64      `json:"timezone_offset_seconds"`
	Condition             string     `json:"condition"`
	Icon                  string     `json:"icon"`

	AvgTempC              float64    `json:"temp_c"`
	AvgTempF              float64    `json:"temp_f"`
	MinTempC              float64    `json:"min_temp_c"`
	MinTempF              float64    `json:"min_temp_f"`
	MaxTempC              float64    `json:"max_temp_c"`
	MaxTempF              float64    `json:"max_temp_f"`

	AvgWindKph            float64    `json:"avg_wind_kph"`
	AvgWindMph            float64    `json:"avg_wind_mph"`
	MinWindKph            float64    `json:"min_wind_kph"`
	MinWindMph            float64    `json:"min_wind_mph"`
	MaxWindKph            float64    `json:"max_wind_kph"`
	MaxWindMph            float64    `json:"max_wind_mph"`

	HourlyWeathers        []Weather  `json:"hourly_weathers"`
}
