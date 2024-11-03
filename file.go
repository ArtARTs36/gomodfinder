package gomodfinder

import (
	"path/filepath"

	"golang.org/x/mod/modfile"
)

type ModFile struct {
	*modfile.File

	Path string
}

func (f *ModFile) Dir() string {
	return filepath.Dir(f.Path)
}
