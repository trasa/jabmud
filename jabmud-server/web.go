package main
import (
	"log"
	"net/http"
	"fmt"
	"io"
	"time"
)

const webBaseDir = "web"
const port = 8888
const defaultPage = "index.html"

func ConnectHttpServer() {
	http.HandleFunc("/", render)
	log.Printf("http listening on port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
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