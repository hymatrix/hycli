package main

import (
	"log"

	"github.com/hymatrix/hycli/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
