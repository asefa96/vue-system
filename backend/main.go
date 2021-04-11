package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", getInfos).Methods("GET")

	fmt.Println("run on: http://localhost:8888")

	http.ListenAndServe(":8888", r)

}
