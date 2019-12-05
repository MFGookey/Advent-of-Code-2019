package main

import (
	//"io"
	"fmt"
	"io/ioutil"
	"strings"
	//"os"
	guppy "github.com/MFGookey/Advent-of-Code-2019/day-03/wireguppy"
	//"strconv"
)

func main() {
	wireArrays, err := readMultiLineCSV("input.txt")
	check(err)

	wires, err := guppy.ParseWires(wireArrays)
	check(err)

	var wiresCoordinates [][]guppy.Coordinate
	for _, wire := range wires {
		wiresCoordinates = append(wiresCoordinates, guppy.WireCoordinates(wire))
	}

	common := guppy.FindCommonCoordinates(wiresCoordinates[0], wiresCoordinates[1])

	for i:=2;i<len(wiresCoordinates);i++ {
		common = guppy.FindCommonCoordinates(common, wiresCoordinates[i])
	}

	fmt.Println(common)

	smallest := guppy.CalculateManhattanFromOrigin(common[0])
	var current int
	for _, coordinate := range common {
		current = guppy.CalculateManhattanFromOrigin(coordinate)
		if current < smallest {
			smallest = current
		}
	}

	fmt.Println(smallest)

	//fmt.Println(wireArrays)
	//fmt.Println(wires)
	//fmt.Println(wiresCoordinates)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func breakStringOnNewlines(data string) []string {
	return strings.Split(data, "\r\n")
}

func breakStringOnCommas(data string) []string {
	return strings.Split(data, ",")
}

func readMultiLineCSV(path string) ([][]string, error) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return [][]string{{}}, err
	}

	stringified := string(dat)
	lines := breakStringOnNewlines(stringified)
	var fields [][]string
	for _, line := range lines {
		fields = append(fields, breakStringOnCommas(line))
	}

	return fields, nil
}
