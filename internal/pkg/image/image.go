package image

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
)

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

func ReadImage(filename string) image.Image {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

func WriteImage(filename string, img *image.Gray) {

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	png.Encode(f, img)
}

func WriteImages(imgs []*image.Gray) {

	for i, img := range imgs {
		filename := "img_" + strconv.Itoa(i) + ".png"
		WriteImage(filename, img)
	}
}

func ReadImages(names []string) []image.Image {

	images := make([]image.Image, 0)
	for _, name := range names {
		images = append(images, ReadImage(name))
	}
	return images
}
