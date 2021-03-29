package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	serverPort := envhelper.GetEnvString("PORT", "4000")
	contentDir := envhelper.GetEnvString("CONTENT_DIR", ".")
	muxrouter := mux.NewRouter()

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
}
