package main

import (
	"fmt"
	"encoding/json"
)

type Book struct {
	Title string
	Issue uint
	Authors []string
}

func main() {
	book := new(Book)
	book.Title = "Structure and Interpretation of Computer Programs"
	book.Issue = 1984
	authors := []string{"Harold Abelson", "Gerald Jay Sussman", "Julie Sussman"}
	book.Authors = authors
	data, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
