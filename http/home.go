package http

import (
	"fmt"
	"net/http"
)

// HelloWorld return Hello World string
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
