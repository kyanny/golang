package main

import (
	"fmt"
	"html"
)

func main() {
	fmt.Println(html.EscapeString("<br />"))
	fmt.Println(html.UnescapeString("&lt;/span&gt;"))
}