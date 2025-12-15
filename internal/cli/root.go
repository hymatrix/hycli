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
	// new
	rootCmd.AddCommand(newCmd)
	// vmm
	vmmCmd.Flags().StringP("name", "n", "", "Name of the vmm")
	vmmCmd.Flags().StringP("format", "f", "", "Module format of the vmm")
	rootCmd.AddCommand(vmmCmd)
	// mount
	mountCmd.Flags().StringP("name", "n", "", "Name of the vmm")
	rootCmd.AddCommand(mountCmd)
	// module
	moduleCmd.Flags().StringP("name", "n", "", "Name of the module")
	moduleCmd.Flags().StringP("node-url", "u", "http://127.0.0.1:8080", "Node URL")
	moduleCmd.Flags().StringP("private-key", "k", "", "Private key")
	rootCmd.AddCommand(moduleCmd)
	// run
	rootCmd.AddCommand(runCmd)
}

func Execute() error { return rootCmd.Execute() }
