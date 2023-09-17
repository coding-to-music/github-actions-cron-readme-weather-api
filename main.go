package main

import (
	"log/slog"
	"os"

	"github.com/coding-to-music/github-actions-cron-readme-weather-api/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var loggingLevel = new(slog.LevelVar)
	loggingLevel.Set(slog.LevelInfo)
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: loggingLevel}))
	slog.SetDefault(logger)

	root := &cobra.Command{}
	root.AddCommand(cmd.UpdateWeather("update-weather"))
	err := root.Execute()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
