package server

import "net/http"

// NewServeMux returns a new instance of ServeMux with registered handlers.
func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /routes", RoutesHandler)
	return mux
}
