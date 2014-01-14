package main

import (
	"os"
	"fmt"
)

func main() {
	file, err := os.Open("/etc/passwd")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", file)
}