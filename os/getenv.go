package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Println("USER=" + os.Getenv("USER"))
	fmt.Println("User=" + os.Getenv("User"))
	fmt.Println("user=" + os.Getenv("user"))
}