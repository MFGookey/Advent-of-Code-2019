package main

import (
	"fmt"
	//"io"
	"io/ioutil"
	//"os"
	"strings"
	"github.com/MFGookey/Advent-of-Code-2019/day-02/Intcode"
	"strconv"
)

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	var program []int
	program, err = stringToIntArray(string(dat), ",")
	check(err)

	var programResult []int
	

	multiplier := 1
	max := len(program) * multiplier

	// Reading opcodes is for scrubs.  Do a search of the solution space instead.
	for i:=0; i< max; i++ {
		for j:= 0;j<max; j++ {
			program, err = stringToIntArray(string(dat), ",")
			program[1] = i
			program[2] = j
			programResult, err = Intcode.Execute(program)
			if err != nil {
				fmt.Println(err)
			//} else{
				//fmt.Println(stringToIntArray(string(dat), ","))
				//fmt.Println(programResult)
			}

			if programResult[0] == 19690720 {
				fmt.Println(programResult[:4])
				return
			}
		}
	}
}

func stringToIntArray(toSplit string, delimeter string) ([]int, error) {
	splitString := strings.Split(toSplit, delimeter)
	decodedInts := make([]int, len(splitString))
	var err error
	for i:=0; i<len(splitString); i++ {
		decodedInts[i], err = strconv.Atoi(splitString[i])
		if err != nil {
			return nil, err
		}
	}

	return decodedInts, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
