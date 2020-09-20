package http

import (
	"github.com/gorilla/mux"
)

type router struct {
	client *mux.Router
}

func (router router) setupRoutes() {
	router.client.HandleFunc("/", Index)
	router.client.HandleFunc("/sensor/{name}", GetSensor)
	router.client.HandleFunc("/sensor/{name}/{value}", AddSensor)
}

func getHandler() *mux.Router {
	r := router{mux.NewRouter()}
	r.setupRoutes()
	return r.client
}
