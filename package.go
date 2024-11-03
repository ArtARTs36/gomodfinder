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

func (p *Package) Equal(that *Package) bool {
	return p.Module.Module.Mod.Path == that.Module.Module.Mod.Path && p.ModuleRelativePath == that.ModuleRelativePath
}
