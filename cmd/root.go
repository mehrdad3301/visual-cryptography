package cmd

import (
	"os"

	"github.com/mehrdad3301/visual-cryptography/cmd/cli"
	"github.com/mehrdad3301/visual-cryptography/cmd/server"
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
				cli.Enc(cmd, args, numTransparencies)
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
				cli.Dec(cmd, args)
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
