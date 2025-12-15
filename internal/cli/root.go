package cli

import (
	"github.com/spf13/cobra"
)

var (
	outPath string
)

var rootCmd = &cobra.Command{
	Use:   "hycli",
	Short: "Generate a Golang project scaffold",
	Long:  "A CLI to generate a standard Golang project directory and bootstrap files.",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outPath, "out", "o", ".", "Output base directory for the generated project")

	rootCmd.AddCommand(newCmd)
	vmmCmd.Flags().StringP("name", "n", "", "Name of the vmm")
	vmmCmd.Flags().StringP("format", "f", "", "Module format of the vmm")
	rootCmd.AddCommand(vmmCmd)
}

func Execute() error { return rootCmd.Execute() }
