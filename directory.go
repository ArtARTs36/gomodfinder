package gomodfinder

import (
	"errors"
	"fmt"
	"os"

	"path/filepath"
)

var ErrFileNotFound = fmt.Errorf("file not found")

type Directory interface {
	ReadFile(path string) ([]byte, error)
	Parent() (Directory, error)
	Path() string
	PathTo(filename string) string
}

type OsDirectory struct {
	dir string
}

func NewOsDirectory(dir string) *OsDirectory {
	return &OsDirectory{dir: dir}
}

func (d *OsDirectory) ReadFile(path string) ([]byte, error) {
	content, err := os.ReadFile(d.PathTo(path))
	if errors.Is(err, os.ErrNotExist) {
		return nil, ErrFileNotFound
	}
	return content, err
}

func (d *OsDirectory) Parent() (Directory, error) {
	currDir := d.dir
	if currDir[len(currDir)-1] == os.PathSeparator {
		currDir = currDir[0 : len(currDir)-2]
	}

	parent := filepath.Dir(currDir)

	return NewOsDirectory(parent), nil
}

func (d *OsDirectory) Path() string {
	return d.dir
}

func (d *OsDirectory) PathTo(filename string) string {
	return fmt.Sprintf("%s%s%s", d.dir, string(os.PathSeparator), filename)
}
