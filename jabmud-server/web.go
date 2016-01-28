package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"time"
)

const webBaseDir = "web"
const port = 8888
const defaultPage = "index.html"

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func connectHttpServer() {

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Printf("http listening on port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

func render(w http.ResponseWriter, r *http.Request) {
	requestedFile := r.URL.Path[1:]
	// TODO switch this to use debug logging
	log.Printf("render: '%s'", requestedFile)
	if requestedFile == "" {
		requestedFile = defaultPage
	}

	f, err := http.Dir(webBaseDir).Open(requestedFile)
	defer f.Close()

	if err == nil {
		content := io.ReadSeeker(f)
		http.ServeContent(w, r, requestedFile, time.Now(), content)
		return
	}
	log.Printf("Error opening requested file %s: %v", requestedFile, err)
	// TODO return some sort of 500 error,
	// and return 404 only for when the file 'merely' can't be found..
	http.NotFound(w, r)
}
