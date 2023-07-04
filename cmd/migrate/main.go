package main

import (
	"rapid/app"
	"rapid/app/core"
)

func main() {
	app.New(core.GetBackendApp())
	// postgresql.Migrate(app.DB().Instance())
}
