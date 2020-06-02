package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	opCodeAdd  = 1
	opCodeMult = 2
	opCodeExit = 99
)

var (
	inputText = `1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,6,19,23,2,6,23,27,1,5,27,31,2,31,9,
35,1,35,5,39,1,39,5,43,1,43,10,47,2,6,47,51,1,51,5,55,2,55,6,59,1,5,59,63,2,63,6,67,1,5,67,71,
1,71,6,75,2,75,10,79,1,79,5,83,2,83,6,87,1,87,5,91,2,9,91,95,1,95,6,99,2,9,99,103,2,9,103,107,
1,5,107,111,1,111,5,115,1,115,13,119,1,13,119,123,2,6,123,127,1,5,127,131,1,9,131,135,1,135,9,
139,2,139,6,143,1,143,5,147,2,147,6,151,1,5,151,155,2,6,155,159,1,159,2,163,1,9,163,0,99,2,0,14,0`
)

func copyArr(in []int) []int {
	tmp := make([]int, len(in))
	copy(tmp, in)
	return tmp
}

func opAdd(codeRun int, memory *[]int) int {
	var indexLeft = (*memory)[codeRun+1]
	var indexRight = (*memory)[codeRun+2]
	var indexResult = (*memory)[codeRun+3]

	//fmt.Printf("*add mem[%d]=%d mem[%d]=%d", indexLeft, (*memory)[indexLeft], indexRight, (*memory)[indexRight])
	(*memory)[indexResult] = (*memory)[indexLeft] + (*memory)[indexRight]
	//fmt.Printf(" result=%d\n", (*memory)[indexRight])
	return codeRun + 4
}

func opMult(codeRun int, memory *[]int) int {
	var indexLeft = (*memory)[codeRun+1]
	var indexRight = (*memory)[codeRun+2]
	var indexResult = (*memory)[codeRun+3]

	//fmt.Printf("*mult mem[%d]=%d mem[%d]=%d", indexLeft, (*memory)[indexLeft], indexRight, (*memory)[indexRight])
	(*memory)[indexResult] = (*memory)[indexLeft] * (*memory)[indexRight]
	//fmt.Printf(" result=%d\n", (*memory)[indexRight])
	return codeRun + 4
}

func execute(memory []int) []int {

	//fmt.Printf("code %v\n", memory)

	var codeRun = 0
	var done = false
	for !done {
		opcode := memory[codeRun]

		//fmt.Printf("execute code %d at addr %d\n", opcode, codeRun)
		switch opcode {
		default:
			fmt.Printf("Unknown instruction %d at addr %d\n", opcode, codeRun)
			os.Exit(99)
		case opCodeAdd:
			codeRun = opAdd(codeRun, &memory)
		case opCodeMult:
			codeRun = opMult(codeRun, &memory)
		case opCodeExit:
			done = true
		}
	}

	//fmt.Printf("result %v\n", memory)
	return memory
}

func executeWithParams(memory []int, noun int, verb int) int {
	memory[1] = noun
	memory[2] = verb
	result := execute(memory)
	return result[0]
}

func memoryFromString(inputText string) []int {
	inputValuesString := strings.Split(strings.ReplaceAll(inputText, "\n", ""), ",")
	input := make([]int, len(inputValuesString))
	for inputPosition := range inputValuesString {
		converted, err := strconv.Atoi(inputValuesString[inputPosition])
		if err != nil {
			panic(err)
		}
		input[inputPosition] = converted
	}

	return input
}

func bruteForce(value int) {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			fmt.Printf("bruteForce noun=%d verb=%d ", noun, verb)
			result := executeWithParams(memoryFromString(inputText), noun, verb)
			fmt.Printf("result %d\n", result)

			if value == result {
				fmt.Printf("found noun=%d verb=%d result=%d\n", noun, verb, result)
				fmt.Printf("result %d\n", 100*noun+verb)
				os.Exit(99)
			}
		}
	}
}

func main() {

	// 76092639
	// 19690720
	input := memoryFromString(inputText)
	//fmt.Printf("input before  %v\n", input)
	// replace position 1 with the value 12 and replace position 2 with the value 2
	input[1] = 12
	input[2] = 2
	//fmt.Printf("input after patch  %v\n", input)

	output := execute(input)

	fmt.Printf("output  %v\n", output)
	fmt.Printf("result output[0]  %v\n", output[0])

	result := executeWithParams(memoryFromString(inputText), 12, 2)
	fmt.Printf("executeWithParams output  %v\n", result)

	bruteForce(19690720)
}
