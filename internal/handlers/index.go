package handlers

import (
	"net/http"

	"github.com/slingercode/go-http/internal/web/pages"
)

func HandleIndex(w http.ResponseWriter, req *http.Request) {
	stack := []string{"go", "net/http", "templ", "pnpm", "tailwind", "htmx"}

	pages.IndexPage(stack).Render(req.Context(), w)
}
