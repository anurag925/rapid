package impl

import (
	"context"
	"testing"

	"github.com/anurag925/rapid/app"
	"github.com/anurag925/rapid/app/core"
)

func init() {
	app.New(core.GetBackendApp())
}
func Test_PrintConfigs(t *testing.T) {
	NewHealthCheckServiceImpl().PrintConfigs(context.Background())
}

func Test_HealthCheck(t *testing.T) {
	if NewHealthCheckServiceImpl().HealthCheck(context.Background()) {
		t.Fail()
	}
}
