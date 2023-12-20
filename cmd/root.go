package cmd

import (
	"image"
	"log"
	"os"

	"github.com/mehrdad3301/visual-cryptography/cmd/server"
	img "github.com/mehrdad3301/visual-cryptography/internal/pkg/image"
	"github.com/mehrdad3301/visual-cryptography/internal/viscrypt/nofn"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"
)

const ExitFailure = -1

func Execute() {

	if err := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("Vis", pterm.NewStyle(pterm.FgCyan)),
		putils.LettersFromStringWithStyle("Crypt", pterm.NewStyle(pterm.FgLightMagenta)),
	).Render(); err != nil {
		_ = err
	}

	//nolint: exhaustruct
	root := &cobra.Command{
		Use:   "viscrypt",
		Short: "visual cryptography",
	}

	{

		var numTransparencies int

		cmd := &cobra.Command{
			Use:   "enc",
			Short: "encrypts image",
			Run: func(cmd *cobra.Command, args []string) {
				img, err := img.ReadImage(args[0])
				if err != nil {
					log.Fatal(err)
				}
				nofn.Encrypt(img, numTransparencies)
			},
		}

		cmd.Flags().IntVarP(&numTransparencies, "transparencies", "t", 2,
			"number of transparencies used for encryption")

		root.AddCommand(cmd)

	}

	{

		cmd := &cobra.Command{
			Use:   "dec",
			Short: "decrypts image",
			Run: func(cmd *cobra.Command, args []string) {
				images := make([]image.Image, 0)
				for _, name := range args[0:] {
					img, err := img.ReadImage(name)
					if err != nil {
						log.Fatal("couldn't read image %w", err)
					}
					images = append(images, img)
				}
				nofn.Decrypt(images)
			},
		}

		root.AddCommand(cmd)
	}

	root.AddCommand(server.New())

	if err := root.Execute(); err != nil {
		pterm.Error.Println(err.Error())
		os.Exit(ExitFailure)
	}
}

