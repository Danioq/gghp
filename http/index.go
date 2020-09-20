package http

import (
	"fmt"
	"net/http"
)

// Index return Hello World string
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
