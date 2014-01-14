package main

import (
	"fmt"
	"encoding/json"
)

type Book struct {
	Title string
	Issue uint
	Authors []string
	Publisher Publisher
}

type Publisher struct {
	Name string
}

func main() {
	data := []byte(`
  {
    "title": "Structure and Interpretation of Computer Programs",
    "issue": 1984,
    "authors": ["Harold Abelson", "Gerald Jay Sussman", "Julie Sussman"],
    "publisher": {
      "name": "The MIT Press"
    }
  }
`)
	book := new(Book)
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", book)
}
