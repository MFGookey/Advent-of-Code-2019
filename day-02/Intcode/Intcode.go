package Intcode

import (
	"fmt"
)

// Execute execute a program written in the Intcode language
func Execute(program []int) ([]int, error) {
	var halted bool = false
	var err error
	for programCounter := 0; halted == false && err == nil; programCounter++ {
		fmt.Printf("Program Counter: %d, halted: %v\n", programCounter, halted)
		fmt.Println(program)
		program, halted, err = executeStep(programCounter, program)

		if err != nil {
			return program, err
		}
	}

	return program, nil
}

func executeStep(programCounter int, program []int) ([]int, bool, error) {
	var instruction []int
	instructionTypicalLength := 4

	if programCounter*instructionTypicalLength > len(program) {
		return program, true, fmt.Errorf("program counter %d has exceeded program length", programCounter)
	}

	if programCounter*instructionTypicalLength+instructionTypicalLength > len(program) {
		instruction = program[programCounter*instructionTypicalLength:]
	} else {
		instruction = program[programCounter*instructionTypicalLength : programCounter*instructionTypicalLength+instructionTypicalLength]
	}
	fmt.Println(instruction)
	if instruction[0] == 99 {
		fmt.Println("halting")
		return program, true, nil
	}

	if len(instruction) == instructionTypicalLength {
		switch instruction[0] {
		case 1:
			return executeAdd(instruction[1], instruction[2], instruction[3], program), false, nil
		case 2:
			return executeMultiply(instruction[1], instruction[2], instruction[3], program), false, nil
		default:
			return program, true, fmt.Errorf("unknown opcode at instruction %d: %d", programCounter, instruction[0])
		}
	} else {
		return program, true, fmt.Errorf("Instruction %v is not of an expected length", instruction)
	}
}

func executeAdd(leftAddress int, rightAddress int, resultAddress int, program []int) []int {
	result := add(program[leftAddress], program[rightAddress])
	fmt.Printf("Adding: %d + %d = %d, storing to %d\n", program[leftAddress], program[rightAddress], result, resultAddress)
	program[resultAddress] = result
	return program
}

func executeMultiply(leftAddress int, rightAddress int, resultAddress int, program []int) []int {
	result := multiply(program[leftAddress], program[rightAddress])
	fmt.Printf("Multiplying: %d * %d = %d, storing to %d\n", program[leftAddress], program[rightAddress], result, resultAddress)
	program[resultAddress] = result

	return program
}

func add(left int, right int) int {
	return left + right
}

func multiply(left int, right int) int {
	return left * right
}
