package repositories

import (
	"context"
	"github.com/anurag925/rapid/pkg/models"
)

type AccountRepository interface {
	FindById(context.Context, int64) (models.Account, error)
	FindByEmail(context.Context, string) (models.Account, error)
	FindAll(ctx context.Context) ([]models.Account, error)
	Create(context.Context, *models.Account) error
}
