package main

import (
	"fmt"
)

var stack [10]int;
var i int;

func push(elem int) {
	i++;
	stack[i] = elem;
}

func pop() {
	i--;
	return stack[i];
}

func main() {
	stack = [1,2,3,4,5,6,7,8,9,10];
	fmt.Println(pop());
	fmt.Println(stack);
	push(999);
	fmt.Println(stack);
}