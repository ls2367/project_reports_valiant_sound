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
	http.HandleFunc("/1", handler1)
	http.ListenAndServe(":"+getEnv("PORT", "8080"), nil)
}

func handler1(w http.ResponseWriter, r *http.Request) {
	path := "project-report-1.html"
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Fatalln(err)
		return
	}
	t.Execute(w, nil)
}
