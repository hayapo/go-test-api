package api

import (
	"encoding/json"
	"strconv"
	"fmt"
	"net/http"
)

var books []Book

type Controller interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
	PutBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
}

type controller struct{}

func NewController() Controller {
	return &controller{}
}

func (c *controller) GetAll(w http.ResponseWriter, r *http.Request) {
	if len(books) == 0 {
		w.WriteHeader(http.StatusNoContent)
	}
	w.Header().Set("Content-Type", "application/json")
	output, _ := json.MarshalIndent(books, "", "\t")
	w.Write(output)
}

func (c *controller) GetBook(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("Id")
	if param == "" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	id, err := strconv.Atoi(param)
	
	if err != nil {
		return
	}

	for _, book := range books {
		if book.Id == id {
			w.Header().Set("Content-Type", "application/json")
			output, _ := json.MarshalIndent(book, "", "\t")
			w.Write(output)
			return
		}
	}
		w.Header().Set("Content-Type", "application/json")
		output, _ := json.MarshalIndent(books, "", "\t")
		w.Write(output)
}

func (c *controller) PutBook(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	
	if param == nil {
		return
	}

	id, _ := strconv.Atoi(param["Id"][0])
	year, _ := strconv.Atoi(param["Year"][0])
	newBook := Book{Id: id, Title: param["Title"][0], Author: param["Author"][0], Year: year, Genre: param["Genre"][0]}
	books = append(books, newBook)
	
	w.Header().Set("Content-Type", "application/json")
	output, _ := json.MarshalIndent(newBook, "", "\t")
	w.Write(output)
}

func (c *controller) UpdateBook(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	if param == nil {
	 	return
	}
	id, _ := strconv.Atoi(param["Id"][0])
	year, _ := strconv.Atoi(param["Year"][0])
	for i, book := range books {
		if book.Id == id {
			books[i] = books[len(books)-1]
			books = books[:len(books)-1]
			fmt.Println(books)
			updateBook := Book{Id: id, Title: param["Title"][0], Author: param["Author"][0], Year: year, Genre: param["Genre"][0]}
			books = append(books, updateBook)

			w.Header().Set("Content-Type", "application/json")
			output, _ := json.MarshalIndent(updateBook, "", "\t")
			w.Write(output)
			return
		}
	}
}

func (c *controller) DeleteBook(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	if param == nil {
		return
	}
	id, _ := strconv.Atoi(param["Id"][0])

	for i, book := range books {
		if book.Id == id {
			books[i] = books[len(books)-1]
			books = books[:len(books)-1]
			fmt.Println(books)
			return	
		}
	}
}
