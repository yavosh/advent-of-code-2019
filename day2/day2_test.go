package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCode(t *testing.T) {
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
		result := execute(testCase.inputCode)
		assert.Equal(t, testCase.outputCode, result, fmt.Sprintf("for [%v!=%v] ", testCase.inputCode, result))
	}
}
