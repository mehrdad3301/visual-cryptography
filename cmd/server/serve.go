package server

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mehrdad3301/visual-cryptography/internal/viscrypt/nofn"
	"github.com/spf13/cobra"
)

const (
	imageDir = "./images"
)

func New() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "starts the web server",
		Run: func(cmd *cobra.Command, args []string) {
			app := echo.New()
			app.HideBanner = true
			app.Use(middleware.CORS())

			if _, err := os.Stat(imageDir); os.IsNotExist(err) {
				if err = os.Mkdir(imageDir, os.ModePerm); err != nil {
					log.Print(fmt.Errorf("couldn't create directory: %w", err))
					log.Fatal(err)
				}
			}

			app.POST("/upload", uploadImage)
			app.Static("/images", imageDir)
			log.Fatal(app.Start(":1234"))
		},
	}

	return cmd
}

func uploadImage(c echo.Context) error {

	file, err := c.FormFile("image")
	if err != nil {
		log.Print(err)
		return echo.ErrInternalServerError
	}

	k, err := strconv.Atoi(c.QueryParam("k"))
	if err != nil {
		log.Print(err)
		return echo.ErrInternalServerError
	}

	src, err := file.Open()
	if err != nil {
		log.Print(err)
		return echo.ErrInternalServerError
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		log.Print(err)
		return echo.ErrInternalServerError
	}

	encImages := nofn.Encrypt(img, k)

	requestID := uuid.New().String()
	requestDir := filepath.Join(imageDir, requestID)
	if err := os.Mkdir(requestDir, os.ModePerm); err != nil {
		log.Print(fmt.Errorf("couldn't create directory: %w", err))
		return echo.ErrInternalServerError
	}

	imagePaths := []string{}
	for i, encryptedImage := range encImages {
		imagePath := filepath.Join(requestDir, fmt.Sprintf("image_%d.png", i+1))
		imageFile, err := os.Create(imagePath)
		if err != nil {
			log.Print(fmt.Errorf("couldn't create file: %w", err))
			return echo.ErrInternalServerError
		}
		defer imageFile.Close()

		err = png.Encode(imageFile, encryptedImage)
		if err != nil {
			log.Print(fmt.Errorf("image couldn't be encoded in jpeg: %w", err))
			return echo.ErrInternalServerError
		}
		imagePaths = append(imagePaths, imagePath)
	}

	return c.JSON(http.StatusOK,
		map[string]interface{}{"requestID": requestID,
			"imagePaths": imagePaths})
}
