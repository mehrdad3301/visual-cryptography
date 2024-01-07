package nofn

import (
	"image"
	"image/color"
)

func getTransparency(n int, rect image.Rectangle) []*image.RGBA {

	transparencies := make([]*image.RGBA, 0, n)
	for i := 0; i < n; i++ {
		transparencies = append(transparencies, image.NewRGBA(rect))
	}
	return transparencies

}

func setShare(transparency *image.RGBA, share int, x, y, c int) {

	for i := 0; i < c; i++ {
		for j := 0; j < c; j++ {

			bit := getKthBit(share, i*c+j)
			clr := color.RGBA{0, 0, 0, 0}
			if bit == 1 {
				clr = color.RGBA{0, 0, 0, 255}
			}
			transparency.Set(c*x+i, c*y+j, clr)
		}
	}
}

func setTransparency(transparencies []*image.RGBA, shares []int, x, y, c int) {

	for i := range transparencies {
		setShare(transparencies[i], shares[i], x, y, c)
	}
}
