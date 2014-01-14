package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(bytes))
}