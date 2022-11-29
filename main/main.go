package main

import (
	"net/http"
	"api"
)

func main() {
	mux := http.NewServeMux()
	c := api.NewController()
	mux.Handle("/all", http.HandlerFunc(c.GetAll))
	mux.Handle("/book", http.HandlerFunc(c.GetBook))
	mux.Handle("/create", http.HandlerFunc(c.PutBook))
	mux.Handle("/update", http.HandlerFunc(c.UpdateBook))
	mux.Handle("/delete", http.HandlerFunc(c.DeleteBook))
	http.ListenAndServe(":8080", mux)
}
