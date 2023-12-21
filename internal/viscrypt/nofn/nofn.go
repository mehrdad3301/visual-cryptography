package nofn

/*
nofn package takes an image and argument n. It encrypts
image into n transparencies. The hidden image is only revealed
by stacking all the transparencies. A transparency is one of the
n resulting encrypted images which appears to be nothing more
than randome noise. nofn is short for n out of n scheme.

See examples at assets/example_{2_2 , 3_3 , 4_4}
*/

import (
	"image"
	"image/color"
)

func setPixels(transparencies []*image.Gray, x, y, c int, black bool) {

	n := len(transparencies)
	var shares []int
	if black {
		shares = getBlackShares(n)
	} else {
		shares = getWhiteShares(n)
	}
	setTransparency(transparencies, shares, x, y, c)

}

func Encrypt(pic image.Image, n int) []*image.Gray {

	startPoint, endPoint := pic.Bounds().Min, pic.Bounds().Max
	rect, c := getRectangle(startPoint, endPoint, n)
	transparencies := getTransparency(n, rect)

	for x := startPoint.X; x < endPoint.X; x++ {
		for y := startPoint.Y; y < endPoint.Y; y++ {
			setPixels(transparencies, x, y, c, IsBlack(pic.At(x, y)))
		}
	}

	return transparencies

}

func Decrypt(images []image.Image) *image.Gray {

	startPoint, endPoint := images[0].Bounds().Min, images[0].Bounds().Max
	mergedImage := image.NewGray(image.Rect(startPoint.X, startPoint.Y,
		endPoint.X, endPoint.Y))
	for x := startPoint.X; x < endPoint.X; x++ {
		for y := startPoint.Y; y < endPoint.Y; y++ {
			c := color.White
			for _, image := range images {
				if IsBlack(image.At(x, y)) {
					c = color.Black
				}
			}
			mergedImage.Set(x, y, c)
		}
	}
	return mergedImage
}

func IsBlack(c color.Color) bool {

	r, g, b, _ := c.RGBA()
	y := 0.299*float32(r) +
		0.587*float32(g) +
		0.114*float32(b)

	if y <= 255/2 {
		return true
	} else {
		return false
	}
}
