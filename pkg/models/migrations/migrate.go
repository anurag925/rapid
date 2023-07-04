package migrations

import (
	"embed"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

//go:embed *.sql
var sqlMigrations embed.FS

func Migrate(db *bun.DB) {
	if err := migrate.NewMigrations().Discover(sqlMigrations); err != nil {
		panic(err)
	}
}
