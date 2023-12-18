package encrypt

import (
	"github.com/mehrdad3301/visual-cryptography/internal/viscrypt/nofn"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {

	var numTransparencies int

	cmd := &cobra.Command{
		Use:   "enc",
		Short: "encrypts image",
		Run: func(cmd *cobra.Command, args []string) {
			nofn.Encrypt(args[0], numTransparencies)
		},
	}

	cmd.Flags().IntVarP(&numTransparencies, "transparencies", "t", 2,
		"number of transparencies used for encryption")

	return cmd
}
