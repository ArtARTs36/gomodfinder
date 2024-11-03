package gomodfinder

import "fmt"

type FileNotFoundError struct {
	File      string
	Locations []string
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("file %q not found in: %v", e.File, e.Locations)
}
