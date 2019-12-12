package main

import (
	"github.com/wahle/AdventOfCode2019https://github.com/wahle/AdventOfCode2019/fileReader"
	"github.com/wahle/AdventOfCode2019https://github.com/wahle/AdventOfCode2019/wire"
)

func main() {
	wire1 := wire.NewWire()
	wire2 := wire.NewWire()
	wire1.Movements = wire.DecodeMoveOrders(fileReader.ReadLine("days/day03/input.txt", 0))
	wire2.Movements = wire.DecodeMoveOrders(fileReader.ReadLine("days/day03/input.txt", 1))
	wire1.MapWire()
	wire2.MapWire()
	intersects := wire.FindIntersects(wire1, wire2)
	_ = wire.FindShortestPath(intersects)
	_ = wire.FindLeastSteps(wire1, wire2, intersects)
}
