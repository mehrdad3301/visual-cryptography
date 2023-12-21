package cli

import (
	"image"
	"image/png"
	"log"
	"os"
	"strconv"

	"github.com/mehrdad3301/visual-cryptography/internal/viscrypt/nofn"
	"github.com/spf13/cobra"
)

const (
	namePrefix      = "img_"
	nameSuffix      = ".png"
	mergedImageName = "merged" + nameSuffix
)

func Enc(cmd *cobra.Command, args []string, k int) {

	img, err := decodeImage(args[0])
	if err != nil {
		log.Fatal(err)
	}

	imgs := nofn.Encrypt(img, k)

	for i := range imgs {
		filename := namePrefix + strconv.Itoa(i) + nameSuffix
		err = writeImage(filename, imgs[i])
		if err != nil {
			log.Fatal(err)
		}

	}

}

func Dec(cmd *cobra.Command, args []string) {

	images := make([]image.Image, 0)
	for _, name := range args[0:] {
		img, err := decodeImage(name)
		if err != nil {
			log.Fatal(err)
		}
		images = append(images, img)
	}
	img := nofn.Decrypt(images)
	writeImage(mergedImageName, img)
}

func decodeImage(filename string) (image.Image, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func writeImage(filename string, img *image.Gray) error {

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
