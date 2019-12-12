package module

import (
	"testing"
)

func TestValidateModule(t *testing.T) {
	m := New(1, 1969)
	if m.Mass != 1969 {
		t.Errorf("Mass: %v did not match expected %v", m.Mass, 1969)
	}
	if m.Number != 1 {
		t.Errorf("Number: %v did not match expected %v", m.Number, 1)
	}
}

func TestFuelNeeded(t *testing.T) {
	fuel := fuelRequirement(1969)
	if fuel != 654 {
		t.Errorf("Fuel: %v did not match expected %v", fuel, 654)
	}
}

func TestFuel4Fuel(t *testing.T) {
	fuel := fuel4Fuel(654)
	if fuel != 966 {
		t.Errorf("Fuel: %v did not match expected %v", fuel, 966)
	}
}
