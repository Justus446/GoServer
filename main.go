package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)



	// fmt.Fprintf(w, "Welcome, to the HomePage justus!")
	o := os.Stdout
	fmt.Fprintf(o, "Server running!")
	if r.Method == "GET" {
		fmt.Println("HTTP method GET")
	}
	fmt.Println("Endpoint Hit: homePage")
	fmt.Println("Endpoint Hit: path", r.URL.Path,"\n header....",r.Header,"\n Body.....",r.Body,"\nMethod>>>>>",r.Method,"\n Query>>>>>",r.URL.Query())

}

func main() {
	// Code
	http.HandleFunc("/", index)
	http.ListenAndServe(":8089", nil)
}