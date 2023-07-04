package main

import (
	"context"
	"github.com/anurag925/rapid/app"
	"github.com/anurag925/rapid/app/core"
	"github.com/anurag925/rapid/pkg/scripts"
	"github.com/anurag925/rapid/utils/logger"
)

func main() {
	app.New(core.GetBackendApp())
	logger.Info(context.Background(), "App init done ...")
	scripts.SendTransactionOtpMailDynamically()
}
