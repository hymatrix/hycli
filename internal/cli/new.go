package cli

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hymatrix/hycli/internal/generator"
	genSchema "github.com/hymatrix/hycli/internal/generator/schema"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new Golang project",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the organization name: ")
		org, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		org = strings.TrimSpace(org)
		if org == "" {
			return fmt.Errorf("organization name cannot be empty")
		}

		base := outPath
		if base == "" {
			base = "."
		}
		projectDir := filepath.Join(base)
		absProjectDir, err := filepath.Abs(projectDir)
		if err != nil {
			return err
		}
		fmt.Println("Project directory:", absProjectDir)
		pkg := filepath.Base(absProjectDir)

		fmt.Printf("Generating project in: %s\n", projectDir)
		if err := generator.GenerateProject(genSchema.Options{
			Org:        org,
			Package:    pkg,
			ProjectDir: projectDir,
		}); err != nil {
			return err
		}

		fmt.Println("Project generation completed")
		return nil
	},
}
