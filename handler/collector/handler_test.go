package collector

import (
	_ "embed"
	"encoding/json"
	"os"
	"testing"

	"github.com/coding-to-music/github-actions-cron-readme-weather-api/model"
	"github.com/stretchr/testify/assert"
)

//go:embed testdata/weathers.json
var weathersData []byte

func TestGenerateReadme(t *testing.T) {
	var weathers []model.Weather
	err := json.Unmarshal(weathersData, &weathers)
	if err != nil {
		panic(err)
	}

	// Construct the path to data/test.txt relative to the test file
	readme, err := generateOutput(weathers, "../../template/README.md.template")
	assert.NoError(t, err)
	assert.NotNil(t, readme)
	assert.NotEmpty(t, *readme)
	os.WriteFile(".README.md", []byte(*readme), 0644)
}
