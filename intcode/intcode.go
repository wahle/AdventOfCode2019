package intcode

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Comp is struct that has an InitialState array and the OutPutState array. The OutputState array is the processed initialState.
type Comp struct {
	State         []int
	IntsructPoint int
}

// NewComp generates a new Intcode Computer object
func NewComp(input interface{}) Comp { // This New function is outrageous and was more just an excersise in proving it was possible
	compState := []int{}
	switch reflect.TypeOf(input).String() {
	case "string":
		for _, eachInt := range strings.Split(reflect.ValueOf(input).String(), ",") {
			if actualInt, err := strconv.Atoi(eachInt); err != nil {
				panic(err)
			} else {
				compState = append(compState, actualInt)
			}
		}
	case "[]int":
		compState = input.([]int)

	default:
		panic("Either no values or incorrect input format for New computer Object was passed to constructor")
	}

	computer := Comp{
		State:         compState,
		IntsructPoint: 0,
	}
	return computer
}

// ProcessProgram executes a single program
func (comp *Comp) ProcessProgram(instPoint int) {
	switch comp.State[instPoint] {
	case 1: //Addition
		comp.State[comp.State[instPoint+3]] = comp.State[comp.State[instPoint+1]] + comp.State[comp.State[instPoint+2]]
		comp.IntsructPoint = instPoint + 4
	case 2: //Multiplication
		comp.State[comp.State[instPoint+3]] = comp.State[comp.State[instPoint+1]] * comp.State[comp.State[instPoint+2]]
		comp.IntsructPoint = instPoint + 4
	case 99: //Exit Code
		fmt.Printf("case = 99")
	default:
		panic(fmt.Sprintf("No Operation for program. Input given: %v", comp.State[instPoint]))
	}
}

// Process executes all of a computers programs against its current state.
func (comp *Comp) Process() {
	for comp.IntsructPoint < len(comp.State) && comp.State[comp.IntsructPoint] != 99 {
		//fmt.Printf("Opperation: %v, Noun: %v, Verb: %v\n", comp.State[comp.IntsructPoint], comp.State[comp.IntsructPoint+1], comp.State[comp.IntsructPoint+2])
		comp.ProcessProgram(comp.IntsructPoint)
	}
}
