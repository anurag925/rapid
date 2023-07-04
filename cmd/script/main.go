package main

import (
	"context"
	"rapid/app"
	"rapid/app/core"
	"rapid/pkg/scripts"
	"rapid/utils/logger"
)

func main() {
	app.New(core.GetBackendApp())
	logger.Info(context.Background(), "App init done ...")
	scripts.SendTransactionOtpMailDynamically()
}
