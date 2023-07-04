package impl

import (
	"context"

	"rapid/app"
	"rapid/pkg/services"
	"rapid/utils/logger"
)

type healthCheckService struct {
}

var _ services.HealthCheckService = (*healthCheckService)(nil)

func NewHealthCheckServiceImpl() *healthCheckService {
	return &healthCheckService{}
}

func (s *healthCheckService) PrintConfigs(ctx context.Context) {
	logger.Debug(ctx, "printing configs", "configs", app.Config())
}

// HealthCheck
func (s *healthCheckService) HealthCheck(ctx context.Context) bool {
	logger.Info(ctx, "health check")
	if app.Config().UP {
		logger.Info(ctx, "health check db")
		db := app.DB().Instance()
		logger.Info(ctx, "health check db", "val", db)
		if err := db.PingContext(ctx); err != nil {
			logger.Error(ctx, "error in db ping", "error", err)
			return false
		}
		logger.Info(ctx, "health check db done")
		return true
	}
	logger.Info(ctx, "app not up", "up", app.Config().UP)
	return false
}
