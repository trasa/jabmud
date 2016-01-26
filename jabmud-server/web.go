package main
import (
	"log"
	"net/http"
	"html/template"
)


func ConnectHttpServer() {
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/", rootHandler)
	log.Println("http listening on port 8888")
	http.ListenAndServe(":8888", nil)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	// nothing happens
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("rootHandler: %s", r.URL.Path[1:])
	// if name 'console.html' doesn't match template filename then stuff doesn't work.
	t := template.New("console.html")
	t, _ = t.ParseFiles("web/console.html")
	t.Execute(w, "")
}