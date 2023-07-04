package handlers

import (
	"github.com/anurag925/rapid/pkg/models"

	"github.com/labstack/echo/v4"
)

func Account(c echo.Context) models.Account {
	return c.Get("account").(models.Account)
}
