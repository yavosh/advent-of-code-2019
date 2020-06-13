package main

import (
	"fmt"
	"io/ioutil"

	"github.com/yavosh/advent-of-code-2019/computer"
)

func thrusters(code string, input []int) int {
	codeInput := []int{0, 0}
	for _, thrusterInput := range input {
		codeInput[0] = thrusterInput
		fmt.Printf("int: %v\n", codeInput)
		_, output := computer.Run(computer.LoadInstructions(code), codeInput)
		fmt.Printf("out: %v\n", output)
		codeInput[1] = output[0]
	}
	// last output
	return codeInput[1]
}

var ()

func main() {
	thrusterCode, err := ioutil.ReadFile("./thruster_code_d7.txt")
	if err != nil {
		panic(err)
	}

	maxValue := 0
	var maxInput []int

	for _, input := range permutations {
		out := thrusters(string(thrusterCode), input)

		if out > maxValue {
			maxValue = out
			maxInput = input
		}
	}

	fmt.Printf("max output: %d input: %v", maxValue, maxInput)
}
