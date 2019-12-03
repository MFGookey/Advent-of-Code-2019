package Intcode

import (
	"fmt"
)

// Execute execute a program written in the Intcode language
func Execute(program []int) ([]int, error) {
	var halted bool = false
	var err error
	for programCounter := 0; halted == false && err == nil; programCounter++ {
		//fmt.Printf("Program Counter: %d, halted: %v\n", programCounter, halted)
		//fmt.Println(program)
		program, halted, err = executeStep(programCounter, program)

		if err != nil {
			return program, err
		}
	}

	return program, nil
}

func executeStep(programCounter int, program []int) ([]int, bool, error) {
	var instruction []int
	var executeResult []int
	var err error
	instructionTypicalLength := 4

	if programCounter*instructionTypicalLength > len(program) {
		return program, true, fmt.Errorf("program counter %d has exceeded program length", programCounter)
	}

	if programCounter*instructionTypicalLength+instructionTypicalLength > len(program) {
		instruction = program[programCounter*instructionTypicalLength:]
	} else {
		instruction = program[programCounter*instructionTypicalLength : programCounter*instructionTypicalLength+instructionTypicalLength]
	}

	//fmt.Println(instruction)

	if instruction[0] == 99 {
		//fmt.Println("halting")
		return program, true, nil
	}

	if len(instruction) == instructionTypicalLength {
		switch instruction[0] {
		case 1:
			executeResult, err = executeAdd(instruction[1], instruction[2], instruction[3], program)
			break
		case 2:
			executeResult, err = executeMultiply(instruction[1], instruction[2], instruction[3], program)
			break
		default:
			err = fmt.Errorf("unknown opcode at instruction %d: %d", programCounter, instruction[0])
		}

		var forceHalt bool

		forceHalt = (err != nil)

		return executeResult, forceHalt, err
	}

	return program, true, fmt.Errorf("Instruction %v is not of an expected length", instruction)
}

func executeAdd(leftAddress int, rightAddress int, resultAddress int, program []int) ([]int, error) {
	if leftAddress >= len(program) {
		return program, fmt.Errorf("leftAddress %d is not within memory", leftAddress)
	}

	if rightAddress >= len(program) {
		return program, fmt.Errorf("rightAddress %d is not within memory", rightAddress)
	}

	if resultAddress >= len(program) {
		return program, fmt.Errorf("resultAddress %d is not within memory", resultAddress)
	}

	result := add(program[leftAddress], program[rightAddress])
	//fmt.Printf("Adding: %d + %d = %d, storing to %d\n", program[leftAddress], program[rightAddress], result, resultAddress)

	program[resultAddress] = result
	return program, nil
}

func executeMultiply(leftAddress int, rightAddress int, resultAddress int, program []int) ([]int, error) {
	if leftAddress >= len(program) {
		return program, fmt.Errorf("leftAddress %d is not within memory", leftAddress)
	}

	if rightAddress >= len(program) {
		return program, fmt.Errorf("rightAddress %d is not within memory", rightAddress)
	}

	if resultAddress >= len(program) {
		return program, fmt.Errorf("resultAddress %d is not within memory", resultAddress)
	}

	result := multiply(program[leftAddress], program[rightAddress])
	//fmt.Printf("Multiplying: %d * %d = %d, storing to %d\n", program[leftAddress], program[rightAddress], result, resultAddress)

	program[resultAddress] = result

	return program, nil
}

func add(left int, right int) int {
	return left + right
}

func multiply(left int, right int) int {
	return left * right
}
