package gomodfinder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModFile_CalcPackageFromAbsPath(t *testing.T) {
	mod := &ModFile{
		Path: "/Users/artarts36/GolandProjects/gomodfinder/go.mod",
	}

	got := mod.CalcPackageFromAbsPath("/Users/artarts36/GolandProjects/gomodfinder/a/b/c/d")

	assert.Equal(t, &Package{
		Name:               "d",
		ModuleRelativePath: "a/b/c/d",
		Module:             mod,
	}, got)
}
