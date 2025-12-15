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
  - `./hycli new`
  - `./hycli vmm --name <vmm> --format <format>`
  - `./hycli mount --name <vmm>`
  - `./hycli module --name <module> [-u <nodeURL>] [-k <privateKey>]`
  - `./hycli run`
- Specify output directory:
  - `./hycli new -o ./_sandbox`
- The command prompts for a package name and generates a scaffolded Go project under the specified base directory.

### Command: new
- Description: Create a new Golang project scaffold for hymx Node.
- Flags:
  - `--out`, `-o`: Output base directory. Default: `.`.
- Flow:
  - Run: `./bin/hycli new -o ./_sandbox`
  - Input: package name (e.g., `myproj`)
  - The tool initializes `go.mod` and runs `go mod tidy` inside the generated project.

### Command: vmm
- Description: Manage or scaffold a VM module, and auto-mount it into `cmd/main.go`.
- Flags:
  - `--name`, `-n`: Name of the vmm
  - `--format`, `-f`: Module format of the vmm
- Notes:
  - Automatically inserts imports (`<vmm>`, `<vmm>Schema`) and adds `s.Mount(<vmm>Schema.ModuleFormat, <vmm>.Spawn)` into `cmd/main.go`.
  - Run inside the generated project root (the directory containing `cmd/main.go`).

### Command: mount
- Description: Mount an existing VM module into `cmd/main.go`.
- Flags:
  - `--name`, `-n`: Name of the vmm
- Notes:
  - Reads the module format from `<projectDir>/<vmm>/schema/schema.go`.
  - Inserts imports and the `s.Mount(...)` call under the mount hint comments in `cmd/main.go`.

### Command: module
- Description: Generate and mount a module by name, reading module format from schema.
- Flags:
  - `--name`, `-n`: Name of the module
  - `--node-url`, `-u`: Node URL. Default: `http://127.0.0.1:8080`
  - `--private-key`, `-k`: Private key
- Notes:
  - Reads `ModuleFormat` from `<projectDir>/<name>/schema/schema.go`.
  - Passes `node-url` and `private-key` to the generation pipeline.
  - Mounts the module into `cmd/main.go` using the same insertion logic as `vmm`.
  - When saving a module via SDK, generated `mod-<itemId>.json` will be placed under `cmd/mod/`.

### Command: run
- Description: Run the generated project.
- Flow:
  - Executes: `cd cmd && go run ./`
  - From the generated project root, runs the `cmd/main.go` entrypoint.

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
    - `mod/mod-<itemId>.json` (when SDK-based module save is performed)
  - `<pkg>/<pkg>.go` (interface file)

### Build Generated Project
- From the generated project root:
  - `cd <out>/<pkg>`
  - `go build -o ./<pkg> ./cmd`

## Notes
- Dependencies used by the scaffold (e.g., `github.com/spf13/viper`, `github.com/urfave/cli/v2`, `github.com/hymatrix/hymx`) are fetched via `go mod tidy` during generation.

## TODO
- [x] Create project directories
- [x] Initialize Golang environment
- [x] Create Vmm & Mount
- [x] Generate Module
- [x] Run
- [ ] Setup Redis environment

