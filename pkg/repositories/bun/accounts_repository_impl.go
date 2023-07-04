package bun

import (
	"context"
	"rapid/pkg/models"
	"rapid/pkg/repositories"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type accountRepositoryImpl struct {
	*RepositoryImpl[models.Account]
}

var _ repositories.AccountRepository = (*accountRepositoryImpl)(nil)

func NewAccountRepositoryImpl(db *bun.DB) *accountRepositoryImpl {
	return &accountRepositoryImpl{RepositoryImpl: NewRepositoryImpl[models.Account](db)}
}

func (r *accountRepositoryImpl) FindByEmail(ctx context.Context, email string) (o models.Account, err error) {
	err = r.NewSelect().Where("email = ?", email).Scan(ctx, &o)
	return
}

func (r *accountRepositoryImpl) Create(ctx context.Context, o *models.Account) (err error) {
	if o.Password.Valid {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(o.Password.String), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		o.Password.SetValid(string(hashedPassword))
	}
	err = r.Create(ctx, o)
	return
}
