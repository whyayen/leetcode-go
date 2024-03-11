package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}

	originalColor := image[sr][sc]
	pixels := [][]int{{sr, sc}}

	for len(pixels) > 0 {
		currentR, currentC := pixels[0][0], pixels[0][1]
		pixels = pixels[1:]
		image[currentR][currentC] = color

		tx, ty := max(currentR-1, 0), currentC
		lx, ly := currentR, max(currentC-1, 0)
		bx, by := min(currentR+1, len(image)-1), currentC
		rx, ry := currentR, min(currentC+1, len(image[currentR])-1)

		if image[tx][ty] == originalColor {
			pixels = append(pixels, []int{tx, ty})
		}

		if image[lx][ly] == originalColor {
			pixels = append(pixels, []int{lx, ly})
		}

		if image[bx][by] == originalColor {
			pixels = append(pixels, []int{bx, by})
		}

		if image[rx][ry] == originalColor {
			pixels = append(pixels, []int{rx, ry})
		}
	}

	return image
}

func main() {
	i := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	result := floodFill(i, 1, 1, 2)
	fmt.Printf("%+v \n", result)
}
