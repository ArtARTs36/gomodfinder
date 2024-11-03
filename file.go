package gomodfinder

import (
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

type ModFile struct {
	*modfile.File

	Path string
}

func (f *ModFile) Dir() string {
	return filepath.Dir(f.Path)
}

func (f *ModFile) CalcPackageFromAbsPath(pkgAbsPath string) *Package {
	parts := strings.Split(pkgAbsPath, string(os.PathSeparator))

	return f.CalcPackageFromAbsPathWithName(parts[len(parts)-1], pkgAbsPath)
}

func (f *ModFile) CalcPackageFromAbsPathWithName(pkgName string, pkgAbsPath string) *Package {
	return &Package{
		Name:               pkgName,
		ModuleRelativePath: strings.TrimLeft(strings.TrimPrefix(pkgAbsPath, f.Dir()), `/\`),
		Module:             f,
	}
}

func (f *ModFile) Package(name string) *Package {
	return &Package{
		Name:               name,
		ModuleRelativePath: name,

		Module: f,
	}
}
