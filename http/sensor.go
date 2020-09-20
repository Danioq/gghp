package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Danioq/gghp/sensors"
	"github.com/gorilla/mux"
)

func AddSensor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	value, _ := strconv.ParseFloat(vars["value"], 64)
	sensors.AddSensor(name, value)
}

func GetSensor(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	sensor, err := sensors.GetSensor(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	response, _ := json.Marshal(sensor)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
