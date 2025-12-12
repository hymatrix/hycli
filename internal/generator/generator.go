package generator

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"hycli/internal/generator/schema"
	"hycli/internal/templates"
)

func GenerateProject(opts schema.Options) error {
	pkg := opts.Package
	projectDir := opts.OutputDir
	if err := os.MkdirAll(projectDir, 0o755); err != nil {
		return err
	}
	dirs := []string{
		filepath.Join(projectDir, pkg, "cmd"),
		filepath.Join(projectDir, pkg, pkg),
	}
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0o755); err != nil {
			return err
		}
	}

	data := schema.Options{Package: pkg}

	// cmd templates
	if err := renderTemplateFile("cmd/main.go.tmpl", filepath.Join(projectDir, pkg, "cmd", "main.go"), data); err != nil {
		return err
	}
	if err := renderTemplateFile("cmd/flags.go.tmpl", filepath.Join(projectDir, pkg, "cmd", "flags.go"), data); err != nil {
		return err
	}
	if err := renderTemplateFile("cmd/const.go.tmpl", filepath.Join(projectDir, pkg, "cmd", "const.go"), data); err != nil {
		return err
	}
	if err := renderTemplateFile("cmd/cmds.go.tmpl", filepath.Join(projectDir, pkg, "cmd", "cmds.go"), data); err != nil {
		return err
	}
	if err := renderTemplateFile("cmd/cfgchainkit.go.tmpl", filepath.Join(projectDir, pkg, "cmd", "cfgchainkit.go"), data); err != nil {
		return err
	}
	if err := renderTemplateFile("cmd/cfgpay.go.tmpl", filepath.Join(projectDir, pkg, "cmd", "cfgpay.go"), data); err != nil {
		return err
	}
	if err := renderTemplateFile("cmd/cfgnode.go.tmpl", filepath.Join(projectDir, pkg, "cmd", "cfgnode.go"), data); err != nil {
		return err
	}

	if err := renderTemplateFile("cmd/interface.go.tmpl", filepath.Join(projectDir, pkg, pkg, pkg+".go"), data); err != nil {
		return err
	}

	// configs (static + templated variants)
	if err := writeRawFile("cmd/config.yaml", filepath.Join(projectDir, pkg, "cmd", "config.yaml")); err != nil {
		return err
	}
	if err := writeRawFile("cmd/config_chainkit.yaml.tmpl", filepath.Join(projectDir, pkg, "cmd", "config_chainkit.yaml")); err != nil {
		return err
	}
	if err := writeRawFile("cmd/config_payment.yaml.tmpl", filepath.Join(projectDir, pkg, "cmd", "config_payment.yaml")); err != nil {
		return err
	}
	if err := writeRawFile("cmd/config_test_network.yaml.tmpl", filepath.Join(projectDir, pkg, "cmd", "config_test_network.yaml")); err != nil {
		return err
	}

	// mod: copy all files under templates/cmd/mod into generated cmd/mod, rename *.tmpl -> remove suffix
	{
		outModDir := filepath.Join(projectDir, pkg, "cmd", "mod")
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
			name := e.Name() // e.g., mod-xxx.json.tmpl
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
	}

	// go tidy
	if err := runGoInitAndTidy(filepath.Join(projectDir, pkg), pkg); err != nil {
		return err
	}

	return nil
}

func renderTemplateFile(tmplName string, outPath string, data any) error {
	b, err := templates.FS.ReadFile(tmplName)
	if err != nil {
		return err
	}
	funcs := template.FuncMap{
		"capitalize": func(s string) string {
			if s == "" {
				return s
			}
			r := []rune(s)
			r[0] = unicode.ToUpper(r[0])
			return string(r)
		},
	}
	t, err := template.New(tmplName).Funcs(funcs).Parse(string(b))
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return err
	}
	if err := os.WriteFile(outPath, buf.Bytes(), 0o644); err != nil {
		return err
	}
	return nil
}

func writeRawFile(srcInFS string, outPath string) error {
	b, err := templates.FS.ReadFile(srcInFS)
	if err != nil {
		return err
	}
	if err := os.WriteFile(outPath, b, 0o644); err != nil {
		return err
	}
	return nil
}

func runGoInitAndTidy(projectDir string, module string) error {
	if err := runCmd(projectDir, "go", "mod", "init", module); err != nil {
		return err
	}
	if err := runCmd(projectDir, "go", "mod", "tidy"); err != nil {
		return err
	}
	return nil
}

func runCmd(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s %v failed: %w", name, args, err)
	}
	return nil
}
