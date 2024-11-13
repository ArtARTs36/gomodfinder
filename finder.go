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

		goModContent, err = currDir.ReadFile("go.mod")
		if err != nil {
			nfErr := &FileNotFoundError{}
			if errors.As(err, &nfErr) {
				scanned = append(scanned, currDir.Path())

				currDir, err = currDir.Parent()
				if err != nil {
					return nil, fmt.Errorf("failed to get parent directory: %w", err)
				}

				continue
			}

			return nil, fmt.Errorf("failed to read file: %w", err)
		}

		goModFound = true
	}

	if !goModFound {
		return nil, &FileNotFoundError{
			File:      "go.mod",
			Locations: scanned,
		}
	}

	path, err := filepath.Abs(currDir.PathTo("go.mod"))
	if err != nil {
		return nil, fmt.Errorf("failed to get absoulte path: %w", err)
	}

	stdModFile, err := modfile.ParseLax(path, goModContent, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to parse go mod file: %w", err)
	}

	return &ModFile{
		File: stdModFile,
		Path: path,
	}, nil
}
