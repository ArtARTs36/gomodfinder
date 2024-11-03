package gomodfinder

import (
	"errors"
	"fmt"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

func Find(dir string, levels int) (*ModFile, error) {
	return FindIn(NewOsDirectory(dir), levels)
}

func FindIn(dir Directory, levels int) (*ModFile, error) {
	goModContent := []byte{}
	goModFound := false

	currDir := dir
	scanned := []string{}

	for i := 0; i < levels; i++ {
		var err error

		currDir, err = currDir.Parent()
		if err != nil {
			return nil, fmt.Errorf("failed to get parent directory: %w", err)
		}

		goModContent, err = currDir.ReadFile("go.mod")
		if err != nil && errors.Is(err, ErrFileNotFound) {
			scanned = append(scanned, currDir.Path())

			continue
		}

		goModFound = true
	}

	if !goModFound {
		return nil, fmt.Errorf("go mod file not found in: %v", scanned)
	}

	path, err := filepath.Abs(currDir.PathTo("go.mod"))
	if err != nil {
		return nil, fmt.Errorf("failed to get absoulte path: %w", err)
	}

	stdModFile, err := modfile.ParseLax(currDir.PathTo("go.mod"), goModContent, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to parse go mod file: %w", err)
	}

	return &ModFile{
		File: stdModFile,
		Path: path,
	}, nil
}
