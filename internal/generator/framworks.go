package generator

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/hymatrix/hycli/internal/generator/schema"
	"github.com/hymatrix/hycli/internal/templates"
)

func genFrameworks(opts schema.Options) error {
	pkg := opts.Package
	projectDir := opts.ProjectDir

	dirs := []string{
		filepath.Join(projectDir, "cmd"),
		filepath.Join(projectDir, pkg),
	}
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0o755); err != nil {
			return err
		}
	}

	data := schema.Options{Package: pkg}

	if err := renderTemplateFile("cmd/main.go.tmpl", filepath.Join(projectDir, "cmd", "main.go"), data); err != nil {
		return err
	}
	if err := renderTemplateFile("cmd/flags.go.tmpl", filepath.Join(projectDir, "cmd", "flags.go"), data); err != nil {
		return err
	}
	if err := renderTemplateFile("cmd/const.go.tmpl", filepath.Join(projectDir, "cmd", "const.go"), data); err != nil {
		return err
	}
	if err := renderTemplateFile("cmd/cmds.go.tmpl", filepath.Join(projectDir, "cmd", "cmds.go"), data); err != nil {
		return err
	}
	if err := renderTemplateFile("cmd/cfgchainkit.go.tmpl", filepath.Join(projectDir, "cmd", "cfgchainkit.go"), data); err != nil {
		return err
	}
	if err := renderTemplateFile("cmd/cfgpay.go.tmpl", filepath.Join(projectDir, "cmd", "cfgpay.go"), data); err != nil {
		return err
	}
	if err := renderTemplateFile("cmd/cfgnode.go.tmpl", filepath.Join(projectDir, "cmd", "cfgnode.go"), data); err != nil {
		return err
	}

	if err := renderTemplateFile("cmd/interface.go.tmpl", filepath.Join(projectDir, pkg, pkg+".go"), data); err != nil {
		return err
	}

	if err := writeRawFile("cmd/config.yaml", filepath.Join(projectDir, "cmd", "config.yaml")); err != nil {
		return err
	}
	if err := writeRawFile("cmd/config_chainkit.yaml.tmpl", filepath.Join(projectDir, "cmd", "config_chainkit.yaml")); err != nil {
		return err
	}
	if err := writeRawFile("cmd/config_payment.yaml.tmpl", filepath.Join(projectDir, "cmd", "config_payment.yaml")); err != nil {
		return err
	}
	if err := writeRawFile("cmd/config_test_network.yaml.tmpl", filepath.Join(projectDir, "cmd", "config_test_network.yaml")); err != nil {
		return err
	}

	outModDir := filepath.Join(projectDir, "cmd", "mod")
	if err := os.MkdirAll(outModDir, 0o755); err != nil {
		return err
	}
	entries, err := templates.FS.ReadDir("cmd/mod")
	if err != nil {
		return err
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		src := filepath.Join("cmd", "mod", name)
		outName := name
		if strings.HasSuffix(outName, ".tmpl") {
			outName = strings.TrimSuffix(outName, ".tmpl")
		}
		dst := filepath.Join(outModDir, outName)
		if err := writeRawFile(src, dst); err != nil {
			return err
		}
	}

	return nil
}
