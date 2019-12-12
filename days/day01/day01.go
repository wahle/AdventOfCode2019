package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/wahle/AdventOfCode2019/fileReader"
	"github.com/wahle/AdventOfCode2019/module"
	"github.com/wahle/AdventOfCode2019/ship"
)

func main() {
	inputs := fileReader.ReadLines("days/day01/inputs.txt")
	spaceShip := ship.New()

	for i, eachline := range inputs {
		Mass, err := strconv.Atoi(eachline)
		if err != nil {
			log.Fatalf("Failed to convert %s to int: %s", eachline, err)
		}
		spaceShip.AddModule(module.New(i, Mass))
	}

	fmt.Printf("ComponentFuel Needed %d\n", spaceShip.ModulesFuel)
	fmt.Printf("TotalFuel Needed %d\n", spaceShip.TotalFuel)
}
