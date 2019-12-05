package wireguppy

import (
	"errors"
	"fmt"
	"strconv"
)

// WireInstruction represents a single direction and distance for a wire trace
type WireInstruction struct {
	Direction CompassDirection
	Distance  int
}

// Coordinate represents a cartesian coordinate in 2 dimensions
type Coordinate struct {
	X int
	Y int
}

// Add a vector to a given coordinate
func (start Coordinate) Add(offset Coordinate) Coordinate {
	return Coordinate{start.X + offset.X, start.Y + offset.Y}
}

// CompassDirection represents a direction on the compass
type CompassDirection int

// The compass directions
const (
	Up    CompassDirection = 0
	Right CompassDirection = 1
	Down  CompassDirection = 2
	Left  CompassDirection = 3
)

func (direction CompassDirection) String() string {
	names := []string{
		"Up",
		"Right",
		"Down",
		"Left",
	}

	if direction < Up || direction > Left {
		return "Unknown"
	}

	return names[direction]
}

func parseDirection(direction string) (CompassDirection, error) {
	encodings := map[string]CompassDirection{
		"u": Up,
		"U": Up,
		"r": Right,
		"R": Right,
		"d": Down,
		"D": Down,
		"l": Left,
		"L": Left,
	}

	compassDirection, ok := encodings[direction]

	if ok {
		return compassDirection, nil
	}

	return compassDirection, fmt.Errorf("Could not parse direction %s", direction)
}

// Parse a string like U42 into a WireInstruction
func Parse(instruction string) (WireInstruction, error) {
	var returnValue WireInstruction
	var err error
	var direction CompassDirection

	if len(instruction) <= 1 {
		return returnValue, fmt.Errorf("Instruction %s is not long enough to be valid", instruction)
	}

	direction, err = parseDirection(instruction[:1])

	if err != nil {
		return returnValue, err
	}

	returnValue.Direction = direction

	returnValue.Distance, err = strconv.Atoi(instruction[1:])

	return returnValue, err
}

// ParseWires Turn an array of stringified wire instructions into a typed array of arrays
func ParseWires(wireArrays [][]string) ([][]WireInstruction, error) {
	var wires [][]WireInstruction
	var wire []WireInstruction
	var err error
	for _, wireArray := range wireArrays {

		wire, err = parseWire(wireArray)
		if err != nil {
			return nil, err
		}

		wires = append(wires, wire)
	}

	return wires, nil
}

func parseWire(wireArray []string) ([]WireInstruction, error) {
	var instructions []WireInstruction
	var instruction WireInstruction
	var err error
	for _, wireStep := range wireArray {
		instruction, err = Parse(wireStep)
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, instruction)

	}

	if len(instructions) == 0 {
		return nil, errors.New("Empty wire is not valid")
	}
	return instructions, nil
}

// Calculates the manhattan distance between two points
func CalculateManhattan(point1 Coordinate, point2 Coordinate) int {
	return abs(point1.X-point2.X) + abs(point1.Y-point2.Y)
}

// Calculates the manhattan distance of a point from the origin
func CalculateManhattanFromOrigin(point Coordinate) int {
	return CalculateManhattan(point, Coordinate{0, 0})
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// WireCoordinates traverse a set of wire instructions and list all of the Coordinates through which they pass
func WireCoordinates(wire []WireInstruction) []Coordinate {

	coordinateMoves := map[CompassDirection]Coordinate{
		// Basic moves for a coordinate
		Up:    Coordinate{0, 1},
		Right: Coordinate{1, 0},
		Down:  Coordinate{0, -1},
		Left:  Coordinate{-1, 0},
	}

	var returnValue []Coordinate
	current := Coordinate{0, 0}

	for _, instruction := range wire {
		for i := 0; i < instruction.Distance; i++ {
			current = current.Add(coordinateMoves[instruction.Direction])
			returnValue = append(returnValue, current)
		}
	}

	return returnValue
}

// FindCommonCoordinates Find the coordinates common to two lists
func FindCommonCoordinates(list1 []Coordinate, list2 []Coordinate) []Coordinate {
	var common []Coordinate

	for _, coordinate1 := range list1 {
		for _, coordinate2 := range list2 {
			if coordinate1.X == coordinate2.X && coordinate1.Y == coordinate2.Y {
				common = append(common, coordinate1)
			}
		}
	}

	return common
}
