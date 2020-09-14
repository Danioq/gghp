package http

import (
	"github.com/Danioq/gghp/apps"
	"net/http"

	"github.com/gorilla/mux"
)


// Index return Hello World string
func Index(w http.ResponseWriter, r *http.Request) {
	q := mux.Vars(r)["q"]
	response := apps.List[q]
	response(w, r)
}
