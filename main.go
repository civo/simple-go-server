package main

import (
	"embed"
        "html/template"
	"fmt"
	"log"
	"net/http"
	"os"
)

//go:embed static
var static embed.FS

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
        t, err := template.ParseFS(static, "static/demo.html")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Starting server on port", port)
        var staticFS = http.FS(static)
	fs := http.FileServer(staticFS)

	// Serve static files
	http.Handle("/static/", fs)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
                t.Execute(w,"" )
	})
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		log.Fatal(err)
	}
}
