package handlers

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	indexContent := `
	<html>
	<head> <title> Deneme</title> </head>
	<body> Hello World </body>
	</html>
	`
	fmt.Fprintf(w, indexContent)
}
