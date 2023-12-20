package server

import (
	"fmt"
	"log"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "starts the web server",
		Run: func(cmd *cobra.Command, args []string) {
			app := echo.New()
			app.HideBanner = true

			app.POST("/upload", uploadImage)
			log.Fatal(app.Start(":1234"))
		},
	}

	return cmd
}

func uploadImage(c echo.Context) error {

	image, err := c.FormFile("image")
	if err != nil {
		return err
	}

	k := c.QueryParam("k")

	src, err := image.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	return c.String(http.StatusOK, fmt.Sprintf("File uploaded successfully: %s, k=%s", image.Filename, k))
}
