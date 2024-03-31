package handlers

import "net/http"

func HandleCustom(value string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(value))
	}
}
