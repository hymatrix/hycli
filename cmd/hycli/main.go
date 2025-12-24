package main

import (
	"log"

	"github.com/hymatrix/hycli/internal/cli"
	"github.com/hymatrix/hycli/internal/generator"
	"github.com/hymatrix/hycli/internal/generator/schema"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func mount(projectDir string, vmmName string) error {
	return generator.Mount(schema.Options{
		ProjectDir: projectDir,
		VmmName:    vmmName,
	})
}
