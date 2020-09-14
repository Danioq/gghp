package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type router struct {
	client *mux.Router
}


func (router router) setupRoutes() {
	router.client.HandleFunc("/{q}", Index)
}

func setupRouter() {
	r := router{mux.NewRouter()}
	r.setupRoutes()
	http.Handle("/", r.client)
}
