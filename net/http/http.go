package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Top)
	fmt.Println("http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func Top(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

// $ curl -i localhost:8080