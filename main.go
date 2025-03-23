package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"app/data"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Server listening port")
	flag.Parse()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		pageData := data.GetPageData()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pageData)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pageData := data.GetPageData()
		tmplData := data.PrepareTemplateData(pageData)

		tmpl, err := template.ParseFiles("templates/template.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, tmplData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Server started at http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
