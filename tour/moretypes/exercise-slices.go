package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for y := range p {
		p[y] = make([]uint8, dx)
	}
	for y, row := range p {
		for x := range row {
			row[x] = uint8(x ^ y)
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}
