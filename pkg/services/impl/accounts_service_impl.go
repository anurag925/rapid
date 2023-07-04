package impl

import (
	"context"
	"rapid/app"
	"rapid/pkg/models"
	"rapid/pkg/repositories"
	"rapid/pkg/services"
	"rapid/utils/jwt"

	goJwt "github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
)

type accountServiceImpl struct {
	accountRepo repositories.AccountRepository
}

func NewAccountService(repo repositories.AccountRepository) *accountServiceImpl {
	return &accountServiceImpl{accountRepo: repo}
}

func (s *accountServiceImpl) GetAccount(ctx context.Context, id int64) (models.Account, error) {
	return s.accountRepo.FindById(ctx, id)
}
func (s *accountServiceImpl) GetAccountByEmail(ctx context.Context, email string) (models.Account, error) {
	return s.accountRepo.FindByEmail(ctx, email)
}

func (s *accountServiceImpl) Create(ctx context.Context, a *models.Account) error {
	return s.accountRepo.Create(ctx, a)
}

func (s *accountServiceImpl) Login(ctx context.Context, req services.LoginRequest) (services.LoginResponse, error) {
	account, err := s.accountRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return services.LoginResponse{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Password.String), []byte(req.Password))
	if err != nil {
		return services.LoginResponse{}, err
	}
	token, err := s.generateUserToken(ctx, account)
	if err != nil {
		return services.LoginResponse{}, err
	}
	return services.LoginResponse{
		Account: account,
		Token:   token,
	}, nil
}

func (s *accountServiceImpl) generateUserToken(ctx context.Context, a models.Account) (string, error) {
	claims := goJwt.MapClaims{
		"account_id": a.ID,
		"email":      a.Email,
		"role":       a.Type.String(),
	}
	return jwt.Encode(ctx, claims, []byte(app.Config().JwtSecret))
}
