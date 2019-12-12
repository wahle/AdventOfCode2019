package main

import (
	"fmt"

	"github.com/wahle/AdventOfCode2019/fileReader"
	"github.com/wahle/AdventOfCode2019/intcode"
)

func main() {
	inputs := fileReader.ReadLine("days/day02/inputs.txt", 0)

	i := 0
	j := 0
	for i = 0; i <= 99; i++ {
		for j = 0; j <= 99; j++ {
			intCodeComp := intcode.NewComp(inputs)
			intCodeComp.State[1] = i
			intCodeComp.State[2] = j
			intCodeComp.Process()
			if intCodeComp.State[0] == 19690720 {
				fmt.Printf("Noun: %d, Verb: %d\n", i, j)
				break
			}
		}
	}
}
