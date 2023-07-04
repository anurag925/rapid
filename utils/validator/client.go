package validator

import (
	"context"
	"rapid/utils/logger"
	"sync"

	"github.com/go-playground/validator/v10"
)

type ValidationRegistrationFunc func(v *validator.Validate) error

var (
	v     *validator.Validate
	vOnce sync.Once
) // ValidatorService validator service

func Client(ctx context.Context, f ValidationRegistrationFunc) *validator.Validate {
	vOnce.Do(func() {
		v = validator.New()
		if err := f(v); err != nil {
			logger.Error(ctx, "validation registration error", err)
		}
	})
	if v == nil {
		logger.Error(context.Background(), "validator is nil")
		panic("validator is nil even after init")
	}
	return v
}

func ValidateStruct(ctx context.Context, s any) error {
	return v.StructCtx(ctx, s)
}
