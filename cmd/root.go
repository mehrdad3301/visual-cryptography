package cmd

import (
	"os"

	"github.com/mehrdad3301/visual-cryptography/cmd/decrypt"
	"github.com/mehrdad3301/visual-cryptography/cmd/encrypt"
	"github.com/spf13/cobra"
)

const ExitFailure = -1

func Execute() {

	//nolint: exhaustruct
	root := &cobra.Command{
		Use:   "viscrypt",
		Short: "visual cryptography",
	}

	root.AddCommand(encrypt.New())
	root.AddCommand(decrypt.New())

	if err := root.Execute(); err != nil {
		os.Exit(ExitFailure)
	}
}
