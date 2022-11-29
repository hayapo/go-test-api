package api

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
}

func init() {
	books = append(books, Book{Id: 1, Title: "吾輩は猫である", Author: "夏目漱石", Year: 1906, Genre: "novel"})
	books = append(books, Book{Id: 2, Title: "走れメロス", Author: "太宰治", Year: 1940, Genre: "novel"})
	books = append(books, Book{Id: 3, Title: "こころ", Author: "夏目漱石", Year: 1914, Genre: "novel"})
}
