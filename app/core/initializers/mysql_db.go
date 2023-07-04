package initializers

import (
	"context"
	"database/sql"
	"rapid/app"
	"rapid/app/configs"
	"rapid/utils/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type MySQL struct {
	db *bun.DB
}

var _ app.SQL = (*MySQL)(nil)

func (p *MySQL) Init(ctx context.Context, c configs.Config, l logger.Logger) error {
	l.Info(ctx, "DB connection init ...")
	sqldb, err := sql.Open("mysql", c.DBUrl)
	if err != nil {
		return err
	}
	p.db = bun.NewDB(sqldb, mysqldialect.New())
	if err := p.db.PingContext(ctx); err != nil {
		return err
	}
	l.Info(ctx, "DB connection completed ...")
	return nil
}

func (p *MySQL) Instance() *bun.DB {
	return p.db
}

func (p *MySQL) Close(ctx context.Context) error {
	return p.db.Close()
}
