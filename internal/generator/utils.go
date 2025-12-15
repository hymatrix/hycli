package generator

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"text/template"
	"unicode"

	"github.com/hymatrix/hycli/internal/templates"
)

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

func readLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var lines []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
