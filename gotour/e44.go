package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var fib func (n int) int
    n := 0
    fib = func(n int) int {
        if n == 0 {
            return 0
        } else if n == 1 {
            return 1
        } else {
            return fib(n-1) + fib(n-2)
        }
    }
    return func() int {
        acc := fib(n)
		n++
		return acc        
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 22; i++ {
        fmt.Println(f())
    }
}