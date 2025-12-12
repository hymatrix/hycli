package main

import (
    "log"

    "hycli/internal/cli"
)

func main() {
    if err := cli.Execute(); err != nil {
        log.Fatalf("error: %v", err)
    }
}

