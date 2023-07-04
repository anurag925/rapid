package bun

import (
	"context"

	"github.com/uptrace/bun"
)

type RepositoryImpl[T any] struct {
	*bun.DB
}

func NewRepositoryImpl[T any](db *bun.DB) *RepositoryImpl[T] {
	return &RepositoryImpl[T]{DB: db}
}

// func (r *RepositoryImpl[T]) db(ctx context.Context) *bun.DB {
// 	return r.Tx
// }

func (r *RepositoryImpl[T]) FindAll(ctx context.Context) (t []T, err error) {
	err = r.DB.NewSelect().Model(&t).Scan(ctx)
	return
}

func (r *RepositoryImpl[T]) FindById(ctx context.Context, id int64) (t T, err error) {
	err = r.DB.NewSelect().Model(&t).Where("id = ?", id).Scan(ctx)
	return
}

// func (r *RepositoryImpl[T]) PreloadFindById(ctx context.Context, id int64) (t T, err error) {
// 	err = r.db.WithContext(ctx).Preload(clause.Associations).First(&t, id).Error
// 	return
// }

func (r *RepositoryImpl[T]) Create(ctx context.Context, o *T) (err error) {
	_, err = r.DB.NewInsert().Model(o).Exec(ctx)
	return
}

func (r *RepositoryImpl[T]) Save(ctx context.Context, o *T) (err error) {
	_, err = r.DB.NewUpdate().Model(o).Exec(ctx)
	return
}
