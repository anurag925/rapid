package services

import (
	"context"
	"rapid/pkg/models"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Account models.Account `json:"account"`
	Token   string         `json:"token"`
}

type AccountService interface {
	GetAccount(ctx context.Context, id int64) (models.Account, error)
	GetAccountByEmail(ctx context.Context, email string) (models.Account, error)
	Create(context.Context, *models.Account) error
	Login(context.Context, LoginRequest) (LoginResponse, error)
}
