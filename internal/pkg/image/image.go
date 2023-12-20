package image

import (
	"image"
	"image/color"
	"image/png"
	"os"
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

func ReadImage(filename string) (image.Image, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func WriteImage(filename string, img *image.Gray) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}
	return nil
}
