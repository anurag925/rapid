package main

import (
	"github.com/anurag925/rapid/app"
	"github.com/anurag925/rapid/app/core"
)

func main() {
	app.New(core.GetBackendApp())
	// postgresql.Migrate(app.DB().Instance())
}
