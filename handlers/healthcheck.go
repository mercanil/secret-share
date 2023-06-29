package handlers

import (
	"fmt"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok")
}
