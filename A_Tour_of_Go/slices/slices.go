package main

import pic "golang.org/x/tour/pic"

// Pic ... Output image as 2d array (uint8)
func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for i := range image {
		image[i] = make([]uint8, dx)
	}

	for x := range image {
		for y := range image[x] {
			image[y][x] = uint8((x + y) / 2)
		}
	}

	return image
}

func main() {
	pic.Show(Pic)
}
