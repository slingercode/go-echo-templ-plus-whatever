package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleVote(ctx echo.Context) error {
  return ctx.String(http.StatusOK, "Vote API")
}
