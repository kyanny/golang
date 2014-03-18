package main

import (
    "fmt"
    "math"
)

func Sqrt1(x float64) float64 {
    z := 1.0
    for i:=0; i<10; i++ {
        z = z - (z*z-x)/(2*z)
    }
    fmt.Printf("loop count: 10\n")
    return z
}

func Sqrt2(x float64) float64 {
    z, y := 1.0, 1.0
    z = Newton(x, z)

    i := 0
    for ; math.Abs(z-y) >= 0.000000000000001 ; i++ {
        y = z
        z = Newton(x, z)
    }
    
    fmt.Printf("loop count: %d\n", i)
    return z
}

func Newton(x, z float64) float64 {
    z = z - (z*z-x) / (2*z)
    return z
}

func main() {
    fmt.Println(Sqrt1(2))
    fmt.Println(Sqrt2(2))
    fmt.Println(Sqrt1(3))
    fmt.Println(Sqrt2(3))
    fmt.Println(Sqrt1(5))
    fmt.Println(Sqrt2(5))
}
