package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("/etc/passwd")
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(bytes))
}