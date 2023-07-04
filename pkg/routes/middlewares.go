package routes

import (
	"context"
	"github.com/anurag925/rapid/app"
	"github.com/anurag925/rapid/pkg/handlers"
	"github.com/anurag925/rapid/utils/logger"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	return middleware.RequestIDWithConfig(
		middleware.RequestIDConfig{
			RequestIDHandler: func(ctx echo.Context, s string) {
				ctx.Set("context", context.WithValue(
					ctx.Request().Context(), logger.ContextKeyValues,
					logger.ContextValue{
						logger.ContextKeyRequestID: s,
						logger.ContextKeyAccountID: ctx.Get("account_id"),
					}),
				)
			},
		},
	)
}

type JwtCustomClaims struct {
	AccountID int64  `json:"account_id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	jwt.RegisteredClaims
}

func AuthMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			SigningKey: []byte(app.Config().JwtSecret),
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(JwtCustomClaims)
			},
			ErrorHandler: func(c echo.Context, err error) error {
				return handlers.UnauthorizedResponse(c, "unauthorized account", err)
			},
		})
}
