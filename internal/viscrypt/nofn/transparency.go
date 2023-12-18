package nofn

import (
	"image"
	"image/color"
)

func getTransparency(n int, rect image.Rectangle) []*image.Gray {

	transparencies := make([]*image.Gray, 0, n)
	for i := 0; i < n; i++ {
		transparencies = append(transparencies, image.NewGray(rect))
	}
	return transparencies

}

func setShare(transparency *image.Gray, share int, x, y, c int) {

	for i := 0; i < c; i++ {
		for j := 0; j < c; j++ {

			bit := getKthBit(share, i*c+j)
			clr := color.White
			if bit == 1 {
				clr = color.Black
			}
			transparency.Set(c*x+i, c*y+j, clr)
		}
	}
}

func setTransparency(transparencies []*image.Gray, shares []int, x, y, c int) {

	for i := range transparencies {
		setShare(transparencies[i], shares[i], x, y, c)
	}
}
