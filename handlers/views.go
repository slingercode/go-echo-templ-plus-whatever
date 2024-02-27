package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func RenderPage(page templ.Component, ctx echo.Context) error {
  return page.Render(ctx.Request().Context(), ctx.Response())
}
