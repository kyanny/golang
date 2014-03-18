package main

import "fmt"

func main() {
	var a [4]int
	fmt.Println(a)
	b := [4]int{1,2,3,4}
	fmt.Println(b)
	c := [...]int{5,6,7,8}
	fmt.Println(c)
}