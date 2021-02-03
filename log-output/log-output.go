package main

import (
	"github.com/gorilla/mux"
	. "github.com/paynejacob/rancher-logging-examples/log-output/pkg/index"
	"log"
	"net/http"
)

func main() {
	service := NewIndexService()

	r := mux.NewRouter()
	r.HandleFunc("/", service.ListIndices).Methods("GET")
	r.HandleFunc("/{index_name}/", service.ListLogs).Methods("GET")
	r.HandleFunc("/{index_name}/", service.WriteLog).Methods("POST")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":80", nil))
}
