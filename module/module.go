package module

type Module struct {
	Number        int
	Mass          int
	ComponentFuel int
	TotalFuel     int
}

func New(moduleNumber int, moduleMass int) Module {
	m := Module{moduleNumber, moduleMass, 0, 0}
	m.ComponentFuel = fuelRequirement(m.Mass)
	m.TotalFuel = fuel4Fuel(m.ComponentFuel)
	return m
}

func fuelRequirement(m int) int {
	fuelNeeded := (m / 3) - 2
	if fuelNeeded < 0 {
		return 0
	}
	return fuelNeeded
}

func fuel4Fuel(f int) int {
	if f > 0 {
		fuelNeeds := fuelRequirement(f)
		f = f + fuel4Fuel(fuelNeeds)
	}
	return f
}
