package http

import (
	"log"
	"net/http"
)

// RunHTTPServer setups and runs the http server
func RunHTTPServer() {

	setupRouter()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
