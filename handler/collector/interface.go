package collector

import (
	"context"

	"github.com/coding-to-music/github-actions-cron-readme-weather-api/model"
)

type WeatherService interface {
	Forecast(ctx context.Context, city string, days int) ([]model.Weather, error)
}
