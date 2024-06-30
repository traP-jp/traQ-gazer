package oapi

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func getUserIdFromSession(ctx echo.Context) (string, error) {
	userId := ctx.Request().Header.Get("X-Forwarded-User")
	if userId == "" {
		return "", errors.New("X-Forwarded-User is empty")
	}
	return userId, nil
}
