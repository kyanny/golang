package main

import (
	"fmt"
	"encoding/base64"
)

var b64a = "QUJDREVGRw=="
var b64b = "c29tZSBkYXRhIHdpdGggACBhbmQg77u/"

func main() {
	b, err := base64.StdEncoding.DecodeString(b64a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", string(b))

	b, err = base64.URLEncoding.DecodeString(b64a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", string(b))

	b, err = base64.StdEncoding.DecodeString(b64b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", string(b))

	b, err = base64.URLEncoding.DecodeString(b64b)
	if err != nil {
		panic(err) // panic: illegal base64 data at input byte 31 ("/" character)
	}
	fmt.Printf("%q\n", string(b))
}