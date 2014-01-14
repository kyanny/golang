package main

import (
	"fmt"
	"net/url"
)

func main() {
	url, err := url.Parse("http://example.com/foo?bar=baz")
	if err != nil {
		panic(err)
	}
	fmt.Println("path: " + url.Path)
	fmt.Println("params: " + url.RawQuery)
	fmt.Println("params[\"bar\"]: " + url.Query().Get("bar"))
}