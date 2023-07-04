package main

import (
	"context"
	"rapid/app"
	"rapid/app/core"
	"rapid/utils/logger"
)

func main() {
	app.New(core.GetBackendApp())
	logger.Info(context.Background(), "App init done ...")
	if err := app.Worker().Instance().Start(); err != nil {
		logger.Fatal(context.Background(), "unable to start worker", "error", err)
	}
}
