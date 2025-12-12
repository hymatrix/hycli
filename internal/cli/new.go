package cli

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"hycli/internal/generator"
	genSchema "hycli/internal/generator/schema"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new Golang project",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the package name to create: ")
		pkg, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		pkg = strings.TrimSpace(pkg)
		if pkg == "" {
			return fmt.Errorf("package name cannot be empty")
		}

		base := outPath
		if base == "" {
			base = "."
		}
		projectDir := filepath.Join(base)

		fmt.Printf("Generating project in: %s\n", projectDir)
		if err := generator.GenerateProject(genSchema.Options{
			Package:   pkg,
			OutputDir: projectDir,
		}); err != nil {
			return err
		}

		fmt.Println("Project generation completed")
		return nil
	},
}
