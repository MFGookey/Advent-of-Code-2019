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
	program, err := stringToIntArray(string(dat), ",")
	check(err)
	
	var programResult []int

	programResult, err = Intcode.Execute(program)
	if err != nil {
		panic(err)
	} else{
		fmt.Println(programResult)
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
