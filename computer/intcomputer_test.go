package computer_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yavosh/advent-of-code-2019/computer"
)

func TestBasicOperations(t *testing.T) {
	testCases := []struct {
		inputCode  []int
		outputCode []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, testCase := range testCases {
		result, _ := computer.Run(testCase.inputCode, []int{})
		assert.Equal(t, testCase.outputCode, result, fmt.Sprintf("for [%v!=%v] ", testCase.inputCode, result))
	}

}

func TestInputOperations(t *testing.T) {
	testCases := []struct {
		inputCode   []int
		inputValue  int
		outputCode  []int
		outputValue int
	}{
		{[]int{3, 0, 4, 0, 99}, 10, []int{10, 0, 4, 0, 99}, 10},
		{[]int{1002, 4, 3, 4, 33}, 0, []int{1002, 4, 3, 4, 99}, 0},
		{[]int{1101, 100, -1, 4, 0}, 0, []int{1101, 100, -1, 4, 99}, 0},
	}

	for _, testCase := range testCases {
		result, outputs := computer.Run(testCase.inputCode, []int{testCase.inputValue})
		outValue := outputs[0]

		assert.Equal(t, testCase.outputCode, result,
			fmt.Sprintf("for [%v!=%v] ", testCase.inputCode, result))
		assert.Equal(t, testCase.outputValue, outValue,
			fmt.Sprintf("for [%v!=%v] ", testCase.outputValue, outValue))
	}
}

func TestCompareOperations(t *testing.T) {

	var equalTo8Program = []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	var equalTo8ImmediateProgram = []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	var lessThan8Program = []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	var lessThan8ImmediateProgram = []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}

	testCases := []struct {
		inputCode   []int
		inputValue  int
		outputCode  []int
		outputValue int
	}{
		{equalTo8Program, 8, []int{0}, 1},
		{equalTo8Program, 7, []int{0}, 0},
		{equalTo8Program, 5, []int{0}, 0},
		{equalTo8Program, 11, []int{0}, 0},

		{equalTo8ImmediateProgram, 8, []int{0}, 1},
		{equalTo8ImmediateProgram, 7, []int{0}, 0},
		{equalTo8ImmediateProgram, 5, []int{0}, 0},
		{equalTo8ImmediateProgram, 11, []int{0}, 0},

		{lessThan8Program, 5, []int{0}, 1},
		{lessThan8Program, 7, []int{0}, 1},
		{lessThan8Program, 8, []int{0}, 0},
		{lessThan8Program, 11, []int{0}, 0},

		{lessThan8ImmediateProgram, 5, []int{0}, 1},
		{lessThan8ImmediateProgram, 7, []int{0}, 1},
		{lessThan8ImmediateProgram, 8, []int{0}, 0},
		{lessThan8ImmediateProgram, 11, []int{0}, 0},
	}

	for _, testCase := range testCases {
		_, outputs := computer.Run(testCase.inputCode, []int{testCase.inputValue})
		outValue := outputs[0]
		assert.Equal(t, testCase.outputValue, outValue,
			fmt.Sprintf("for [%v!=%v] ", testCase.outputValue, outValue))
	}

}

func TestJumpOperations(t *testing.T) {
	testCases := []struct {
		inputCode   []int
		inputValue  int
		outputCode  []int
		outputValue int
	}{
		{[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, 0, []int{0}, 0},
		{[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, 1, []int{0}, 1},
		{[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, 2, []int{0}, 1},
		{[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, 0, []int{0}, 0},
		{[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, 1, []int{0}, 1},
		{[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, 2, []int{0}, 1},
	}

	for _, testCase := range testCases {
		_, outputs := computer.Run(testCase.inputCode, []int{testCase.inputValue})
		outValue := outputs[0]
		assert.Equal(t, testCase.outputValue, outValue,
			fmt.Sprintf("for [%v!=%v] ", testCase.outputValue, outValue))
	}

}

func TestWorkingExample(t *testing.T) {
	var code = `3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
	1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
	999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99`
	instructions := computer.LoadInstructions(code)

	testCases := []struct {
		inputValue  int
		outputValue int
	}{
		{1, 999},
		{7, 999},
		{8, 1000},
		{9, 1001},
		{100, 1001},
	}

	for _, testCase := range testCases {
		_, outputs := computer.Run(instructions, []int{testCase.inputValue})
		outValue := outputs[0]
		assert.Equal(t, testCase.outputValue, outValue,
			fmt.Sprintf("for input %d unexpetec output [%v!=%v] ", testCase.inputValue, testCase.outputValue, outValue))
	}
}
