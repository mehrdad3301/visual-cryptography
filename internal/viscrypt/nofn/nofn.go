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
	"strconv"

	img "github.com/mehrdad3301/visual-cryptography/internal/pkg/image"
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

func Encrypt(pic image.Image, n int) {

	startPoint, endPoint := pic.Bounds().Min, pic.Bounds().Max
	rect, c := getRectangle(startPoint, endPoint, n)
	transparencies := getTransparency(n, rect)

	for x := startPoint.X; x < endPoint.X; x++ {
		for y := startPoint.Y; y < endPoint.Y; y++ {
			setPixels(transparencies, x, y, c, img.IsBlack(pic.At(x, y)))
		}
	}

	for i, _ := range transparencies {
		filename := "img_" + strconv.Itoa(i) + ".png"
		img.WriteImage(filename, transparencies[i])
	}
}

func Decrypt(images []image.Image) {

	startPoint, endPoint := images[0].Bounds().Min, images[0].Bounds().Max
	mergedImage := image.NewGray(image.Rect(startPoint.X, startPoint.Y,
		endPoint.X, endPoint.Y))
	for x := startPoint.X; x < endPoint.X; x++ {
		for y := startPoint.Y; y < endPoint.Y; y++ {
			c := color.White
			for _, image := range images {
				if img.IsBlack(image.At(x, y)) {
					c = color.Black
				}
			}
			mergedImage.Set(x, y, c)
		}
	}
	img.WriteImage("merged.png", mergedImage)
}
