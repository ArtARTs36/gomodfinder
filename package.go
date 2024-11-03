package gomodfinder

import "fmt"

type Package struct {
	Name               string
	ModuleRelativePath string

	Module *ModFile
}

func (p *Package) Child(name string) *Package {
	return &Package{
		Name:               name,
		ModuleRelativePath: fmt.Sprintf("%s/%s", p.ModuleRelativePath, name),

		Module: p.Module,
	}
}
