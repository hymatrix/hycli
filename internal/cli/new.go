package cli

import (
	"fmt"
	"path/filepath"

	"github.com/hymatrix/hycli/internal/generator"
	genSchema "github.com/hymatrix/hycli/internal/generator/schema"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new Golang project",
	RunE: func(cmd *cobra.Command, args []string) error {
		// --module / -m
		module, err := cmd.Flags().GetString("module")
		if err != nil {
			return err
		}

		// --out / -o
		outPath, err := cmd.Flags().GetString("out")
		if err != nil {
			return err
		}
		base := outPath
		if base == "" {
			base = "."
		}

		// get project directory, absolute path
		projectDir := filepath.Join(base)
		absProjectDir, err := filepath.Abs(projectDir)
		if err != nil {
			return err
		}
		fmt.Printf("Generating project in: %s\n", projectDir)
		fmt.Println("Project directory:", absProjectDir)
		fmt.Printf("Go Module: %s\n", module)

		// package name
		pkg := filepath.Base(absProjectDir)

		// generate project
		if err := generator.GenerateProject(genSchema.Options{
			Package:    pkg,
			ProjectDir: projectDir,
			GoModule:   module,
		}); err != nil {
			return err
		}

		fmt.Println("Project generation completed")
		return nil
	},
}
