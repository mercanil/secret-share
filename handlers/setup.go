package handlers

import (
	"embed"
	"net/http"
)

//go:embed static
var staticFS embed.FS

func SetupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/healthcheck", healthCheckHandler)
	mux.HandleFunc("/api", apiHandler)
	mux.HandleFunc("/", indexHandler)

	fileServer := http.FileServer(http.FS(staticFS))
	mux.Handle("/static/", fileServer)
}
