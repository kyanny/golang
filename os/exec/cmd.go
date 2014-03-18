package main

import (
	"os/exec"
	"fmt"
)

func main() {
	out, err := exec.Command("uname", "-a").Output()
	if err != nil {
		panic(err)
	}
	fmt.Print(string(out));
}