package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
)

const webBaseDir = "web"
const port = 8888
const defaultPage = "index.html"

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func connectHttpServer() {

	router := mux.NewRouter().StrictSlash(true)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	log.Printf("http listening on port %d", port)
	http.Handle("/", router)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
