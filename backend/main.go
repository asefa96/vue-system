package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", getInfos).Methods("GET")

	fmt.Println("run on: http://localhost:8888")
	handler := cors.Default().Handler(r)
	http.ListenAndServe(":8888", handler)

}
