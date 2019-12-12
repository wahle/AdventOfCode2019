package wire

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// Wire is arrays of points.
type Wire struct {
	Path      []Coordinate
	Movements []MoveOrder
}

// Coordinate is a 2D coordinate containing an x and y value
type Coordinate struct {
	x int
	y int
}

// MoveOrder tells the traverse function how to create the wire
type MoveOrder struct {
	direction string
	distance  int
}

//NewWire returns a new wire with no
func NewWire() Wire {
	origin := Coordinate{x: 0, y: 0}
	coords := []Coordinate{origin}

	return Wire{Path: coords}
}

// DecodeMoveOrders takes string with comma separated move orders and returns and array of parsed moveOrders
func DecodeMoveOrders(moves string) []MoveOrder {
	listOfMoves := []MoveOrder{}

	for _, eachMove := range strings.Split(moves, ",") {
		eachDirection := string(eachMove[0])
		eachDistance := eachMove[1:len(eachMove)]
		if actualDistance, err := strconv.Atoi(eachDistance); err != nil {
			panic(err)
		} else {
			actualMove := MoveOrder{eachDirection, actualDistance}
			listOfMoves = append(listOfMoves, actualMove)
		}
	}
	return listOfMoves
}

func processMovement(wire *Wire, movement MoveOrder) {
	switch movement.direction {
	case "U":
		for i := 0; i < movement.distance; i++ {
			lastCoord := wire.Path[len(wire.Path)-1]
			wire.Path = append(wire.Path, Coordinate{x: lastCoord.x, y: lastCoord.y + 1})
		}
	case "D":
		for i := 0; i < movement.distance; i++ {
			lastCoord := wire.Path[len(wire.Path)-1]
			wire.Path = append(wire.Path, Coordinate{x: lastCoord.x, y: lastCoord.y - 1})
		}

	case "R":
		for i := 0; i < movement.distance; i++ {
			lastCoord := wire.Path[len(wire.Path)-1]
			wire.Path = append(wire.Path, Coordinate{x: lastCoord.x + 1, y: lastCoord.y})
		}

	case "L":
		for i := 0; i < movement.distance; i++ {
			lastCoord := wire.Path[len(wire.Path)-1]
			wire.Path = append(wire.Path, Coordinate{x: lastCoord.x - 1, y: lastCoord.y})
		}

	default:
		panic(fmt.Sprintf("The direction: %v is not valid", movement.direction))

	}
}

//MapWire will create the full path of the wire
func (wire *Wire) MapWire() {
	for i := 0; i < len(wire.Movements); i++ {
		processMovement(wire, wire.Movements[i])
	}
}

//StepsUnitlPoint returns the number of steps a wire takes before landing on a given coordinate
func StepsUnitlPoint(wire Wire, coord Coordinate) int {
	numberOfSteps := 0
	for i := 0; i < len(wire.Path); i++ {
		if wire.Path[i] == coord {
			numberOfSteps = i
		}
	}
	return numberOfSteps
}

func equalCoords(firstCoord Coordinate, secondCord Coordinate) bool {
	equalValues := false
	if firstCoord.x == secondCord.x && firstCoord.y == secondCord.y {
		equalValues = true
	}
	return equalValues
}

// FindIntersects finds the intersect locations for two wires
func FindIntersects(firstWire Wire, secondWire Wire) []Coordinate {
	intersects := []Coordinate{}
	//Set Up multiThreaded for loop
	var wg sync.WaitGroup
	wg.Add(len(firstWire.Path))
	fmt.Printf("\nProcessing Intersects\n")
	for i := 0; i < len(firstWire.Path); i++ {
		go func(i int) {
			for j := 0; j < len(secondWire.Path); j++ {
				if equalCoords(firstWire.Path[i], secondWire.Path[j]) {
					intersects = append(intersects, firstWire.Path[i])
				}
			}
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	return intersects
}

// FindShortestPath will return the closet point to coordinate 0,0 for a given array of coordinates
func FindShortestPath(intersects []Coordinate) Coordinate {
	closestCoord := intersects[0]
	shortestDistance := 0
	for i := 0; i < len(intersects); i++ {
		//find absolutes
		x := intersects[i].x
		if x < 0 {
			x = -intersects[i].x
		}
		y := intersects[i].y
		if y < 0 {
			y = -intersects[i].y
		}
		distance := x + y
		if distance < shortestDistance || shortestDistance == 0 {
			shortestDistance = distance
			closestCoord = intersects[i]
		}
	}
	fmt.Printf("\nShortest Distance: %d Coordinates: %v\n", shortestDistance, closestCoord)
	return closestCoord
}

//FindLeastSteps returns the shortest list of possible steps to get a intesect
func FindLeastSteps(wire1 Wire, wire2 Wire, intersects []Coordinate) int {
	fewestSteps := 0
	fastestIntersect := intersects[0]
	for i := 0; i < len(intersects); i++ {
		numberOfSteps := StepsUnitlPoint(wire1, intersects[i]) + StepsUnitlPoint(wire2, intersects[i])
		if numberOfSteps < fewestSteps || fewestSteps == 0 {
			fewestSteps = numberOfSteps
			fastestIntersect = intersects[i]
		}
	}
	fmt.Printf("\nLeast Steps: %v FewestSteps: %d\n", fastestIntersect, fewestSteps)
	return fewestSteps
}
