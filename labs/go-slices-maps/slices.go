package main

import "golang.org/x/tour/pic"

//Pic .
func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for i := range picture {
		picture[i] = make([]uint8, dx)
	}

	return picture
}

func main() {
	pic.Show(Pic)
}
