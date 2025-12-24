package generator

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/hymatrix/hycli/internal/generator/schema"
)

type ModMeta struct {
	Dir string
}

func Get(opts schema.Options) error {
	projectDir := opts.ProjectDir
	pkg := opts.Package
	if projectDir == "" || pkg == "" {
		return fmt.Errorf("invalid get options")
	}
	basePath := strings.SplitN(pkg, "@", 2)[0]
	name := path.Base(basePath)

	if err := runCmd(projectDir, "go", "get", pkg); err != nil {
		return err
	}

	meta, err := ListModule(projectDir, basePath)
	if err != nil {
		return err
	}
	if meta.Dir == "" {
		return fmt.Errorf("module directory not found for %s", pkg)
	}

	outDir := filepath.Join(projectDir, name)
	if err := copyDir(meta.Dir, outDir); err != nil {
		return err
	}

	return Mount(schema.Options{
		ProjectDir: projectDir,
		VmmName:    name,
	})
}

func ListModule(dir string, pkg string) (ModMeta, error) {
	basePath := strings.SplitN(pkg, "@", 2)[0]
	c := exec.Command("go", "list", "-json", basePath)
	c.Dir = dir
	out, err := c.Output()
	if err != nil {
		var stderr []byte
		if exitErr, ok := err.(*exec.ExitError); ok {
			stderr = exitErr.Stderr
		}
		return ModMeta{}, fmt.Errorf("go list failed: %w, stderr: %s", err, string(stderr))
	}
	var m ModMeta
	if err := json.Unmarshal(out, &m); err != nil {
		return ModMeta{}, fmt.Errorf("failed to unmarshal go list output: %w", err)
	}
	return m, nil
}
