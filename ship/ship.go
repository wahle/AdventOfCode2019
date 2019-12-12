package ship

import (
	"github.com/wahle/AdventOfCode2019/module"
)

type Ship struct {
	Modules     []module.Module
	TotalFuel   int
	ModulesFuel int
}

func New() Ship {
	s := Ship{}
	s.Modules = []module.Module{}
	s.TotalFuel = 0
	s.ModulesFuel = 0

	return s
}

func (s *Ship) AddModule(m module.Module) {
	s.Modules = append(s.Modules, m)
	s.ModulesFuel, s.TotalFuel = s.findFuelNeeded()
}

func (s Ship) findFuelNeeded() (int, int) {
	modulesFuelNeeded := 0
	totalFuelNeeded := 0

	for _, eachModule := range s.Modules {
		modulesFuelNeeded = modulesFuelNeeded + eachModule.ComponentFuel
		totalFuelNeeded = totalFuelNeeded + eachModule.TotalFuel
	}

	return modulesFuelNeeded, totalFuelNeeded
}
