package intcode

import (
	"reflect"
	"testing"
)

func TestValidateCompString(t *testing.T) {
	input := "1,0,0,0,99"
	state := []int{1, 0, 0, 0, 99}
	createdComp := NewComp(input)
	if !reflect.DeepEqual(createdComp.State, state) {
		t.Errorf("Created Comp has improper state: %v did not match expected %v", createdComp.State, state)
	}
	if createdComp.IntsructPoint != 0 {
		t.Errorf("Created Comp has improper Instruction Point: %v did not match expected %v", createdComp.IntsructPoint, 0)
	}
}

func TestValidateCompIntArray(t *testing.T) {
	input := []int{1, 0, 0, 0, 99}
	createdComp := NewComp(input)
	if !reflect.DeepEqual(createdComp.State, input) {
		t.Errorf("Created Comp has improper state: %v did not match expected %v", createdComp.State, input)
	}
	if createdComp.IntsructPoint != 0 {
		t.Errorf("Created Comp has improper Instruction Point: %v did not match expected %v", createdComp.IntsructPoint, 0)
	}
}

func TestAdditionProgram(t *testing.T) {
	input := []int{1, 0, 0, 0, 99}
	newState := []int{2, 0, 0, 0, 99}
	createdComp := NewComp(input)
	createdComp.Process()
	if !reflect.DeepEqual(createdComp.State, newState) {
		t.Errorf("Created Comp has improper state: %v did not match expected %v", createdComp.State, newState)
	}
}

func TestMultiplicationProgram(t *testing.T) {
	input := []int{2, 3, 0, 3, 99}
	newState := []int{2, 3, 0, 6, 99}
	createdComp := NewComp(input)
	createdComp.Process()
	if !reflect.DeepEqual(createdComp.State, newState) {
		t.Errorf("Created Comp has improper state: %v did not match expected %v", createdComp.State, newState)
	}
}

func TestMoreComplexProgram(t *testing.T) {
	input := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	newState := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	createdComp := NewComp(input)
	createdComp.Process()
	if !reflect.DeepEqual(createdComp.State, newState) {
		t.Errorf("Created Comp has improper state: %v did not match expected %v", createdComp.State, newState)
	}
}
