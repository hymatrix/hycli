# hycli
**hycli** is the official project scaffolding and management CLI tool for the entire **hymx** Node.

## Build
- Use `make`:
  - `make build`
  - Output binary: `bin/hycli`
- Or use `go` directly:
  - `go build -o bin/hycli ./cmd/hycli`

## Usage
- Basic:
  - `./bin/hycli new`
- Specify output directory:
  - `./bin/hycli new -o ./_sandbox`
- The command prompts for a package name and generates a scaffolded Go project under the specified base directory.

### Command: new
- Description: Create a new Golang project scaffold for hymx Node.
- Flags:
  - `--out`, `-o`: Output base directory. Default: `.`.
- Flow:
  - Run: `./bin/hycli new -o ./_sandbox`
  - Input: package name (e.g., `myproj`)
  - The tool initializes `go.mod` and runs `go mod tidy` inside the generated project.

### Generated Structure
- Base path: `<out>/<pkg>/`
- Contents:
  - `cmd/`
    - `main.go`
    - `flags.go`
    - `const.go`
    - `cmds.go`
    - `cfgchainkit.go`
    - `cfgnode.go`
    - `cfgpay.go`
    - `config.yaml`
    - `config_chainkit.yaml`
    - `config_payment.yaml`
    - `config_test_network.yaml`
    - `mod/*.json` (copied from templates, `.tmpl` suffix removed)
  - `<pkg>/<pkg>.go` (interface file)

### Build Generated Project
- From the generated project root:
  - `cd <out>/<pkg>`
  - `go build -o ./<pkg> ./cmd`

## Notes
- Dependencies used by the scaffold (e.g., `github.com/spf13/viper`, `github.com/urfave/cli/v2`, `github.com/hymatrix/hymx`) are fetched via `go mod tidy` during generation.
