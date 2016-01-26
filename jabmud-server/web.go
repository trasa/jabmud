package main
import (
	"log"
	"net/http"
	"html/template"
	"fmt"
)

const webBaseDir = "web"
const port = 8888

func ConnectHttpServer() {
	http.HandleFunc("/", render)
	log.Printf("http listening on port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}


func render(w http.ResponseWriter, r *http.Request) {
	requestedFile := r.URL.Path[1:]
	log.Printf("render: %s", requestedFile)

	templ := fmt.Sprintf("%s/%s", webBaseDir, requestedFile)
	t, err := template.ParseFiles(templ)
	if err != nil {
		log.Printf("template parsing error: %s - %s", requestedFile, err)
		http.NotFound(w, r)
		return
	}
	err = t.Execute(w, "") // TODO data ?
	if err != nil {
		log.Printf("template execution error: %s - %s", requestedFile, err)
		// some sort of 500 error
	}


/*
	// if name 'console.html' doesn't match template filename then stuff doesn't work.
	t := template.New("console.html")
	t, _ = t.ParseFiles("web/console.html")
	t.Execute(w, "")
	*/
}