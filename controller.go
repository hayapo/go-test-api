package api

import (
	"encoding/json"
	"net/http"
)

var books []Book

type Controller interface {
	GetAllBooks(w http.ResponseWriter, _ *http.Request)
}

type controller struct{}

func NewController() Controller{
	return &controller{}
}

func init() {
	books = append(books, Book{Id: 1, Title: "吾輩は猫である", Author: "夏目漱石", Year: 1906, Genre: "novel"})
	books = append(books, Book{Id: 2, Title: "走れメロス", Author: "太宰治", Year: 1940, Genre: "novel"})
}

func (c *controller) GetAllBooks(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	outs, _ := json.MarshalIndent(books, "", "\t")
	w.Write(outs)
}
