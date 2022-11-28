package main

import (
	"net/http"
	"api"
)

func main() {
	mux := http.NewServeMux()
	c := api.NewController()
	mux.Handle("/all", http.HandlerFunc(c.GetAllBooks))
	http.ListenAndServe(":8080", mux)
}
