package decrypt

import (
	img "github.com/mehrdad3301/visual-cryptography/internal/pkg/image"
	"github.com/mehrdad3301/visual-cryptography/internal/viscrypt/nofn"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {

	return &cobra.Command{
		Use:   "dec",
		Short: "decrypts cipher",
		Run: func(cmd *cobra.Command, args []string) {
			images := img.ReadImages(args)
			nofn.Decrypt(images)
		},
	}
}
