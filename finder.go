package gomodfinder

import (
	"errors"
	"fmt"

	"golang.org/x/mod/modfile"
)

func Find(dir string, levels int) (*modfile.File, error) {
	return FindIn(NewOsDirectory(dir), levels)
}

func FindIn(dir Directory, levels int) (*modfile.File, error) {
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

	return modfile.ParseLax(currDir.PathTo("go.mod"), goModContent, nil)
}
