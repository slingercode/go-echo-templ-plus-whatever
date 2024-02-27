package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/slingercode/go-echo-templ-plus-whatever/views/pages"
)

func HandleIndex(ctx echo.Context) error {
  thisDataIsBeingFetchedBySomeDatabase := "Noel"

  return RenderPage(pages.Index(thisDataIsBeingFetchedBySomeDatabase), ctx)
}
