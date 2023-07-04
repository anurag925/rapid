package main

import (
	"context"
	"github.com/anurag925/rapid/app"
	"github.com/anurag925/rapid/app/core"
	"github.com/anurag925/rapid/utils/logger"
)

func main() {
	app.New(core.GetBackendApp())
	logger.Info(context.Background(), "App init done ...")
	if err := app.Worker().Instance().Start(); err != nil {
		logger.Fatal(context.Background(), "unable to start worker", "error", err)
	}
}
