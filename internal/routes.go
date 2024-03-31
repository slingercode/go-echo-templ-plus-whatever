package internal

import (
	"net/http"

	h "github.com/slingercode/go-http/internal/handlers"
)

func routes(mux *http.ServeMux) {
	staticFolder := http.FileServer(http.Dir("./internal/web/static"))

	mux.Handle("/static/", http.StripPrefix("/static", staticFolder))

	mux.HandleFunc("/", h.HandleIndex)
	mux.HandleFunc("/custom", h.HandleCustom("Custom message"))
}
