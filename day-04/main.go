package main

import (
	"fmt"
	cracker "github.com/MFGookey/Advent-of-Code-2019/day-04/cracker"
)

func main() {
	passwordCount, err := cracker.CountPasswords(240298, 784956)
	check(err)
	fmt.Println(passwordCount)
}

func check(err error){
	if err != nil {
		panic(err)
	}
}