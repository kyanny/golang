package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    matrix := make([][]uint8, dy)
    for y := 0; y < dy; y++ {
        row := make([]uint8, dx)
        for x := 0; x < dx; x++ {
			row[x] = uint8(x/2+y/2)
        }
        matrix[y] = row
    }
    return matrix
}

func main() {
    pic.Show(Pic)
}