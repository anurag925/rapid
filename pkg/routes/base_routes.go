package routes

import (
	"encoding/json"
	"os"

	"github.com/anurag925/rapid/app"
	"github.com/anurag925/rapid/app/configs"
	"github.com/anurag925/rapid/pkg/handlers"
)

func Init() {
	serverInstance := app.Server()
	baseRoutes(serverInstance)
	printRoutes(serverInstance)
}

// baseRoutes
func baseRoutes(s app.HttpServer) {
	server := s.Instance()
	healthCheckController := handlers.NewHealthCheckController()
	server.GET("/health_check", healthCheckController.Ping)
	server.GET("/hello", healthCheckController.Hello)

	// add request id to the request and other request related data
	server.Use(LoggerMiddleware())
}

// printRoutes writes the routes to a file for debugging
func printRoutes(s app.HttpServer) error {
	if app.Config().Env != configs.Development {
		return nil
	}
	data, err := json.MarshalIndent(s.Instance().Routes(), "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile("routes.json", data, 0644); err != nil {
		return err
	}
	return nil
}
