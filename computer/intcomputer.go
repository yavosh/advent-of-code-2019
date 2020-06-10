package computer

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	opCodeAdd         = 1
	opCodeMult        = 2
	opCodeInput       = 3
	opCodeOutput      = 4
	opCodeJumpIfTrue  = 5
	opCodeJumpIfFalse = 6
	opCodeLessThan    = 7
	opCodeEquals      = 8
	opCodeExit        = 99
)

func copyArr(in []int) []int {
	tmp := make([]int, len(in))
	copy(tmp, in)
	return tmp
}

func getValueImmediate(position int, mem *[]int) int {
	return (*mem)[position]
}

func getValueByAddress(position int, mem *[]int) int {
	address := (*mem)[position]
	return (*mem)[address]
}

func getValueBy(position int, flag int, mem *[]int) int {
	if flag > 0 {
		return getValueImmediate(position, mem)
	}

	return getValueByAddress(position, mem)
}

func putValueAtAddress(position int, value int, mem *[]int) int {
	address := (*mem)[position]
	(*mem)[address] = value
	return (*mem)[address]
}

func opInput(codeRun int, flags *[]int, mem *[]int, input *[]int) int {

	inputValue := (*input)[0]
	left := getValueImmediate(codeRun+1, mem)
	// left := getValueBy(codeRun+1, (*flags)[0], mem)
	// fmt.Printf("putValueAtAddress left=%d inputValue=%d mem[]=%d\n", left, inputValue, (*mem)[codeRun+1])
	(*mem)[left] = inputValue
	return codeRun + 2
}

func opOutput(codeRun int, flags *[]int, mem *[]int, output *[]int) int {

	if (*flags)[0] > 0 {
		left := getValueImmediate(codeRun+1, mem)
		fmt.Printf("DEBUG: codeRun=%d left=%d \n", codeRun, left)
		(*output)[0] = left
		return codeRun + 2
	}

	left := getValueByAddress(codeRun+1, mem)
	fmt.Printf("DEBUG: codeRun=%d left=%d \n", codeRun, left)

	(*output)[0] = left
	return codeRun + 2
}

func opAdd(codeRun int, flags *[]int, mem *[]int) int {
	left := getValueBy(codeRun+1, (*flags)[0], mem)
	right := getValueBy(codeRun+2, (*flags)[1], mem)
	// fmt.Printf("OP Add flags=%v left=%d right=%d\n", flags, left, right)
	putValueAtAddress(codeRun+3, left+right, mem)
	return codeRun + 4
}

func opMult(codeRun int, flags *[]int, mem *[]int) int {
	left := getValueBy(codeRun+1, (*flags)[0], mem)
	right := getValueBy(codeRun+2, (*flags)[1], mem)
	// fmt.Printf("OP Mul flags=%v left=%d right=%d\n", flags, left, right)
	putValueAtAddress(codeRun+3, left*right, mem)
	return codeRun + 4
}

func opJumpTrue(codeRun int, flags *[]int, mem *[]int) int {
	left := getValueBy(codeRun+1, (*flags)[0], mem)
	right := getValueBy(codeRun+2, (*flags)[1], mem)
	if left > 0 {
		return right
	}
	return codeRun + 3
}

func opJumpFalse(codeRun int, flags *[]int, mem *[]int) int {
	left := getValueBy(codeRun+1, (*flags)[0], mem)
	right := getValueBy(codeRun+2, (*flags)[1], mem)
	if left == 0 {
		return right
	}
	return codeRun + 3
}

func opLessThan(codeRun int, flags *[]int, mem *[]int) int {
	left := getValueBy(codeRun+1, (*flags)[0], mem)
	right := getValueBy(codeRun+2, (*flags)[1], mem)

	if left < right {
		putValueAtAddress(codeRun+3, 1, mem)
		return codeRun + 4

	}

	putValueAtAddress(codeRun+3, 0, mem)
	return codeRun + 4
}

func opEquals(codeRun int, flags *[]int, mem *[]int) int {
	left := getValueBy(codeRun+1, (*flags)[0], mem)
	right := getValueBy(codeRun+2, (*flags)[1], mem)

	if left == right {
		putValueAtAddress(codeRun+3, 1, mem)
		return codeRun + 4
	}

	putValueAtAddress(codeRun+3, 0, mem)
	return codeRun + 4
}

// Run execute a program in the int computer
// accepts program memory and inputs
// returns program memory and outputs
func Run(memory []int, input []int) ([]int, []int) {

	var output = make([]int, len(input))

	var codeRun = 0
	var done = false
	for !done {
		instruction := memory[codeRun]
		opcode := instruction % 100

		flags := []int{
			(instruction / 100) % 10,
			(instruction / 1000) % 10,
			(instruction / 10000) % 10,
		}

		// fmt.Printf("execute code instruction=%d opcode=%d flags=%v at addr %d\n",
		// 	instruction, opcode, flags, codeRun)
		// fmt.Printf("mem=%v input=%v output=%v\n",
		// 	memory, input, output)

		switch opcode {
		default:
			fmt.Printf("Unknown instruction %d at addr %d\n", opcode, codeRun)
			os.Exit(99)
		case opCodeAdd:
			codeRun = opAdd(codeRun, &flags, &memory)
		case opCodeMult:
			codeRun = opMult(codeRun, &flags, &memory)
		case opCodeInput:
			codeRun = opInput(codeRun, &flags, &memory, &input)
		case opCodeOutput:
			codeRun = opOutput(codeRun, &flags, &memory, &output)
		case opCodeJumpIfTrue:
			codeRun = opJumpTrue(codeRun, &flags, &memory)
		case opCodeJumpIfFalse:
			codeRun = opJumpFalse(codeRun, &flags, &memory)
		case opCodeLessThan:
			codeRun = opLessThan(codeRun, &flags, &memory)
		case opCodeEquals:
			codeRun = opEquals(codeRun, &flags, &memory)
		case opCodeExit:
			done = true
		}
	}

	//fmt.Printf("result memory %v\n", memory)
	//fmt.Printf("result output  %v\n", output)
	return memory, output
}

// LoadInstructions .
func LoadInstructions(inputText string) []int {
	inputClean := strings.ReplaceAll(inputText, "\n", "")
	inputClean = strings.ReplaceAll(inputClean, " ", "")
	inputClean = strings.ReplaceAll(inputClean, "\t", "")

	inputValuesString := strings.Split(
		inputClean,
		",")
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
