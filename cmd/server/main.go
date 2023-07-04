package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"rapid/app"
	"rapid/app/core"
	"rapid/pkg/routes"
	"rapid/utils/logger"
)

func main() {
	app.New(core.GetBackendApp())
	logger.Info(context.Background(), "App init done ...")
	routes.Init()
	logger.Info(context.Background(), "Router init done ...")

	// go func() {
	// 	logger.Info(context.Background(), "Async schedular server starting ...")
	// 	if err := app.Worker().Instance().StartScheduler(); err != nil {
	// 		logger.Fatal(context.Background(), "shutting down the async schedular because", "error", err)
	// 	}
	// 	logger.Info(context.Background(), "Async schedular server shutting down ...")
	// }()

	// Start server
	go func() {
		logger.Info(context.Background(), "Http Server starting ...")
		if err := app.Server().Instance().Start(fmt.Sprintf(":%d", app.Config().Port)); err != nil && err != http.ErrServerClosed {
			logger.Fatal(context.Background(), "shutting down the server because", "error", err)
		}
		logger.Info(context.Background(), "Http Server shutting down ...")
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info(context.Background(), "Shutting down the application ...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	logger.Info(context.Background(), "Cancelling context ...")
	defer cancel()
	if err := Shutdown(ctx); err != nil {
		logger.Fatal(context.Background(), "error shutting down server", "error", err)
	}
	logger.Info(context.Background(), "Bye ...")
}

func Shutdown(ctx context.Context) error {
	if err := app.DB().Close(ctx); err != nil {
		return err
	}
	if err := app.Server().Close(ctx); err != nil {
		return err
	}
	if err := app.Cache().Close(ctx); err != nil {
		return err
	}
	if err := app.Worker().Close(ctx); err != nil {
		return err
	}
	return nil
}
