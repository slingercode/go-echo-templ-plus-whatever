package internal

import "net/http"

func newServer() http.Handler {
	mux := http.NewServeMux()

	routes(mux)

	var handler http.Handler = mux

	return handler
}
