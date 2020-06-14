package computer

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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
	opCodeSetBase     = 9
	opCodeExit        = 99
)

func copyArr(in []int64) []int64 {
	tmp := make([]int64, len(in))
	copy(tmp, in)
	return tmp
}

func getValueImmediate(position int64, mem *[]int64) int64 {
	return (*mem)[position]
}

func getValueByAddress(position int64, mem *[]int64) int64 {
	address := (*mem)[position]
	return (*mem)[address]
}

func getValueByAddressRelative(baseAddress int64, position int64, mem *[]int64) int64 {
	address := (*mem)[position]
	return (*mem)[baseAddress+address]
}

func getValueBy(baseAddress int64, position int64, flag int, mem *[]int64) int64 {
	if flag == 2 {
		return getValueByAddressRelative(baseAddress, position, mem)
	}

	if flag == 1 {
		return getValueImmediate(position, mem)
	}

	return getValueByAddress(position, mem)
}

func putValueAtAddress(baseAddress int64, position int64, flag int, value int64, mem *[]int64) int64 {
	address := (*mem)[position]
	if flag == 2 {
		address = address + baseAddress
	}

	(*mem)[address] = value
	return (*mem)[address]
}

func opInput(baseAddress int64, codeRun int64, flags *[]int, mem *[]int64, input *[]int64, inputIndex int64) int64 {
	inputValue := (*input)[inputIndex]
	left := getValueImmediate(codeRun+1, mem)

	if (*flags)[0] == 2 {
		left = left + baseAddress
	}

	(*mem)[left] = inputValue
	return codeRun + 2
}

func opInputFromChannel(name string, codeRun int64, flags *[]int, mem *[]int64, input chan int64) int64 {
	inputValue := int64(0)
	select {
	case res := <-input:
		inputValue = res
	case <-time.After(3 * time.Second):
		fmt.Printf("input timeout name=%s\n", name)
		panic(fmt.Sprintf("input timeout name=%s\n", name))
	}

	left := getValueImmediate(codeRun+1, mem)
	(*mem)[left] = inputValue
	return codeRun + 2
}

func opOutput(baseAddress int64, codeRun int64, flags *[]int, mem *[]int64, output *[]int64) int64 {
	if (*flags)[0] > 0 {
		left := int64(0)
		if (*flags)[0] == 2 {
			left = getValueByAddressRelative(baseAddress, codeRun+1, mem)
		} else {
			left = getValueImmediate(codeRun+1, mem)
		}

		fmt.Printf("* DEBUG OUTPUT: %d\n", left)
		(*output)[0] = left
		return codeRun + 2
	}

	left := getValueByAddress(codeRun+1, mem)
	fmt.Printf("* DEBUG OUTPUT: %d\n", left)
	(*output)[0] = left
	return codeRun + 2
}

func opOutputIntoChannel(name string, codeRun int64, flags *[]int, mem *[]int64, output chan int64) (int64, int64) {
	if (*flags)[0] > 0 {
		left := getValueImmediate(codeRun+1, mem)
		fmt.Printf("* DEBUG OUTPUT: codeRun=%d left=%d  \n", codeRun, left)
		output <- left
		return codeRun + 2, left
	}

	left := getValueByAddress(codeRun+1, mem)
	fmt.Printf("* DEBUG OUTPUT: codeRun=%d left=%d\n", codeRun, left)
	output <- left
	return codeRun + 2, left
}

func opAdd(baseAddress int64, codeRun int64, flags *[]int, mem *[]int64) int64 {
	left := getValueBy(baseAddress, codeRun+1, (*flags)[0], mem)
	right := getValueBy(baseAddress, codeRun+2, (*flags)[1], mem)
	putValueAtAddress(baseAddress, codeRun+3, (*flags)[2], left+right, mem)
	return codeRun + 4
}

func opMult(baseAddress int64, codeRun int64, flags *[]int, mem *[]int64) int64 {
	left := getValueBy(baseAddress, codeRun+1, (*flags)[0], mem)
	right := getValueBy(baseAddress, codeRun+2, (*flags)[1], mem)
	putValueAtAddress(baseAddress, codeRun+3, (*flags)[2], left*right, mem)
	return codeRun + 4
}

func opJumpTrue(baseAddress int64, codeRun int64, flags *[]int, mem *[]int64) int64 {
	left := getValueBy(baseAddress, codeRun+1, (*flags)[0], mem)
	right := getValueBy(baseAddress, codeRun+2, (*flags)[1], mem)
	if left > 0 {
		return right
	}
	return codeRun + 3
}

func opJumpFalse(baseAddress int64, codeRun int64, flags *[]int, mem *[]int64) int64 {
	left := getValueBy(baseAddress, codeRun+1, (*flags)[0], mem)
	right := getValueBy(baseAddress, codeRun+2, (*flags)[1], mem)
	if left == 0 {
		return right
	}
	return codeRun + 3
}

func opLessThan(baseAddress int64, codeRun int64, flags *[]int, mem *[]int64) int64 {
	left := getValueBy(baseAddress, codeRun+1, (*flags)[0], mem)
	right := getValueBy(baseAddress, codeRun+2, (*flags)[1], mem)

	if left < right {
		putValueAtAddress(baseAddress, codeRun+3, (*flags)[2], 1, mem)
		return codeRun + 4

	}

	putValueAtAddress(baseAddress, codeRun+3, (*flags)[2], 0, mem)
	return codeRun + 4
}

func opEquals(baseAddress int64, codeRun int64, flags *[]int, mem *[]int64) int64 {
	left := getValueBy(baseAddress, codeRun+1, (*flags)[0], mem)
	right := getValueBy(baseAddress, codeRun+2, (*flags)[1], mem)

	if left == right {
		putValueAtAddress(baseAddress, codeRun+3, (*flags)[2], 1, mem)
		return codeRun + 4
	}

	putValueAtAddress(baseAddress, codeRun+3, (*flags)[2], 0, mem)
	return codeRun + 4
}

// Run execute a program in the int computer
// accepts program memory and inputs
// returns program memory and outputs
func Run(memory []int64, input []int64) ([]int64, []int64) {

	var baseAddress = int64(0)
	var output = make([]int64, len(input))
	var intputIndex = int64(0)

	var codeRun = int64(0)
	var done = false
	for !done {
		instruction := memory[codeRun]
		opcode := int(instruction % 100)

		flags := []int{
			int(instruction/100) % 10,
			int(instruction/1000) % 10,
			int(instruction/10000) % 10,
		}

		switch opcode {
		default:
			fmt.Printf("Unknown instruction %d at addr %d\n", opcode, codeRun)
			os.Exit(99)
		case opCodeAdd:
			codeRun = opAdd(baseAddress, codeRun, &flags, &memory)
		case opCodeMult:
			codeRun = opMult(baseAddress, codeRun, &flags, &memory)
		case opCodeInput:
			codeRun = opInput(baseAddress, codeRun, &flags, &memory, &input, intputIndex)
			intputIndex++
		case opCodeOutput:
			codeRun = opOutput(baseAddress, codeRun, &flags, &memory, &output)
		case opCodeJumpIfTrue:
			codeRun = opJumpTrue(baseAddress, codeRun, &flags, &memory)
		case opCodeJumpIfFalse:
			codeRun = opJumpFalse(baseAddress, codeRun, &flags, &memory)
		case opCodeLessThan:
			codeRun = opLessThan(baseAddress, codeRun, &flags, &memory)
		case opCodeEquals:
			codeRun = opEquals(baseAddress, codeRun, &flags, &memory)
		case opCodeSetBase:
			value := getValueBy(baseAddress, codeRun+1, flags[0], &memory)
			baseAddress = baseAddress + value
			//fmt.Printf("new base address: (flag:%d) %d\n", flags[0], baseAddress)
			codeRun = codeRun + 2
		case opCodeExit:
			done = true
		}
	}

	return memory, output
}

// RunWithChannels run but inputs and outputs are callbacks
func RunWithChannels(memory []int64, name string, input chan int64, output chan int64, exit chan int64) ([]int64, int64) {
	var baseAddress = int64(0)

	var intputIndex = 0
	var lastOutput = int64(0)

	var codeRun = int64(0)
	var done = false
	for !done {
		instruction := memory[codeRun]
		opcode := int(instruction % 100)

		flags := []int{
			int(instruction/100) % 10,
			int(instruction/1000) % 10,
			int(instruction/10000) % 10,
		}

		switch opcode {
		default:
			fmt.Printf("Unknown instruction %d at addr %d\n", opcode, codeRun)
			os.Exit(99)
		case opCodeAdd:
			codeRun = opAdd(baseAddress, codeRun, &flags, &memory)
		case opCodeMult:
			codeRun = opMult(baseAddress, codeRun, &flags, &memory)
		case opCodeInput:
			codeRun = opInputFromChannel(name, codeRun, &flags, &memory, input)
			intputIndex++
		case opCodeOutput:
			codeRun, lastOutput = opOutputIntoChannel(name, codeRun, &flags, &memory, output)
		case opCodeJumpIfTrue:
			codeRun = opJumpTrue(baseAddress, codeRun, &flags, &memory)
		case opCodeJumpIfFalse:
			codeRun = opJumpFalse(baseAddress, codeRun, &flags, &memory)
		case opCodeLessThan:
			codeRun = opLessThan(baseAddress, codeRun, &flags, &memory)
		case opCodeEquals:
			codeRun = opEquals(baseAddress, codeRun, &flags, &memory)
		case opCodeExit:
			done = true
			exit <- lastOutput
		}
	}

	return memory, lastOutput
}

// LoadInstructions .
func LoadInstructions(inputText string) []int64 {
	inputClean := strings.ReplaceAll(inputText, "\n", "")
	inputClean = strings.ReplaceAll(inputClean, " ", "")
	inputClean = strings.ReplaceAll(inputClean, "\t", "")

	inputValuesString := strings.Split(
		inputClean,
		",")
	input := make([]int64, len(inputValuesString))
	for inputPosition := range inputValuesString {
		converted, err := strconv.ParseInt(inputValuesString[inputPosition], 10, 64)
		if err != nil {
			panic(err)
		}
		input[inputPosition] = converted
	}

	return input
}

// LoadInstructionsWithMemoryAlloc .
func LoadInstructionsWithMemoryAlloc(inputText string, memoryLength int64) []int64 {
	inputClean := strings.ReplaceAll(inputText, "\n", "")
	inputClean = strings.ReplaceAll(inputClean, " ", "")
	inputClean = strings.ReplaceAll(inputClean, "\t", "")

	inputValuesString := strings.Split(
		inputClean,
		",")
	input := make([]int64, memoryLength)
	for inputPosition := range inputValuesString {
		converted, err := strconv.ParseInt(inputValuesString[inputPosition], 10, 64)
		if err != nil {
			panic(err)
		}
		input[inputPosition] = converted
	}

	return input
}
