package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type QueryResponse struct {
	Message string `json:"message"`
	Method string `json:"method"`
	Query string `json:"query"`
}

func index(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("name")

	response := QueryResponse{
		Message: "Welcome to the HomePage" + query,
		Method: r.Method,
		Query: query,
	}


	// Parse and render template with data
	indexPage, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalf("Error while parsing the template: %v", err)
	}

	// Pass response to template
	err = indexPage.Execute(w, response)
	if err != nil {
		log.Printf("Error executing template: %v", err)
	}



	o := os.Stdout
	fmt.Fprintf(o, "Server running!")


}

func main() {
	// Code
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8085", nil))
}