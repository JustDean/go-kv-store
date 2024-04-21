package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/get/{key}", handleGet).Methods(http.MethodGet)
	r.HandleFunc("/set/{key}", handleSet).Methods(http.MethodPut)
	r.HandleFunc("/del/{key}", handleDel).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":9090", r))
}
