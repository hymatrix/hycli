package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "hycli",
	Short:   "Hymx project scaffolding and management CLI",
	Long:    "Command-line tool to generate Hymx project structure, manage modules, and run the sample project.",
	Version: "v0.0.1",
}

func init() {
	// new
	newCmd.Flags().StringP("out", "o", ".", "Output base directory for the generated project")
	newCmd.Flags().StringP("module", "m", "", "Go module name (e.g. github.com/hymatrix/hycli)")
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
	moduleCmd.Flags().StringP("private-key", "k", "", "Ethereum ECDSA secp256k1 private key hex (0x-prefixed)")
	rootCmd.AddCommand(moduleCmd)
	// run
	rootCmd.AddCommand(runCmd)

	// version
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print version information")
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		show, err := cmd.Flags().GetBool("version")
		if err != nil {
			return err
		}
		if show {
			fmt.Println(rootCmd.Version)
			os.Exit(0)
		}
		return nil
	}
}

func Execute() error { return rootCmd.Execute() }
