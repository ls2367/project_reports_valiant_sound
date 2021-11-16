package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	http.HandleFunc("/", handlerindex)
	http.HandleFunc("/1", handler1)
	http.HandleFunc("/2", handler2)
	http.HandleFunc("/3", handler3)
	http.ListenAndServe(":"+getEnv("PORT", "8080"), nil)
}

func handlerindex(w http.ResponseWriter, r *http.Request) {
	path := "project-reports.html"
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Fatalln(err)
		return
	}
	t.Execute(w, nil)
}

func handler1(w http.ResponseWriter, r *http.Request) {
	path := "project-reports.1.html"
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Fatalln(err)
		return
	}
	t.Execute(w, nil)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	path := "project-reports.2.html"
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Fatalln(err)
		return
	}
	t.Execute(w, nil)
}

func handler3(w http.ResponseWriter, r *http.Request) {
	path := "project-reports.3.html"
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Fatalln(err)
		return
	}
	t.Execute(w, nil)
}
