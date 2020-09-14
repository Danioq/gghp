package http

import "net/http"

func setupRouter() {
	http.HandleFunc("/", HelloWorld)
}