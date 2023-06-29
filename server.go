package main

import (
	"log"
	"net/http"

	"github.com/mercanil/secret-share.git/handlers"
)

func main() {
	mux := http.NewServeMux()
	handlers.SetupHandlers(mux)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
